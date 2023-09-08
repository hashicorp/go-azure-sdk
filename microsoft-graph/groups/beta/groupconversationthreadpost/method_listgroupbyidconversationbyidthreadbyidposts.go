package groupconversationthreadpost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdConversationByIdThreadByIdPostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PostCollectionResponse
}

type ListGroupByIdConversationByIdThreadByIdPostsCompleteResult struct {
	Items []models.PostCollectionResponse
}

// ListGroupByIdConversationByIdThreadByIdPosts ...
func (c GroupConversationThreadPostClient) ListGroupByIdConversationByIdThreadByIdPosts(ctx context.Context, id GroupConversationThreadId) (result ListGroupByIdConversationByIdThreadByIdPostsOperationResponse, err error) {
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

// ListGroupByIdConversationByIdThreadByIdPostsComplete retrieves all the results into a single object
func (c GroupConversationThreadPostClient) ListGroupByIdConversationByIdThreadByIdPostsComplete(ctx context.Context, id GroupConversationThreadId) (ListGroupByIdConversationByIdThreadByIdPostsCompleteResult, error) {
	return c.ListGroupByIdConversationByIdThreadByIdPostsCompleteMatchingPredicate(ctx, id, models.PostCollectionResponseOperationPredicate{})
}

// ListGroupByIdConversationByIdThreadByIdPostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupConversationThreadPostClient) ListGroupByIdConversationByIdThreadByIdPostsCompleteMatchingPredicate(ctx context.Context, id GroupConversationThreadId, predicate models.PostCollectionResponseOperationPredicate) (result ListGroupByIdConversationByIdThreadByIdPostsCompleteResult, err error) {
	items := make([]models.PostCollectionResponse, 0)

	resp, err := c.ListGroupByIdConversationByIdThreadByIdPosts(ctx, id)
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

	result = ListGroupByIdConversationByIdThreadByIdPostsCompleteResult{
		Items: items,
	}
	return
}
