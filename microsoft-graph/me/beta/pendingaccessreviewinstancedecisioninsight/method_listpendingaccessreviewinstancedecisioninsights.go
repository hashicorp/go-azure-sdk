package pendingaccessreviewinstancedecisioninsight

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

type ListPendingAccessReviewInstanceDecisionInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GovernanceInsight
}

type ListPendingAccessReviewInstanceDecisionInsightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GovernanceInsight
}

type ListPendingAccessReviewInstanceDecisionInsightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceDecisionInsightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceDecisionInsights ...
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsights(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId) (result ListPendingAccessReviewInstanceDecisionInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceDecisionInsightsCustomPager{},
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

// ListPendingAccessReviewInstanceDecisionInsightsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsightsComplete(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId) (ListPendingAccessReviewInstanceDecisionInsightsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate(ctx, id, GovernanceInsightOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId, predicate GovernanceInsightOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInsightsCompleteResult, err error) {
	items := make([]beta.GovernanceInsight, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInsights(ctx, id)
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

	result = ListPendingAccessReviewInstanceDecisionInsightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
