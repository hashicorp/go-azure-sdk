package accessreviewdecisioninstancecontactedreviewer

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

type ListAccessReviewDecisionInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewReviewer
}

type ListAccessReviewDecisionInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewReviewer
}

type ListAccessReviewDecisionInstanceContactedReviewersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAccessReviewDecisionInstanceContactedReviewersOperationOptions() ListAccessReviewDecisionInstanceContactedReviewersOperationOptions {
	return ListAccessReviewDecisionInstanceContactedReviewersOperationOptions{}
}

func (o ListAccessReviewDecisionInstanceContactedReviewersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDecisionInstanceContactedReviewersOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDecisionInstanceContactedReviewersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDecisionInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDecisionInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDecisionInstanceContactedReviewers - Get contactedReviewers from identityGovernance. Returns the
// collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers
// properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers,
// contactedReviewers returns their individual identities. Supports $select. Read-only.
func (c AccessReviewDecisionInstanceContactedReviewerClient) ListAccessReviewDecisionInstanceContactedReviewers(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options ListAccessReviewDecisionInstanceContactedReviewersOperationOptions) (result ListAccessReviewDecisionInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDecisionInstanceContactedReviewersCustomPager{},
		Path:          fmt.Sprintf("%s/instance/contactedReviewers", id.ID()),
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
		Values *[]beta.AccessReviewReviewer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDecisionInstanceContactedReviewersComplete retrieves all the results into a single object
func (c AccessReviewDecisionInstanceContactedReviewerClient) ListAccessReviewDecisionInstanceContactedReviewersComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options ListAccessReviewDecisionInstanceContactedReviewersOperationOptions) (ListAccessReviewDecisionInstanceContactedReviewersCompleteResult, error) {
	return c.ListAccessReviewDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, options, AccessReviewReviewerOperationPredicate{})
}

// ListAccessReviewDecisionInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDecisionInstanceContactedReviewerClient) ListAccessReviewDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options ListAccessReviewDecisionInstanceContactedReviewersOperationOptions, predicate AccessReviewReviewerOperationPredicate) (result ListAccessReviewDecisionInstanceContactedReviewersCompleteResult, err error) {
	items := make([]beta.AccessReviewReviewer, 0)

	resp, err := c.ListAccessReviewDecisionInstanceContactedReviewers(ctx, id, options)
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

	result = ListAccessReviewDecisionInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
