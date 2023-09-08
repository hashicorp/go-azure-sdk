package userauthenticationphonemethod

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

type ListUserByIdAuthenticationPhoneMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PhoneAuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationPhoneMethodsCompleteResult struct {
	Items []models.PhoneAuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationPhoneMethods ...
func (c UserAuthenticationPhoneMethodClient) ListUserByIdAuthenticationPhoneMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationPhoneMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/phoneMethods", id.ID()),
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
		Values *[]models.PhoneAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationPhoneMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationPhoneMethodClient) ListUserByIdAuthenticationPhoneMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationPhoneMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx, id, models.PhoneAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationPhoneMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationPhoneMethodClient) ListUserByIdAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PhoneAuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationPhoneMethodsCompleteResult, err error) {
	items := make([]models.PhoneAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationPhoneMethods(ctx, id)
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

	result = ListUserByIdAuthenticationPhoneMethodsCompleteResult{
		Items: items,
	}
	return
}
