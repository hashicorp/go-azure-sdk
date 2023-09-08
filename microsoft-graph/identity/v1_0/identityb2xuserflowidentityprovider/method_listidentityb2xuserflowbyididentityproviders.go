package identityb2xuserflowidentityprovider

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

type ListIdentityB2xUserFlowByIdIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityProviderCollectionResponse
}

type ListIdentityB2xUserFlowByIdIdentityProvidersCompleteResult struct {
	Items []models.IdentityProviderCollectionResponse
}

// ListIdentityB2xUserFlowByIdIdentityProviders ...
func (c IdentityB2xUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdIdentityProviders(ctx context.Context, id IdentityB2xUserFlowId) (result ListIdentityB2xUserFlowByIdIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/identityProviders", id.ID()),
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
		Values *[]models.IdentityProviderCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowByIdIdentityProvidersComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdIdentityProvidersComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListIdentityB2xUserFlowByIdIdentityProvidersCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdIdentityProvidersCompleteMatchingPredicate(ctx, id, models.IdentityProviderCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdIdentityProvidersCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate models.IdentityProviderCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdIdentityProvidersCompleteResult, err error) {
	items := make([]models.IdentityProviderCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdIdentityProviders(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdIdentityProvidersCompleteResult{
		Items: items,
	}
	return
}
