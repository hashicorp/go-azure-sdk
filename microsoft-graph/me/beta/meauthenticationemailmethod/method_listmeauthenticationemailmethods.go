package meauthenticationemailmethod

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

type ListMeAuthenticationEmailMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EmailAuthenticationMethodCollectionResponse
}

type ListMeAuthenticationEmailMethodsCompleteResult struct {
	Items []models.EmailAuthenticationMethodCollectionResponse
}

// ListMeAuthenticationEmailMethods ...
func (c MeAuthenticationEmailMethodClient) ListMeAuthenticationEmailMethods(ctx context.Context) (result ListMeAuthenticationEmailMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/authentication/emailMethods",
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
		Values *[]models.EmailAuthenticationMethodCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeAuthenticationEmailMethodsComplete retrieves all the results into a single object
func (c MeAuthenticationEmailMethodClient) ListMeAuthenticationEmailMethodsComplete(ctx context.Context) (ListMeAuthenticationEmailMethodsCompleteResult, error) {
	return c.ListMeAuthenticationEmailMethodsCompleteMatchingPredicate(ctx, models.EmailAuthenticationMethodCollectionResponseOperationPredicate{})
}

// ListMeAuthenticationEmailMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAuthenticationEmailMethodClient) ListMeAuthenticationEmailMethodsCompleteMatchingPredicate(ctx context.Context, predicate models.EmailAuthenticationMethodCollectionResponseOperationPredicate) (result ListMeAuthenticationEmailMethodsCompleteResult, err error) {
	items := make([]models.EmailAuthenticationMethodCollectionResponse, 0)

	resp, err := c.ListMeAuthenticationEmailMethods(ctx)
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

	result = ListMeAuthenticationEmailMethodsCompleteResult{
		Items: items,
	}
	return
}
