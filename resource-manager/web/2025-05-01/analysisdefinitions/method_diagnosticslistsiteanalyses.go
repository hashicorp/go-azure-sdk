package analysisdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteAnalysesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AnalysisDefinition
}

type DiagnosticsListSiteAnalysesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AnalysisDefinition
}

type DiagnosticsListSiteAnalysesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteAnalysesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteAnalyses ...
func (c AnalysisDefinitionsClient) DiagnosticsListSiteAnalyses(ctx context.Context, id DiagnosticId) (result DiagnosticsListSiteAnalysesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteAnalysesCustomPager{},
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

// DiagnosticsListSiteAnalysesComplete retrieves all the results into a single object
func (c AnalysisDefinitionsClient) DiagnosticsListSiteAnalysesComplete(ctx context.Context, id DiagnosticId) (DiagnosticsListSiteAnalysesCompleteResult, error) {
	return c.DiagnosticsListSiteAnalysesCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// DiagnosticsListSiteAnalysesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AnalysisDefinitionsClient) DiagnosticsListSiteAnalysesCompleteMatchingPredicate(ctx context.Context, id DiagnosticId, predicate AnalysisDefinitionOperationPredicate) (result DiagnosticsListSiteAnalysesCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	resp, err := c.DiagnosticsListSiteAnalyses(ctx, id)
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

	result = DiagnosticsListSiteAnalysesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
