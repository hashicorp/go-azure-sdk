package usertodolisttaskattachment

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

type ListUserByIdTodoListByIdTaskByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentBaseCollectionResponse
}

type ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentBaseCollectionResponse
}

// ListUserByIdTodoListByIdTaskByIdAttachments ...
func (c UserTodoListTaskAttachmentClient) ListUserByIdTodoListByIdTaskByIdAttachments(ctx context.Context, id UserTodoListTaskId) (result ListUserByIdTodoListByIdTaskByIdAttachmentsOperationResponse, err error) {
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
		Values *[]models.AttachmentBaseCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTodoListByIdTaskByIdAttachmentsComplete retrieves all the results into a single object
func (c UserTodoListTaskAttachmentClient) ListUserByIdTodoListByIdTaskByIdAttachmentsComplete(ctx context.Context, id UserTodoListTaskId) (ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentBaseCollectionResponseOperationPredicate{})
}

// ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTodoListTaskAttachmentClient) ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserTodoListTaskId, predicate models.AttachmentBaseCollectionResponseOperationPredicate) (result ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentBaseCollectionResponse, 0)

	resp, err := c.ListUserByIdTodoListByIdTaskByIdAttachments(ctx, id)
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

	result = ListUserByIdTodoListByIdTaskByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
