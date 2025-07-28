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

type VirtualMachineSizesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachineSize
}

type VirtualMachineSizesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualMachineSize
}

type VirtualMachineSizesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VirtualMachineSizesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VirtualMachineSizesList ...
func (c ComputeRPSClient) VirtualMachineSizesList(ctx context.Context, id LocationId) (result VirtualMachineSizesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VirtualMachineSizesListCustomPager{},
		Path:       fmt.Sprintf("%s/vmSizes", id.ID()),
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
		Values *[]VirtualMachineSize `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VirtualMachineSizesListComplete retrieves all the results into a single object
func (c ComputeRPSClient) VirtualMachineSizesListComplete(ctx context.Context, id LocationId) (VirtualMachineSizesListCompleteResult, error) {
	return c.VirtualMachineSizesListCompleteMatchingPredicate(ctx, id, VirtualMachineSizeOperationPredicate{})
}

// VirtualMachineSizesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeRPSClient) VirtualMachineSizesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate VirtualMachineSizeOperationPredicate) (result VirtualMachineSizesListCompleteResult, err error) {
	items := make([]VirtualMachineSize, 0)

	resp, err := c.VirtualMachineSizesList(ctx, id)
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

	result = VirtualMachineSizesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
