package threadpost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListThreadPostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Post
}

type ListThreadPostsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Post
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

// ListThreadPosts ...
func (c ThreadPostClient) ListThreadPosts(ctx context.Context, id GroupIdThreadId) (result ListThreadPostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListThreadPostsCustomPager{},
		Path:       fmt.Sprintf("%s/posts", id.ID()),
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
		Values *[]stable.Post `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListThreadPostsComplete retrieves all the results into a single object
func (c ThreadPostClient) ListThreadPostsComplete(ctx context.Context, id GroupIdThreadId) (ListThreadPostsCompleteResult, error) {
	return c.ListThreadPostsCompleteMatchingPredicate(ctx, id, PostOperationPredicate{})
}

// ListThreadPostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostClient) ListThreadPostsCompleteMatchingPredicate(ctx context.Context, id GroupIdThreadId, predicate PostOperationPredicate) (result ListThreadPostsCompleteResult, err error) {
	items := make([]stable.Post, 0)

	resp, err := c.ListThreadPosts(ctx, id)
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
