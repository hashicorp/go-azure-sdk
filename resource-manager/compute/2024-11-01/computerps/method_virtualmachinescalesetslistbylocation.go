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

type VirtualMachineScaleSetsListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachineScaleSet
}

type VirtualMachineScaleSetsListByLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualMachineScaleSet
}

type VirtualMachineScaleSetsListByLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VirtualMachineScaleSetsListByLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VirtualMachineScaleSetsListByLocation ...
func (c ComputeRPSClient) VirtualMachineScaleSetsListByLocation(ctx context.Context, id LocationId) (result VirtualMachineScaleSetsListByLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VirtualMachineScaleSetsListByLocationCustomPager{},
		Path:       fmt.Sprintf("%s/virtualMachineScaleSets", id.ID()),
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
		Values *[]VirtualMachineScaleSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VirtualMachineScaleSetsListByLocationComplete retrieves all the results into a single object
func (c ComputeRPSClient) VirtualMachineScaleSetsListByLocationComplete(ctx context.Context, id LocationId) (VirtualMachineScaleSetsListByLocationCompleteResult, error) {
	return c.VirtualMachineScaleSetsListByLocationCompleteMatchingPredicate(ctx, id, VirtualMachineScaleSetOperationPredicate{})
}

// VirtualMachineScaleSetsListByLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeRPSClient) VirtualMachineScaleSetsListByLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate VirtualMachineScaleSetOperationPredicate) (result VirtualMachineScaleSetsListByLocationCompleteResult, err error) {
	items := make([]VirtualMachineScaleSet, 0)

	resp, err := c.VirtualMachineScaleSetsListByLocation(ctx, id)
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

	result = VirtualMachineScaleSetsListByLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
