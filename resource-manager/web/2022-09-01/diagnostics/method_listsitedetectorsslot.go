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

type ListSiteDetectorsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorDefinitionResource
}

type ListSiteDetectorsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorDefinitionResource
}

// ListSiteDetectorsSlot ...
func (c DiagnosticsClient) ListSiteDetectorsSlot(ctx context.Context, id SlotDiagnosticId) (result ListSiteDetectorsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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
		Values *[]DetectorDefinitionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteDetectorsSlotComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteDetectorsSlotComplete(ctx context.Context, id SlotDiagnosticId) (ListSiteDetectorsSlotCompleteResult, error) {
	return c.ListSiteDetectorsSlotCompleteMatchingPredicate(ctx, id, DetectorDefinitionResourceOperationPredicate{})
}

// ListSiteDetectorsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteDetectorsSlotCompleteMatchingPredicate(ctx context.Context, id SlotDiagnosticId, predicate DetectorDefinitionResourceOperationPredicate) (result ListSiteDetectorsSlotCompleteResult, err error) {
	items := make([]DetectorDefinitionResource, 0)

	resp, err := c.ListSiteDetectorsSlot(ctx, id)
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

	result = ListSiteDetectorsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
