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

type ListSiteAnalysesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AnalysisDefinition
}

type ListSiteAnalysesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AnalysisDefinition
}

// ListSiteAnalysesSlot ...
func (c DiagnosticsClient) ListSiteAnalysesSlot(ctx context.Context, id SlotDiagnosticId) (result ListSiteAnalysesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/analyses", id.ID()),
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
		Values *[]AnalysisDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteAnalysesSlotComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteAnalysesSlotComplete(ctx context.Context, id SlotDiagnosticId) (ListSiteAnalysesSlotCompleteResult, error) {
	return c.ListSiteAnalysesSlotCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// ListSiteAnalysesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteAnalysesSlotCompleteMatchingPredicate(ctx context.Context, id SlotDiagnosticId, predicate AnalysisDefinitionOperationPredicate) (result ListSiteAnalysesSlotCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	resp, err := c.ListSiteAnalysesSlot(ctx, id)
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

	result = ListSiteAnalysesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
