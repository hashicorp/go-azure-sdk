package mechatmessagereplyhostedcontent

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

type ListMeChatByIdMessageByIdReplyByIdHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageHostedContentCollectionResponse
}

type ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteResult struct {
	Items []models.ChatMessageHostedContentCollectionResponse
}

// ListMeChatByIdMessageByIdReplyByIdHostedContents ...
func (c MeChatMessageReplyHostedContentClient) ListMeChatByIdMessageByIdReplyByIdHostedContents(ctx context.Context, id MeChatMessageReplyId) (result ListMeChatByIdMessageByIdReplyByIdHostedContentsOperationResponse, err error) {
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

// ListMeChatByIdMessageByIdReplyByIdHostedContentsComplete retrieves all the results into a single object
func (c MeChatMessageReplyHostedContentClient) ListMeChatByIdMessageByIdReplyByIdHostedContentsComplete(ctx context.Context, id MeChatMessageReplyId) (ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteResult, error) {
	return c.ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx, id, models.ChatMessageHostedContentCollectionResponseOperationPredicate{})
}

// ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatMessageReplyHostedContentClient) ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx context.Context, id MeChatMessageReplyId, predicate models.ChatMessageHostedContentCollectionResponseOperationPredicate) (result ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteResult, err error) {
	items := make([]models.ChatMessageHostedContentCollectionResponse, 0)

	resp, err := c.ListMeChatByIdMessageByIdReplyByIdHostedContents(ctx, id)
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

	result = ListMeChatByIdMessageByIdReplyByIdHostedContentsCompleteResult{
		Items: items,
	}
	return
}
