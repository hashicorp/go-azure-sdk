package pendingaccessreviewinstancestagedecisioninsight

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPendingAccessReviewInstanceStageDecisionInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GovernanceInsight
}

type ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GovernanceInsight
}

type ListPendingAccessReviewInstanceStageDecisionInsightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceStageDecisionInsightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceStageDecisionInsights ...
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsights(ctx context.Context, id UserIdPendingAccessReviewInstanceIdStageIdDecisionId) (result ListPendingAccessReviewInstanceStageDecisionInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceStageDecisionInsightsCustomPager{},
		Path:       fmt.Sprintf("%s/insights", id.ID()),
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
		Values *[]beta.GovernanceInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceStageDecisionInsightsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsightsComplete(ctx context.Context, id UserIdPendingAccessReviewInstanceIdStageIdDecisionId) (ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx, id, GovernanceInsightOperationPredicate{})
}

// ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx context.Context, id UserIdPendingAccessReviewInstanceIdStageIdDecisionId, predicate GovernanceInsightOperationPredicate) (result ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult, err error) {
	items := make([]beta.GovernanceInsight, 0)

	resp, err := c.ListPendingAccessReviewInstanceStageDecisionInsights(ctx, id)
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

	result = ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
