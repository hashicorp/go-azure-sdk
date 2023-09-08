package memessage

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

type ListMeMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MessageCollectionResponse
}

type ListMeMessagesCompleteResult struct {
	Items []models.MessageCollectionResponse
}

// ListMeMessages ...
func (c MeMessageClient) ListMeMessages(ctx context.Context) (result ListMeMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/messages",
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

// ListMeMessagesComplete retrieves all the results into a single object
func (c MeMessageClient) ListMeMessagesComplete(ctx context.Context) (ListMeMessagesCompleteResult, error) {
	return c.ListMeMessagesCompleteMatchingPredicate(ctx, models.MessageCollectionResponseOperationPredicate{})
}

// ListMeMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMessageClient) ListMeMessagesCompleteMatchingPredicate(ctx context.Context, predicate models.MessageCollectionResponseOperationPredicate) (result ListMeMessagesCompleteResult, err error) {
	items := make([]models.MessageCollectionResponse, 0)

	resp, err := c.ListMeMessages(ctx)
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

	result = ListMeMessagesCompleteResult{
		Items: items,
	}
	return
}
