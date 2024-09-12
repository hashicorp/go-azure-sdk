package pendingaccessreviewinstancedecisioninsight

import (
	"context"
	"encoding/json"
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

type ListPendingAccessReviewInstanceDecisionInsightsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPendingAccessReviewInstanceDecisionInsightsOperationOptions() ListPendingAccessReviewInstanceDecisionInsightsOperationOptions {
	return ListPendingAccessReviewInstanceDecisionInsightsOperationOptions{}
}

func (o ListPendingAccessReviewInstanceDecisionInsightsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceDecisionInsightsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListPendingAccessReviewInstanceDecisionInsightsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListPendingAccessReviewInstanceDecisionInsights - Get insights from me. Insights are recommendations to reviewers on
// whether to approve or deny a decision. There can be multiple insights associated with an
// accessReviewInstanceDecisionItem.
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsights(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInsightsOperationOptions) (result ListPendingAccessReviewInstanceDecisionInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceDecisionInsightsCustomPager{},
		Path:          fmt.Sprintf("%s/insights", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.GovernanceInsight, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalGovernanceInsightImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.GovernanceInsight (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListPendingAccessReviewInstanceDecisionInsightsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsightsComplete(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInsightsOperationOptions) (ListPendingAccessReviewInstanceDecisionInsightsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate(ctx, id, options, GovernanceInsightOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInsightClient) ListPendingAccessReviewInstanceDecisionInsightsCompleteMatchingPredicate(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInsightsOperationOptions, predicate GovernanceInsightOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInsightsCompleteResult, err error) {
	items := make([]beta.GovernanceInsight, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInsights(ctx, id, options)
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
