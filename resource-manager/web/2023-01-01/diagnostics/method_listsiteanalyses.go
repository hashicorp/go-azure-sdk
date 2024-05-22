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

type ListSiteAnalysesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AnalysisDefinition
}

type ListSiteAnalysesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AnalysisDefinition
}

// ListSiteAnalyses ...
func (c DiagnosticsClient) ListSiteAnalyses(ctx context.Context, id DiagnosticId) (result ListSiteAnalysesOperationResponse, err error) {
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

// ListSiteAnalysesComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteAnalysesComplete(ctx context.Context, id DiagnosticId) (ListSiteAnalysesCompleteResult, error) {
	return c.ListSiteAnalysesCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// ListSiteAnalysesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteAnalysesCompleteMatchingPredicate(ctx context.Context, id DiagnosticId, predicate AnalysisDefinitionOperationPredicate) (result ListSiteAnalysesCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	resp, err := c.ListSiteAnalyses(ctx, id)
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

	result = ListSiteAnalysesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
