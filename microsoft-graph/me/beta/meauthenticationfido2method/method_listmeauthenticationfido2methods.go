package meauthenticationfido2method

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

type ListMeAuthenticationFido2MethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.Fido2AuthenticationMethodCollectionResponse
}

type ListMeAuthenticationFido2MethodsCompleteResult struct {
	Items []models.Fido2AuthenticationMethodCollectionResponse
}

// ListMeAuthenticationFido2Methods ...
func (c MeAuthenticationFido2MethodClient) ListMeAuthenticationFido2Methods(ctx context.Context) (result ListMeAuthenticationFido2MethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/fido2Methods",
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
		Values *[]models.Fido2AuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeAuthenticationFido2MethodsComplete retrieves all the results into a single object
func (c MeAuthenticationFido2MethodClient) ListMeAuthenticationFido2MethodsComplete(ctx context.Context) (ListMeAuthenticationFido2MethodsCompleteResult, error) {
	return c.ListMeAuthenticationFido2MethodsCompleteMatchingPredicate(ctx, models.Fido2AuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationFido2MethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationFido2MethodClient) ListMeAuthenticationFido2MethodsCompleteMatchingPredicate(ctx context.Context, predicate models.Fido2AuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationFido2MethodsCompleteResult, err error) {
	items := make([]models.Fido2AuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationFido2Methods(ctx)
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

	result = ListMeAuthenticationFido2MethodsCompleteResult{
		Items: items,
	}
	return
}
