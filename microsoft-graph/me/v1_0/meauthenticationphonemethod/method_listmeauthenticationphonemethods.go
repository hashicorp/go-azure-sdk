package meauthenticationphonemethod

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

type ListMeAuthenticationPhoneMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PhoneAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationPhoneMethodsCompleteResult struct {
	Items []models.PhoneAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationPhoneMethods ...
func (c MeAuthenticationPhoneMethodClient) ListMeAuthenticationPhoneMethods(ctx context.Context) (result ListMeAuthenticationPhoneMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/phoneMethods",
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

// ListMeAuthenticationPhoneMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationPhoneMethodClient) ListMeAuthenticationPhoneMethodsComplete(ctx context.Context) (ListMeAuthenticationPhoneMethodsCompleteResult, error) {
	return c.ListMeAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx, models.PhoneAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationPhoneMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationPhoneMethodClient) ListMeAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.PhoneAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationPhoneMethodsCompleteResult, err error) {
	items := make([]models.PhoneAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationPhoneMethods(ctx)
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

	result = ListMeAuthenticationPhoneMethodsCompleteResult{
		Items: items,
	}
	return
}
