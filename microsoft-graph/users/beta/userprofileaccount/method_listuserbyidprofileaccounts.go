package userprofileaccount

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

type ListUserByIdProfileAccountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserAccountInformationCollectionResponse
}

type ListUserByIdProfileAccountsCompleteResult struct {
	Items []models.UserAccountInformationCollectionResponse
}

// ListUserByIdProfileAccounts ...
func (c UserProfileAccountClient) ListUserByIdProfileAccounts(ctx context.Context, id UserId) (result ListUserByIdProfileAccountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/account", id.ID()),
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
		Values *[]models.UserAccountInformationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileAccountsComplete retrieves all the results into a single object
func (c UserProfileAccountClient) ListUserByIdProfileAccountsComplete(ctx context.Context, id UserId) (ListUserByIdProfileAccountsCompleteResult, error) {
	return c.ListUserByIdProfileAccountsCompleteMatchingPredicate(ctx, id, models.UserAccountInformationCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileAccountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileAccountClient) ListUserByIdProfileAccountsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.UserAccountInformationCollectionResponseOperationPredicate) (result ListUserByIdProfileAccountsCompleteResult, err error) {
	items := make([]models.UserAccountInformationCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileAccounts(ctx, id)
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

	result = ListUserByIdProfileAccountsCompleteResult{
		Items: items,
	}
	return
}
