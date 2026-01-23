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

type VirtualMachineImagesListByEdgeZoneOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachineImageResource
}

type VirtualMachineImagesListByEdgeZoneCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualMachineImageResource
}

type VirtualMachineImagesListByEdgeZoneCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VirtualMachineImagesListByEdgeZoneCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VirtualMachineImagesListByEdgeZone ...
func (c ComputeRPSClient) VirtualMachineImagesListByEdgeZone(ctx context.Context, id EdgeZoneId) (result VirtualMachineImagesListByEdgeZoneOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VirtualMachineImagesListByEdgeZoneCustomPager{},
		Path:       fmt.Sprintf("%s/vmimages", id.ID()),
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
		Values *[]VirtualMachineImageResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VirtualMachineImagesListByEdgeZoneComplete retrieves all the results into a single object
func (c ComputeRPSClient) VirtualMachineImagesListByEdgeZoneComplete(ctx context.Context, id EdgeZoneId) (VirtualMachineImagesListByEdgeZoneCompleteResult, error) {
	return c.VirtualMachineImagesListByEdgeZoneCompleteMatchingPredicate(ctx, id, VirtualMachineImageResourceOperationPredicate{})
}

// VirtualMachineImagesListByEdgeZoneCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeRPSClient) VirtualMachineImagesListByEdgeZoneCompleteMatchingPredicate(ctx context.Context, id EdgeZoneId, predicate VirtualMachineImageResourceOperationPredicate) (result VirtualMachineImagesListByEdgeZoneCompleteResult, err error) {
	items := make([]VirtualMachineImageResource, 0)

	resp, err := c.VirtualMachineImagesListByEdgeZone(ctx, id)
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

	result = VirtualMachineImagesListByEdgeZoneCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
