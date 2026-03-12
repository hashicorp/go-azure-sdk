package bysiteslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthMetadataListBySiteSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceHealthMetadata
}

type ResourceHealthMetadataListBySiteSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceHealthMetadata
}

type ResourceHealthMetadataListBySiteSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceHealthMetadataListBySiteSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceHealthMetadataListBySiteSlot ...
func (c BySiteSlotOperationGroupClient) ResourceHealthMetadataListBySiteSlot(ctx context.Context, id SlotId) (result ResourceHealthMetadataListBySiteSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceHealthMetadataListBySiteSlotCustomPager{},
		Path:       fmt.Sprintf("%s/resourceHealthMetadata", id.ID()),
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
		Values *[]ResourceHealthMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceHealthMetadataListBySiteSlotComplete retrieves all the results into a single object
func (c BySiteSlotOperationGroupClient) ResourceHealthMetadataListBySiteSlotComplete(ctx context.Context, id SlotId) (ResourceHealthMetadataListBySiteSlotCompleteResult, error) {
	return c.ResourceHealthMetadataListBySiteSlotCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ResourceHealthMetadataListBySiteSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BySiteSlotOperationGroupClient) ResourceHealthMetadataListBySiteSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ResourceHealthMetadataOperationPredicate) (result ResourceHealthMetadataListBySiteSlotCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	resp, err := c.ResourceHealthMetadataListBySiteSlot(ctx, id)
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

	result = ResourceHealthMetadataListBySiteSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
