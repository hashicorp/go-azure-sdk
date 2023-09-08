package mechat

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

type ListMeChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatCollectionResponse
}

type ListMeChatsCompleteResult struct {
	Items []models.ChatCollectionResponse
}

// ListMeChats ...
func (c MeChatClient) ListMeChats(ctx context.Context) (result ListMeChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/chats",
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

// ListMeChatsComplete retrieves all the results into a single object
func (c MeChatClient) ListMeChatsComplete(ctx context.Context) (ListMeChatsCompleteResult, error) {
	return c.ListMeChatsCompleteMatchingPredicate(ctx, models.ChatCollectionResponseOperationPredicate{})
}

// ListMeChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatClient) ListMeChatsCompleteMatchingPredicate(ctx context.Context, predicate models.ChatCollectionResponseOperationPredicate) (result ListMeChatsCompleteResult, err error) {
	items := make([]models.ChatCollectionResponse, 0)

	resp, err := c.ListMeChats(ctx)
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

	result = ListMeChatsCompleteResult{
		Items: items,
	}
	return
}
