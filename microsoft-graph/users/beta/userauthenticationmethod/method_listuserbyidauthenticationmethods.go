package userauthenticationmethod

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

type ListUserByIdAuthenticationMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationMethodsCompleteResult struct {
	Items []models.AuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationMethods ...
func (c UserAuthenticationMethodClient) ListUserByIdAuthenticationMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/methods", id.ID()),
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
		Values *[]models.AuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationMethodClient) ListUserByIdAuthenticationMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationMethodsCompleteMatchingPredicate(ctx, id, models.AuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationMethodClient) ListUserByIdAuthenticationMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.AuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationMethodsCompleteResult, err error) {
	items := make([]models.AuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationMethods(ctx, id)
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

	result = ListUserByIdAuthenticationMethodsCompleteResult{
		Items: items,
	}
	return
}
