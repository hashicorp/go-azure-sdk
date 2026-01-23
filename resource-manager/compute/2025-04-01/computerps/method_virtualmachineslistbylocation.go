package computerps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachinesListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachine
}

type VirtualMachinesListByLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualMachine
}

type VirtualMachinesListByLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VirtualMachinesListByLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VirtualMachinesListByLocation ...
func (c ComputeRPSClient) VirtualMachinesListByLocation(ctx context.Context, id LocationId) (result VirtualMachinesListByLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VirtualMachinesListByLocationCustomPager{},
		Path:       fmt.Sprintf("%s/virtualMachines", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]VirtualMachine `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VirtualMachinesListByLocationComplete retrieves all the results into a single object
func (c ComputeRPSClient) VirtualMachinesListByLocationComplete(ctx context.Context, id LocationId) (VirtualMachinesListByLocationCompleteResult, error) {
	return c.VirtualMachinesListByLocationCompleteMatchingPredicate(ctx, id, VirtualMachineOperationPredicate{})
}

// VirtualMachinesListByLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeRPSClient) VirtualMachinesListByLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate VirtualMachineOperationPredicate) (result VirtualMachinesListByLocationCompleteResult, err error) {
	items := make([]VirtualMachine, 0)

	resp, err := c.VirtualMachinesListByLocation(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = VirtualMachinesListByLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
