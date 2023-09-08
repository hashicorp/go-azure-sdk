package meauthenticationmicrosoftauthenticatormethod

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

type ListMeAuthenticationMicrosoftAuthenticatorMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteResult struct {
	Items []models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationMicrosoftAuthenticatorMethods ...
func (c MeAuthenticationMicrosoftAuthenticatorMethodClient) ListMeAuthenticationMicrosoftAuthenticatorMethods(ctx context.Context) (result ListMeAuthenticationMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/microsoftAuthenticatorMethods",
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

// ListMeAuthenticationMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationMicrosoftAuthenticatorMethodClient) ListMeAuthenticationMicrosoftAuthenticatorMethodsComplete(ctx context.Context) (ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationMicrosoftAuthenticatorMethodClient) ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]models.MicrosoftAuthenticatorAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationMicrosoftAuthenticatorMethods(ctx)
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

	result = ListMeAuthenticationMicrosoftAuthenticatorMethodsCompleteResult{
		Items: items,
	}
	return
}
