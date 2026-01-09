package pendingaccessreviewinstancestagedecisioninsight

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

type ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions() ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions {
	return ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions{}
}

func (o ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListPendingAccessReviewInstanceStageDecisionInsights - Get insights from users. Insights are recommendations to
// reviewers on whether to approve or deny a decision. There can be multiple insights associated with an
// accessReviewInstanceDecisionItem.
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsights(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions) (result ListPendingAccessReviewInstanceStageDecisionInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceStageDecisionInsightsCustomPager{},
		Path:          fmt.Sprintf("%s/insights", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListPendingAccessReviewInstanceStageDecisionInsightsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsightsComplete(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions) (ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx, id, options, GovernanceInsightOperationPredicate{})
}

// ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageDecisionInsightClient) ListPendingAccessReviewInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInsightsOperationOptions, predicate GovernanceInsightOperationPredicate) (result ListPendingAccessReviewInstanceStageDecisionInsightsCompleteResult, err error) {
	items := make([]beta.GovernanceInsight, 0)

	resp, err := c.ListPendingAccessReviewInstanceStageDecisionInsights(ctx, id, options)
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
