package identityidentityprovider

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

type ListIdentityIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityProviderBaseCollectionResponse
}

type ListIdentityIdentityProvidersCompleteResult struct {
	Items []models.IdentityProviderBaseCollectionResponse
}

// ListIdentityIdentityProviders ...
func (c IdentityIdentityProviderClient) ListIdentityIdentityProviders(ctx context.Context) (result ListIdentityIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/identityProviders",
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
		Values *[]models.IdentityProviderBaseCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityIdentityProvidersComplete retrieves all the results into a single object
func (c IdentityIdentityProviderClient) ListIdentityIdentityProvidersComplete(ctx context.Context) (ListIdentityIdentityProvidersCompleteResult, error) {
	return c.ListIdentityIdentityProvidersCompleteMatchingPredicate(ctx, models.IdentityProviderBaseCollectionResponseOperationPredicate{})
}

// ListIdentityIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityIdentityProviderClient) ListIdentityIdentityProvidersCompleteMatchingPredicate(ctx context.Context, predicate models.IdentityProviderBaseCollectionResponseOperationPredicate) (result ListIdentityIdentityProvidersCompleteResult, err error) {
	items := make([]models.IdentityProviderBaseCollectionResponse, 0)

	resp, err := c.ListIdentityIdentityProviders(ctx)
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

	result = ListIdentityIdentityProvidersCompleteResult{
		Items: items,
	}
	return
}
