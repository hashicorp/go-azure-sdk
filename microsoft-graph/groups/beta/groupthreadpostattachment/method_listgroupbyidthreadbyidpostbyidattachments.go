package groupthreadpostattachment

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

type ListGroupByIdThreadByIdPostByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListGroupByIdThreadByIdPostByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListGroupByIdThreadByIdPostByIdAttachments ...
func (c GroupThreadPostAttachmentClient) ListGroupByIdThreadByIdPostByIdAttachments(ctx context.Context, id GroupThreadPostId) (result ListGroupByIdThreadByIdPostByIdAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attachments", id.ID()),
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

// ListGroupByIdThreadByIdPostByIdAttachmentsComplete retrieves all the results into a single object
func (c GroupThreadPostAttachmentClient) ListGroupByIdThreadByIdPostByIdAttachmentsComplete(ctx context.Context, id GroupThreadPostId) (ListGroupByIdThreadByIdPostByIdAttachmentsCompleteResult, error) {
	return c.ListGroupByIdThreadByIdPostByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadByIdPostByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadPostAttachmentClient) ListGroupByIdThreadByIdPostByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupThreadPostId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListGroupByIdThreadByIdPostByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreadByIdPostByIdAttachments(ctx, id)
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

	result = ListGroupByIdThreadByIdPostByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
