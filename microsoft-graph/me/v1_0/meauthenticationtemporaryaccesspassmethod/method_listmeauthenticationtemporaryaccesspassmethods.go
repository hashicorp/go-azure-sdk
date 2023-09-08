package meauthenticationtemporaryaccesspassmethod

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

type ListMeAuthenticationTemporaryAccessPassMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TemporaryAccessPassAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationTemporaryAccessPassMethodsCompleteResult struct {
	Items []models.TemporaryAccessPassAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationTemporaryAccessPassMethods ...
func (c MeAuthenticationTemporaryAccessPassMethodClient) ListMeAuthenticationTemporaryAccessPassMethods(ctx context.Context) (result ListMeAuthenticationTemporaryAccessPassMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/temporaryAccessPassMethods",
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
		Values *[]models.TemporaryAccessPassAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeAuthenticationTemporaryAccessPassMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationTemporaryAccessPassMethodClient) ListMeAuthenticationTemporaryAccessPassMethodsComplete(ctx context.Context) (ListMeAuthenticationTemporaryAccessPassMethodsCompleteResult, error) {
	return c.ListMeAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx, models.TemporaryAccessPassAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationTemporaryAccessPassMethodClient) ListMeAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.TemporaryAccessPassAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationTemporaryAccessPassMethodsCompleteResult, err error) {
	items := make([]models.TemporaryAccessPassAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationTemporaryAccessPassMethods(ctx)
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

	result = ListMeAuthenticationTemporaryAccessPassMethodsCompleteResult{
		Items: items,
	}
	return
}
