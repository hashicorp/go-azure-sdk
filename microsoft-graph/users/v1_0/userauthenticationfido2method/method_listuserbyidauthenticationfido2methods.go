package userauthenticationfido2method

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdAuthenticationFido2MethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.Fido2AuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationFido2MethodsCompleteResult struct {
	Items []models.Fido2AuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationFido2Methods ...
func (c UserAuthenticationFido2MethodClient) ListUserByIdAuthenticationFido2Methods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationFido2MethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/fido2Methods", id.ID()),
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

// ListUserByIdAuthenticationFido2MethodsComplete retrieves all the results into a single object
func (c UserAuthenticationFido2MethodClient) ListUserByIdAuthenticationFido2MethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationFido2MethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationFido2MethodsCompleteMatchingPredicate(ctx, id, models.Fido2AuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationFido2MethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationFido2MethodClient) ListUserByIdAuthenticationFido2MethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.Fido2AuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationFido2MethodsCompleteResult, err error) {
	items := make([]models.Fido2AuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationFido2Methods(ctx, id)
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

	result = ListUserByIdAuthenticationFido2MethodsCompleteResult{
		Items: items,
	}
	return
}
