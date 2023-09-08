package userauthenticationwindowshelloforbusinessmethod

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

type ListUserByIdAuthenticationWindowsHelloForBusinessMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.WindowsHelloForBusinessAuthenticationMethodCollectionResponse
}

type ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteResult struct {
	Items []models.WindowsHelloForBusinessAuthenticationMethodCollectionResponse
}

// ListUserByIdAuthenticationWindowsHelloForBusinessMethods ...
func (c UserAuthenticationWindowsHelloForBusinessMethodClient) ListUserByIdAuthenticationWindowsHelloForBusinessMethods(ctx context.Context, id UserId) (result ListUserByIdAuthenticationWindowsHelloForBusinessMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authentication/windowsHelloForBusinessMethods", id.ID()),
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
		Values *[]models.WindowsHelloForBusinessAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAuthenticationWindowsHelloForBusinessMethodsComplete retrieves all the results into a single object
func (c UserAuthenticationWindowsHelloForBusinessMethodClient) ListUserByIdAuthenticationWindowsHelloForBusinessMethodsComplete(ctx context.Context, id UserId) (ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteResult, error) {
	return c.ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx, id, models.WindowsHelloForBusinessAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAuthenticationWindowsHelloForBusinessMethodClient) ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.WindowsHelloForBusinessAuthenticationMethodCollectionResponseOperationPredicate) (result ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteResult, err error) {
	items := make([]models.WindowsHelloForBusinessAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListUserByIdAuthenticationWindowsHelloForBusinessMethods(ctx, id)
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

	result = ListUserByIdAuthenticationWindowsHelloForBusinessMethodsCompleteResult{
		Items: items,
	}
	return
}
