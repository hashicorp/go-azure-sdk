package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteDetectorResponsesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type ListSiteDetectorResponsesSlotCompleteResult struct {
	Items []DetectorResponse
}

// ListSiteDetectorResponsesSlot ...
func (c DiagnosticsClient) ListSiteDetectorResponsesSlot(ctx context.Context, id SlotId) (result ListSiteDetectorResponsesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/detectors", id.ID()),
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
		Values *[]DetectorResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteDetectorResponsesSlotComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteDetectorResponsesSlotComplete(ctx context.Context, id SlotId) (ListSiteDetectorResponsesSlotCompleteResult, error) {
	return c.ListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListSiteDetectorResponsesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DetectorResponseOperationPredicate) (result ListSiteDetectorResponsesSlotCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.ListSiteDetectorResponsesSlot(ctx, id)
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

	result = ListSiteDetectorResponsesSlotCompleteResult{
		Items: items,
	}
	return
}
