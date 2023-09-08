package usermailfoldermessage

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

type ListUserByIdMailFolderByIdMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MessageCollectionResponse
}

type ListUserByIdMailFolderByIdMessagesCompleteResult struct {
	Items []models.MessageCollectionResponse
}

// ListUserByIdMailFolderByIdMessages ...
func (c UserMailFolderMessageClient) ListUserByIdMailFolderByIdMessages(ctx context.Context, id UserMailFolderId) (result ListUserByIdMailFolderByIdMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/messages", id.ID()),
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
		Values *[]models.MessageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdMailFolderByIdMessagesComplete retrieves all the results into a single object
func (c UserMailFolderMessageClient) ListUserByIdMailFolderByIdMessagesComplete(ctx context.Context, id UserMailFolderId) (ListUserByIdMailFolderByIdMessagesCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdMessagesCompleteMatchingPredicate(ctx, id, models.MessageCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderMessageClient) ListUserByIdMailFolderByIdMessagesCompleteMatchingPredicate(ctx context.Context, id UserMailFolderId, predicate models.MessageCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdMessagesCompleteResult, err error) {
	items := make([]models.MessageCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdMessages(ctx, id)
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

	result = ListUserByIdMailFolderByIdMessagesCompleteResult{
		Items: items,
	}
	return
}
