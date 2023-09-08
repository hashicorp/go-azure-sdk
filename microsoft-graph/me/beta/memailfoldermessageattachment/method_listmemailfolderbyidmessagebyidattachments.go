package memailfoldermessageattachment

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

type ListMeMailFolderByIdMessageByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeMailFolderByIdMessageByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeMailFolderByIdMessageByIdAttachments ...
func (c MeMailFolderMessageAttachmentClient) ListMeMailFolderByIdMessageByIdAttachments(ctx context.Context, id MeMailFolderMessageId) (result ListMeMailFolderByIdMessageByIdAttachmentsOperationResponse, err error) {
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

// ListMeMailFolderByIdMessageByIdAttachmentsComplete retrieves all the results into a single object
func (c MeMailFolderMessageAttachmentClient) ListMeMailFolderByIdMessageByIdAttachmentsComplete(ctx context.Context, id MeMailFolderMessageId) (ListMeMailFolderByIdMessageByIdAttachmentsCompleteResult, error) {
	return c.ListMeMailFolderByIdMessageByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdMessageByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderMessageAttachmentClient) ListMeMailFolderByIdMessageByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeMailFolderMessageId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeMailFolderByIdMessageByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdMessageByIdAttachments(ctx, id)
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

	result = ListMeMailFolderByIdMessageByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
