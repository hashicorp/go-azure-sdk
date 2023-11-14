package resourcehealthmetadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySiteSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceHealthMetadata
}

type ListBySiteSlotCompleteResult struct {
	Items []ResourceHealthMetadata
}

// ListBySiteSlot ...
func (c ResourceHealthMetadataClient) ListBySiteSlot(ctx context.Context, id SlotId) (result ListBySiteSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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

// ListBySiteSlotComplete retrieves all the results into a single object
func (c ResourceHealthMetadataClient) ListBySiteSlotComplete(ctx context.Context, id SlotId) (ListBySiteSlotCompleteResult, error) {
	return c.ListBySiteSlotCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ListBySiteSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceHealthMetadataClient) ListBySiteSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ResourceHealthMetadataOperationPredicate) (result ListBySiteSlotCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	resp, err := c.ListBySiteSlot(ctx, id)
	if err != nil {
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

	result = ListBySiteSlotCompleteResult{
		Items: items,
	}
	return
}
