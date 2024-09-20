package accessreviewdefinitioninstancedecision

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

type ListAccessReviewDefinitionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceDecisionsOperationOptions struct {
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

func DefaultListAccessReviewDefinitionInstanceDecisionsOperationOptions() ListAccessReviewDefinitionInstanceDecisionsOperationOptions {
	return ListAccessReviewDefinitionInstanceDecisionsOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceDecisions - List decisions. Retrieve the accessReviewInstanceDecisionItem objects
// for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned,
// including all of their nested properties.
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceDecisionsOperationOptions) (result ListAccessReviewDefinitionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceDecisionsCustomPager{},
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

// ListAccessReviewDefinitionInstanceDecisionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisionsComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceDecisionsOperationOptions) (ListAccessReviewDefinitionInstanceDecisionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate(ctx, id, options, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceDecisionsOperationOptions, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListAccessReviewDefinitionInstanceDecisionsCompleteResult, err error) {
	items := make([]beta.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceDecisions(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
