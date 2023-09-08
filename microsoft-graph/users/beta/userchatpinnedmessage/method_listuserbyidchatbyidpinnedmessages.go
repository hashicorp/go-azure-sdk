package userchatpinnedmessage

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

type ListUserByIdChatByIdPinnedMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PinnedChatMessageInfoCollectionResponse
}

type ListUserByIdChatByIdPinnedMessagesCompleteResult struct {
	Items []models.PinnedChatMessageInfoCollectionResponse
}

// ListUserByIdChatByIdPinnedMessages ...
func (c UserChatPinnedMessageClient) ListUserByIdChatByIdPinnedMessages(ctx context.Context, id UserChatId) (result ListUserByIdChatByIdPinnedMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/pinnedMessages", id.ID()),
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
		Values *[]models.PinnedChatMessageInfoCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdChatByIdPinnedMessagesComplete retrieves all the results into a single object
func (c UserChatPinnedMessageClient) ListUserByIdChatByIdPinnedMessagesComplete(ctx context.Context, id UserChatId) (ListUserByIdChatByIdPinnedMessagesCompleteResult, error) {
	return c.ListUserByIdChatByIdPinnedMessagesCompleteMatchingPredicate(ctx, id, models.PinnedChatMessageInfoCollectionResponseOperationPredicate{})
}

// ListUserByIdChatByIdPinnedMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserChatPinnedMessageClient) ListUserByIdChatByIdPinnedMessagesCompleteMatchingPredicate(ctx context.Context, id UserChatId, predicate models.PinnedChatMessageInfoCollectionResponseOperationPredicate) (result ListUserByIdChatByIdPinnedMessagesCompleteResult, err error) {
	items := make([]models.PinnedChatMessageInfoCollectionResponse, 0)

	resp, err := c.ListUserByIdChatByIdPinnedMessages(ctx, id)
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

	result = ListUserByIdChatByIdPinnedMessagesCompleteResult{
		Items: items,
	}
	return
}
