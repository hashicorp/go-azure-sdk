package triggeredanalyticsruleruns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTriggeredAnalyticsRuleRunsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredAnalyticsRuleRun
}

type GetTriggeredAnalyticsRuleRunsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggeredAnalyticsRuleRun
}

type GetTriggeredAnalyticsRuleRunsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetTriggeredAnalyticsRuleRunsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetTriggeredAnalyticsRuleRunsList ...
func (c TriggeredAnalyticsRuleRunsClient) GetTriggeredAnalyticsRuleRunsList(ctx context.Context, id WorkspaceId) (result GetTriggeredAnalyticsRuleRunsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetTriggeredAnalyticsRuleRunsListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/triggeredAnalyticsRuleRuns", id.ID()),
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
		Values *[]TriggeredAnalyticsRuleRun `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetTriggeredAnalyticsRuleRunsListComplete retrieves all the results into a single object
func (c TriggeredAnalyticsRuleRunsClient) GetTriggeredAnalyticsRuleRunsListComplete(ctx context.Context, id WorkspaceId) (GetTriggeredAnalyticsRuleRunsListCompleteResult, error) {
	return c.GetTriggeredAnalyticsRuleRunsListCompleteMatchingPredicate(ctx, id, TriggeredAnalyticsRuleRunOperationPredicate{})
}

// GetTriggeredAnalyticsRuleRunsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggeredAnalyticsRuleRunsClient) GetTriggeredAnalyticsRuleRunsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate TriggeredAnalyticsRuleRunOperationPredicate) (result GetTriggeredAnalyticsRuleRunsListCompleteResult, err error) {
	items := make([]TriggeredAnalyticsRuleRun, 0)

	resp, err := c.GetTriggeredAnalyticsRuleRunsList(ctx, id)
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

	result = GetTriggeredAnalyticsRuleRunsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
