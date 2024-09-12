package pendingaccessreviewinstancecontactedreviewer

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

type ListPendingAccessReviewInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceContactedReviewersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPendingAccessReviewInstanceContactedReviewersOperationOptions() ListPendingAccessReviewInstanceContactedReviewersOperationOptions {
	return ListPendingAccessReviewInstanceContactedReviewersOperationOptions{}
}

func (o ListPendingAccessReviewInstanceContactedReviewersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceContactedReviewersOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstanceContactedReviewersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceContactedReviewers - Get contactedReviewers from users. Returns the collection of
// reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the
// accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their
// individual identities. Supports $select. Read-only.
func (c PendingAccessReviewInstanceContactedReviewerClient) ListPendingAccessReviewInstanceContactedReviewers(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceContactedReviewersOperationOptions) (result ListPendingAccessReviewInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceContactedReviewersCustomPager{},
		Path:          fmt.Sprintf("%s/contactedReviewers", id.ID()),
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

// ListPendingAccessReviewInstanceContactedReviewersComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceContactedReviewerClient) ListPendingAccessReviewInstanceContactedReviewersComplete(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceContactedReviewersOperationOptions) (ListPendingAccessReviewInstanceContactedReviewersCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, options, AccessReviewReviewerOperationPredicate{})
}

// ListPendingAccessReviewInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceContactedReviewerClient) ListPendingAccessReviewInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceContactedReviewersOperationOptions, predicate AccessReviewReviewerOperationPredicate) (result ListPendingAccessReviewInstanceContactedReviewersCompleteResult, err error) {
	items := make([]beta.AccessReviewReviewer, 0)

	resp, err := c.ListPendingAccessReviewInstanceContactedReviewers(ctx, id, options)
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

	result = ListPendingAccessReviewInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
