package meprofilewebaccount

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

type ListMeProfileWebAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WebAccountCollectionResponse
}

type ListMeProfileWebAccountsCompleteResult struct {
	Items []models.WebAccountCollectionResponse
}

// ListMeProfileWebAccounts ...
func (c MeProfileWebAccountClient) ListMeProfileWebAccounts(ctx context.Context) (result ListMeProfileWebAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/profile/webAccounts",
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

// ListMeProfileWebAccountsComplete retrieves all the results into a single object
func (c MeProfileWebAccountClient) ListMeProfileWebAccountsComplete(ctx context.Context) (ListMeProfileWebAccountsCompleteResult, error) {
	return c.ListMeProfileWebAccountsCompleteMatchingPredicate(ctx, models.WebAccountCollectionResponseOperationPredicate{})
}

// ListMeProfileWebAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileWebAccountClient) ListMeProfileWebAccountsCompleteMatchingPredicate(ctx context.Context, predicate models.WebAccountCollectionResponseOperationPredicate) (result ListMeProfileWebAccountsCompleteResult, err error) {
	items := make([]models.WebAccountCollectionResponse, 0)

	resp, err := c.ListMeProfileWebAccounts(ctx)
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

	result = ListMeProfileWebAccountsCompleteResult{
		Items: items,
	}
	return
}
