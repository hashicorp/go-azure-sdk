package meauthenticationpasswordlessmicrosoftauthenticatormethod

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

type ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult struct {
	Items []models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethods ...
func (c MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx context.Context) (result ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/passwordlessMicrosoftAuthenticatorMethods",
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
		Values *[]models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsComplete(ctx context.Context) (ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]models.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx)
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

	result = ListMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult{
		Items: items,
	}
	return
}
