package analysisdefinitionoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteAnalysesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AnalysisDefinition
}

type DiagnosticsListSiteAnalysesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AnalysisDefinition
}

type DiagnosticsListSiteAnalysesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteAnalysesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteAnalysesSlot ...
func (c AnalysisDefinitionOperationGroupClient) DiagnosticsListSiteAnalysesSlot(ctx context.Context, id SlotDiagnosticId) (result DiagnosticsListSiteAnalysesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteAnalysesSlotCustomPager{},
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

// DiagnosticsListSiteAnalysesSlotComplete retrieves all the results into a single object
func (c AnalysisDefinitionOperationGroupClient) DiagnosticsListSiteAnalysesSlotComplete(ctx context.Context, id SlotDiagnosticId) (DiagnosticsListSiteAnalysesSlotCompleteResult, error) {
	return c.DiagnosticsListSiteAnalysesSlotCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// DiagnosticsListSiteAnalysesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AnalysisDefinitionOperationGroupClient) DiagnosticsListSiteAnalysesSlotCompleteMatchingPredicate(ctx context.Context, id SlotDiagnosticId, predicate AnalysisDefinitionOperationPredicate) (result DiagnosticsListSiteAnalysesSlotCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	resp, err := c.DiagnosticsListSiteAnalysesSlot(ctx, id)
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

	result = DiagnosticsListSiteAnalysesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
