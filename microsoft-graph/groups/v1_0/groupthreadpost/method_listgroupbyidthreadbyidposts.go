package groupthreadpost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdThreadByIdPostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PostCollectionResponse
}

type ListGroupByIdThreadByIdPostsCompleteResult struct {
	Items []models.PostCollectionResponse
}

// ListGroupByIdThreadByIdPosts ...
func (c GroupThreadPostClient) ListGroupByIdThreadByIdPosts(ctx context.Context, id GroupThreadId) (result ListGroupByIdThreadByIdPostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.PostCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdThreadByIdPostsComplete retrieves all the results into a single object
func (c GroupThreadPostClient) ListGroupByIdThreadByIdPostsComplete(ctx context.Context, id GroupThreadId) (ListGroupByIdThreadByIdPostsCompleteResult, error) {
	return c.ListGroupByIdThreadByIdPostsCompleteMatchingPredicate(ctx, id, models.PostCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadByIdPostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadPostClient) ListGroupByIdThreadByIdPostsCompleteMatchingPredicate(ctx context.Context, id GroupThreadId, predicate models.PostCollectionResponseOperationPredicate) (result ListGroupByIdThreadByIdPostsCompleteResult, err error) {
	items := make([]models.PostCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreadByIdPosts(ctx, id)
	if err != nil {
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

	result = ListGroupByIdThreadByIdPostsCompleteResult{
		Items: items,
	}
	return
}
