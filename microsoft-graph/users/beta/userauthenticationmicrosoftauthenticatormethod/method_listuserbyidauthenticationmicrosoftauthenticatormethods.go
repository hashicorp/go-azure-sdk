package userauthenticationmicrosoftauthenticatormethod

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

type ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteResult struct {
	Items []models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationMicrosoftAuthenticatorMethods ...
func (c UserAuthenticationMicrosoftAuthenticatorMethodClient) ListUserByIdAuthenticationMicrosoftAuthenticatorMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/microsoftAuthenticatorMethods", id.ID()),
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
		Values *[]models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationMicrosoftAuthenticatorMethodClient) ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, id, models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationMicrosoftAuthenticatorMethodClient) ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationMicrosoftAuthenticatorMethods(ctx, id)
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

	result = ListUserByIdAuthenticationMicrosoftAuthenticatorMethodsCompleteResult{
		Items: items,
	}
	return
}
