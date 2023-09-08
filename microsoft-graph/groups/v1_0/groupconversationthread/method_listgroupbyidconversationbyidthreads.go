package groupconversationthread

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

type ListGroupByIdConversationByIdThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationThreadCollectionResponse
}

type ListGroupByIdConversationByIdThreadsCompleteResult struct {
	Items []models.ConversationThreadCollectionResponse
}

// ListGroupByIdConversationByIdThreads ...
func (c GroupConversationThreadClient) ListGroupByIdConversationByIdThreads(ctx context.Context, id GroupConversationId) (result ListGroupByIdConversationByIdThreadsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/threads", id.ID()),
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
		Values *[]models.ConversationThreadCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdConversationByIdThreadsComplete retrieves all the results into a single object
func (c GroupConversationThreadClient) ListGroupByIdConversationByIdThreadsComplete(ctx context.Context, id GroupConversationId) (ListGroupByIdConversationByIdThreadsCompleteResult, error) {
	return c.ListGroupByIdConversationByIdThreadsCompleteMatchingPredicate(ctx, id, models.ConversationThreadCollectionResponseOperationPredicate{})
}

// ListGroupByIdConversationByIdThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupConversationThreadClient) ListGroupByIdConversationByIdThreadsCompleteMatchingPredicate(ctx context.Context, id GroupConversationId, predicate models.ConversationThreadCollectionResponseOperationPredicate) (result ListGroupByIdConversationByIdThreadsCompleteResult, err error) {
	items := make([]models.ConversationThreadCollectionResponse, 0)

	resp, err := c.ListGroupByIdConversationByIdThreads(ctx, id)
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

	result = ListGroupByIdConversationByIdThreadsCompleteResult{
		Items: items,
	}
	return
}
