package threatintelligence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndicatorQueryIndicatorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ThreatIntelligenceInformationList
}

type IndicatorQueryIndicatorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ThreatIntelligenceInformationList
}

// IndicatorQueryIndicators ...
func (c ThreatIntelligenceClient) IndicatorQueryIndicators(ctx context.Context, id WorkspaceId, input ThreatIntelligenceFilteringCriteria) (result IndicatorQueryIndicatorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/threatIntelligence/main/queryIndicators", id.ID()),
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
		Values *[]ThreatIntelligenceInformationList `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// IndicatorQueryIndicatorsComplete retrieves all the results into a single object
func (c ThreatIntelligenceClient) IndicatorQueryIndicatorsComplete(ctx context.Context, id WorkspaceId, input ThreatIntelligenceFilteringCriteria) (IndicatorQueryIndicatorsCompleteResult, error) {
	return c.IndicatorQueryIndicatorsCompleteMatchingPredicate(ctx, id, input, ThreatIntelligenceInformationListOperationPredicate{})
}

// IndicatorQueryIndicatorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreatIntelligenceClient) IndicatorQueryIndicatorsCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, input ThreatIntelligenceFilteringCriteria, predicate ThreatIntelligenceInformationListOperationPredicate) (result IndicatorQueryIndicatorsCompleteResult, err error) {
	items := make([]ThreatIntelligenceInformationList, 0)

	resp, err := c.IndicatorQueryIndicators(ctx, id, input)
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

	result = IndicatorQueryIndicatorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
