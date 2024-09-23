package pendingaccessreviewinstancedecisioninstancecontactedreviewer

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

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions struct {
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

func DefaultListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions() ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions {
	return ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions{}
}

func (o ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers - Get contactedReviewers from me. Returns the
// collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers
// properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers,
// contactedReviewers returns their individual identities. Supports $select. Read-only.
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions) (result ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager{},
		Path:          fmt.Sprintf("%s/instance/contactedReviewers", id.ID()),
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
		Values *[]beta.AccessReviewReviewer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersComplete(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions) (ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, options, AccessReviewReviewerOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationOptions, predicate AccessReviewReviewerOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult, err error) {
	items := make([]beta.AccessReviewReviewer, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers(ctx, id, options)
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

	result = ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
