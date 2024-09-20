package accessreviewdefinitioninstancestagedecisioninsight

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.GovernanceInsight
}

type ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.GovernanceInsight
}

type ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions() ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions {
	return ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceStageDecisionInsightsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStageDecisionInsightsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStageDecisionInsights - Get insights from identityGovernance. Insights are
// recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with
// an accessReviewInstanceDecisionItem.
func (c AccessReviewDefinitionInstanceStageDecisionInsightClient) ListAccessReviewDefinitionInstanceStageDecisionInsights(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions) (result ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceStageDecisionInsightsCustomPager{},
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

	temp := make([]stable.GovernanceInsight, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalGovernanceInsightImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.GovernanceInsight (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAccessReviewDefinitionInstanceStageDecisionInsightsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageDecisionInsightClient) ListAccessReviewDefinitionInstanceStageDecisionInsightsComplete(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions) (ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx, id, options, GovernanceInsightOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageDecisionInsightClient) ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInsightsOperationOptions, predicate GovernanceInsightOperationPredicate) (result ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteResult, err error) {
	items := make([]stable.GovernanceInsight, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStageDecisionInsights(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceStageDecisionInsightsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
