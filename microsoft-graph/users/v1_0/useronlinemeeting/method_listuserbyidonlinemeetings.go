package useronlinemeeting

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

type ListUserByIdOnlineMeetingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnlineMeetingCollectionResponse
}

type ListUserByIdOnlineMeetingsCompleteResult struct {
	Items []models.OnlineMeetingCollectionResponse
}

// ListUserByIdOnlineMeetings ...
func (c UserOnlineMeetingClient) ListUserByIdOnlineMeetings(ctx context.Context, id UserId) (result ListUserByIdOnlineMeetingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onlineMeetings", id.ID()),
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
		Values *[]models.OnlineMeetingCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnlineMeetingsComplete retrieves all the results into a single object
func (c UserOnlineMeetingClient) ListUserByIdOnlineMeetingsComplete(ctx context.Context, id UserId) (ListUserByIdOnlineMeetingsCompleteResult, error) {
	return c.ListUserByIdOnlineMeetingsCompleteMatchingPredicate(ctx, id, models.OnlineMeetingCollectionResponseOperationPredicate{})
}

// ListUserByIdOnlineMeetingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnlineMeetingClient) ListUserByIdOnlineMeetingsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OnlineMeetingCollectionResponseOperationPredicate) (result ListUserByIdOnlineMeetingsCompleteResult, err error) {
	items := make([]models.OnlineMeetingCollectionResponse, 0)

	resp, err := c.ListUserByIdOnlineMeetings(ctx, id)
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

	result = ListUserByIdOnlineMeetingsCompleteResult{
		Items: items,
	}
	return
}
