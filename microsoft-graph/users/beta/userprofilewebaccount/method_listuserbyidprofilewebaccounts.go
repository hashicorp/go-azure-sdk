package userprofilewebaccount

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

type ListUserByIdProfileWebAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WebAccountCollectionResponse
}

type ListUserByIdProfileWebAccountsCompleteResult struct {
	Items []models.WebAccountCollectionResponse
}

// ListUserByIdProfileWebAccounts ...
func (c UserProfileWebAccountClient) ListUserByIdProfileWebAccounts(ctx context.Context, id UserId) (result ListUserByIdProfileWebAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/webAccounts", id.ID()),
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
		Values *[]models.WebAccountCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileWebAccountsComplete retrieves all the results into a single object
func (c UserProfileWebAccountClient) ListUserByIdProfileWebAccountsComplete(ctx context.Context, id UserId) (ListUserByIdProfileWebAccountsCompleteResult, error) {
	return c.ListUserByIdProfileWebAccountsCompleteMatchingPredicate(ctx, id, models.WebAccountCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileWebAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileWebAccountClient) ListUserByIdProfileWebAccountsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.WebAccountCollectionResponseOperationPredicate) (result ListUserByIdProfileWebAccountsCompleteResult, err error) {
	items := make([]models.WebAccountCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileWebAccounts(ctx, id)
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

	result = ListUserByIdProfileWebAccountsCompleteResult{
		Items: items,
	}
	return
}
