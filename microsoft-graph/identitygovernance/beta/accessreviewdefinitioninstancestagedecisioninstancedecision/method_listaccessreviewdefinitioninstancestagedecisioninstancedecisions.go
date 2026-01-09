package accessreviewdefinitioninstancestagedecisioninstancedecision

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions struct {
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

func DefaultListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions() ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions {
	return ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisions - Get decisions from identityGovernance. Each user
// reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet
// reviewed.
func (c AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions) (result ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCustomPager{},
		Path:          fmt.Sprintf("%s/instance/decisions", id.ID()),
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
		Values *[]beta.AccessReviewInstanceDecisionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions) (ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate(ctx, id, options, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteResult, err error) {
	items := make([]beta.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisions(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
