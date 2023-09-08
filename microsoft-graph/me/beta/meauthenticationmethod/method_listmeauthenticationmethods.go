package meauthenticationmethod

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

type ListMeAuthenticationMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationMethodCollectionResponse
}

type ListMeAuthenticationMethodsCompleteResult struct {
	Items []models.AuthenticationMethodCollectionResponse
}

// ListMeAuthenticationMethods ...
func (c MeAuthenticationMethodClient) ListMeAuthenticationMethods(ctx context.Context) (result ListMeAuthenticationMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/methods",
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

// ListMeAuthenticationMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationMethodClient) ListMeAuthenticationMethodsComplete(ctx context.Context) (ListMeAuthenticationMethodsCompleteResult, error) {
	return c.ListMeAuthenticationMethodsCompleteMatchingPredicate(ctx, models.AuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationMethodClient) ListMeAuthenticationMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.AuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationMethodsCompleteResult, err error) {
	items := make([]models.AuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationMethods(ctx)
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

	result = ListMeAuthenticationMethodsCompleteResult{
		Items: items,
	}
	return
}
