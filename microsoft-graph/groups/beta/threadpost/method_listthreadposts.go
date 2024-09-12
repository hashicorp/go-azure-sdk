package threadpost

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

type ListThreadPostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Post
}

type ListThreadPostsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Post
}

type ListThreadPostsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListThreadPostsOperationOptions() ListThreadPostsOperationOptions {
	return ListThreadPostsOperationOptions{}
}

func (o ListThreadPostsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListThreadPostsOperationOptions) ToOData() *odata.Query {
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

func (o ListThreadPostsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListThreadPostsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadPostsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreadPosts - List posts. Get the posts of the specified thread. You can specify both the parent conversation and
// the thread, or, you can specify the thread without referencing the parent conversation.
func (c ThreadPostClient) ListThreadPosts(ctx context.Context, id beta.GroupIdThreadId, options ListThreadPostsOperationOptions) (result ListThreadPostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListThreadPostsCustomPager{},
		Path:          fmt.Sprintf("%s/posts", id.ID()),
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
		Values *[]beta.Post `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListThreadPostsComplete retrieves all the results into a single object
func (c ThreadPostClient) ListThreadPostsComplete(ctx context.Context, id beta.GroupIdThreadId, options ListThreadPostsOperationOptions) (ListThreadPostsCompleteResult, error) {
	return c.ListThreadPostsCompleteMatchingPredicate(ctx, id, options, PostOperationPredicate{})
}

// ListThreadPostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostClient) ListThreadPostsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdThreadId, options ListThreadPostsOperationOptions, predicate PostOperationPredicate) (result ListThreadPostsCompleteResult, err error) {
	items := make([]beta.Post, 0)

	resp, err := c.ListThreadPosts(ctx, id, options)
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

	result = ListThreadPostsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
