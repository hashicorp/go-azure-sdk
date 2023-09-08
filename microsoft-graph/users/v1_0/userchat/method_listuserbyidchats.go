package userchat

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

type ListUserByIdChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatCollectionResponse
}

type ListUserByIdChatsCompleteResult struct {
	Items []models.ChatCollectionResponse
}

// ListUserByIdChats ...
func (c UserChatClient) ListUserByIdChats(ctx context.Context, id UserId) (result ListUserByIdChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/chats", id.ID()),
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
		Values *[]models.ChatCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdChatsComplete retrieves all the results into a single object
func (c UserChatClient) ListUserByIdChatsComplete(ctx context.Context, id UserId) (ListUserByIdChatsCompleteResult, error) {
	return c.ListUserByIdChatsCompleteMatchingPredicate(ctx, id, models.ChatCollectionResponseOperationPredicate{})
}

// ListUserByIdChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserChatClient) ListUserByIdChatsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ChatCollectionResponseOperationPredicate) (result ListUserByIdChatsCompleteResult, err error) {
	items := make([]models.ChatCollectionResponse, 0)

	resp, err := c.ListUserByIdChats(ctx, id)
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

	result = ListUserByIdChatsCompleteResult{
		Items: items,
	}
	return
}
