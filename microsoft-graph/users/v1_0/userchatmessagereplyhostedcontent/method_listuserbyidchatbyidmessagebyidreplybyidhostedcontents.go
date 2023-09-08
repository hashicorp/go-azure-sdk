package userchatmessagereplyhostedcontent

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

type ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageHostedContentCollectionResponse
}

type ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteResult struct {
	Items []models.ChatMessageHostedContentCollectionResponse
}

// ListUserByIdChatByIdMessageByIdReplyByIdHostedContents ...
func (c UserChatMessageReplyHostedContentClient) ListUserByIdChatByIdMessageByIdReplyByIdHostedContents(ctx context.Context, id UserChatMessageReplyId) (result ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]models.ChatMessageHostedContentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsComplete retrieves all the results into a single object
func (c UserChatMessageReplyHostedContentClient) ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsComplete(ctx context.Context, id UserChatMessageReplyId) (ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteResult, error) {
	return c.ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx, id, models.ChatMessageHostedContentCollectionResponseOperationPredicate{})
}

// ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserChatMessageReplyHostedContentClient) ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx context.Context, id UserChatMessageReplyId, predicate models.ChatMessageHostedContentCollectionResponseOperationPredicate) (result ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteResult, err error) {
	items := make([]models.ChatMessageHostedContentCollectionResponse, 0)

	resp, err := c.ListUserByIdChatByIdMessageByIdReplyByIdHostedContents(ctx, id)
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

	result = ListUserByIdChatByIdMessageByIdReplyByIdHostedContentsCompleteResult{
		Items: items,
	}
	return
}
