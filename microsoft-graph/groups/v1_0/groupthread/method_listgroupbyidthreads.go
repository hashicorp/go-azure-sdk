package groupthread

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

type ListGroupByIdThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationThreadCollectionResponse
}

type ListGroupByIdThreadsCompleteResult struct {
	Items []models.ConversationThreadCollectionResponse
}

// ListGroupByIdThreads ...
func (c GroupThreadClient) ListGroupByIdThreads(ctx context.Context, id GroupId) (result ListGroupByIdThreadsOperationResponse, err error) {
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

// ListGroupByIdThreadsComplete retrieves all the results into a single object
func (c GroupThreadClient) ListGroupByIdThreadsComplete(ctx context.Context, id GroupId) (ListGroupByIdThreadsCompleteResult, error) {
	return c.ListGroupByIdThreadsCompleteMatchingPredicate(ctx, id, models.ConversationThreadCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadClient) ListGroupByIdThreadsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ConversationThreadCollectionResponseOperationPredicate) (result ListGroupByIdThreadsCompleteResult, err error) {
	items := make([]models.ConversationThreadCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreads(ctx, id)
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

	result = ListGroupByIdThreadsCompleteResult{
		Items: items,
	}
	return
}
