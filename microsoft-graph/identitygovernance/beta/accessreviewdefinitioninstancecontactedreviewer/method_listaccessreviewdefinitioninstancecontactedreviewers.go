package accessreviewdefinitioninstancecontactedreviewer

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

type ListAccessReviewDefinitionInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewReviewer
}

type ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewReviewer
}

type ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions struct {
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

func DefaultListAccessReviewDefinitionInstanceContactedReviewersOperationOptions() ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions {
	return ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceContactedReviewers - List contactedReviewers. Get the reviewers for an access
// review instance, irrespective of whether or not they have received a notification. The reviewers are represented by
// an accessReviewReviewer object. A list of zero or more objects are returned, including all of their nested
// properties.
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewers(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions) (result ListAccessReviewDefinitionInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceContactedReviewersCustomPager{},
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

// ListAccessReviewDefinitionInstanceContactedReviewersComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewersComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions) (ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, options, AccessReviewReviewerOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceContactedReviewersOperationOptions, predicate AccessReviewReviewerOperationPredicate) (result ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult, err error) {
	items := make([]beta.AccessReviewReviewer, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceContactedReviewers(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
