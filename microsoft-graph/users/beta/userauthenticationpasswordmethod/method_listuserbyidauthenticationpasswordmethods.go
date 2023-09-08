package userauthenticationpasswordmethod

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

type ListUserByIdAuthenticationPasswordMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PasswordAuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationPasswordMethodsCompleteResult struct {
	Items []models.PasswordAuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationPasswordMethods ...
func (c UserAuthenticationPasswordMethodClient) ListUserByIdAuthenticationPasswordMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationPasswordMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/passwordMethods", id.ID()),
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
		Values *[]models.PasswordAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationPasswordMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationPasswordMethodClient) ListUserByIdAuthenticationPasswordMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationPasswordMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx, id, models.PasswordAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationPasswordMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationPasswordMethodClient) ListUserByIdAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PasswordAuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationPasswordMethodsCompleteResult, err error) {
	items := make([]models.PasswordAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationPasswordMethods(ctx, id)
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

	result = ListUserByIdAuthenticationPasswordMethodsCompleteResult{
		Items: items,
	}
	return
}
