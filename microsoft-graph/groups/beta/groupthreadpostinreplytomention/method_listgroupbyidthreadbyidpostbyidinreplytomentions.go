package groupthreadpostinreplytomention

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

type ListGroupByIdThreadByIdPostByIdInReplyToMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MentionCollectionResponse
}

type ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteResult struct {
	Items []models.MentionCollectionResponse
}

// ListGroupByIdThreadByIdPostByIdInReplyToMentions ...
func (c GroupThreadPostInReplyToMentionClient) ListGroupByIdThreadByIdPostByIdInReplyToMentions(ctx context.Context, id GroupThreadPostId) (result ListGroupByIdThreadByIdPostByIdInReplyToMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inReplyTo/mentions", id.ID()),
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
		Values *[]models.MentionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdThreadByIdPostByIdInReplyToMentionsComplete retrieves all the results into a single object
func (c GroupThreadPostInReplyToMentionClient) ListGroupByIdThreadByIdPostByIdInReplyToMentionsComplete(ctx context.Context, id GroupThreadPostId) (ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteResult, error) {
	return c.ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteMatchingPredicate(ctx, id, models.MentionCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadPostInReplyToMentionClient) ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteMatchingPredicate(ctx context.Context, id GroupThreadPostId, predicate models.MentionCollectionResponseOperationPredicate) (result ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteResult, err error) {
	items := make([]models.MentionCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreadByIdPostByIdInReplyToMentions(ctx, id)
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

	result = ListGroupByIdThreadByIdPostByIdInReplyToMentionsCompleteResult{
		Items: items,
	}
	return
}
