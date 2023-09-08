package user

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

type ListUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserCollectionResponse
}

type ListUsersCompleteResult struct {
	Items []models.UserCollectionResponse
}

// ListUsers ...
func (c UserClient) ListUsers(ctx context.Context) (result ListUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/users",
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
		Values *[]models.UserCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUsersComplete retrieves all the results into a single object
func (c UserClient) ListUsersComplete(ctx context.Context) (ListUsersCompleteResult, error) {
	return c.ListUsersCompleteMatchingPredicate(ctx, models.UserCollectionResponseOperationPredicate{})
}

// ListUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserClient) ListUsersCompleteMatchingPredicate(ctx context.Context, predicate models.UserCollectionResponseOperationPredicate) (result ListUsersCompleteResult, err error) {
	items := make([]models.UserCollectionResponse, 0)

	resp, err := c.ListUsers(ctx)
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

	result = ListUsersCompleteResult{
		Items: items,
	}
	return
}
