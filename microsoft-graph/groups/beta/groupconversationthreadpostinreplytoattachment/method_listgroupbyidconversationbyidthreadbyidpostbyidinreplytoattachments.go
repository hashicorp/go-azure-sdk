package groupconversationthreadpostinreplytoattachment

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

type ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachments ...
func (c GroupConversationThreadPostInReplyToAttachmentClient) ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachments(ctx context.Context, id GroupConversationThreadPostId) (result ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inReplyTo/attachments", id.ID()),
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
		Values *[]models.AttachmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsComplete retrieves all the results into a single object
func (c GroupConversationThreadPostInReplyToAttachmentClient) ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsComplete(ctx context.Context, id GroupConversationThreadPostId) (ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteResult, error) {
	return c.ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupConversationThreadPostInReplyToAttachmentClient) ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupConversationThreadPostId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachments(ctx, id)
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

	result = ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsCompleteResult{
		Items: items,
	}
	return
}
