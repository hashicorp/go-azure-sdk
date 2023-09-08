package groupconversation

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

type ListGroupByIdConversationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationCollectionResponse
}

type ListGroupByIdConversationsCompleteResult struct {
	Items []models.ConversationCollectionResponse
}

// ListGroupByIdConversations ...
func (c GroupConversationClient) ListGroupByIdConversations(ctx context.Context, id GroupId) (result ListGroupByIdConversationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/conversations", id.ID()),
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
		Values *[]models.ConversationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdConversationsComplete retrieves all the results into a single object
func (c GroupConversationClient) ListGroupByIdConversationsComplete(ctx context.Context, id GroupId) (ListGroupByIdConversationsCompleteResult, error) {
	return c.ListGroupByIdConversationsCompleteMatchingPredicate(ctx, id, models.ConversationCollectionResponseOperationPredicate{})
}

// ListGroupByIdConversationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupConversationClient) ListGroupByIdConversationsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ConversationCollectionResponseOperationPredicate) (result ListGroupByIdConversationsCompleteResult, err error) {
	items := make([]models.ConversationCollectionResponse, 0)

	resp, err := c.ListGroupByIdConversations(ctx, id)
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

	result = ListGroupByIdConversationsCompleteResult{
		Items: items,
	}
	return
}
