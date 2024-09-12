package accessreviewdefinitioninstancestagedecision

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

type ListAccessReviewDefinitionInstanceStageDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAccessReviewDefinitionInstanceStageDecisionsOperationOptions() ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions {
	return ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceStageDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStageDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStageDecisions - List decisions (from a multi-stage access review). Get the
// decisions from a stage in a multi-stage access review. The decisions in an accessReviewStage object are represented
// by an accessReviewInstanceDecisionItem object.
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, options ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions) (result ListAccessReviewDefinitionInstanceStageDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceStageDecisionsCustomPager{},
		Path:          fmt.Sprintf("%s/decisions", id.ID()),
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

// ListAccessReviewDefinitionInstanceStageDecisionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionsComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, options ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions) (ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate(ctx, id, options, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, options ListAccessReviewDefinitionInstanceStageDecisionsOperationOptions, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult, err error) {
	items := make([]beta.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStageDecisions(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
