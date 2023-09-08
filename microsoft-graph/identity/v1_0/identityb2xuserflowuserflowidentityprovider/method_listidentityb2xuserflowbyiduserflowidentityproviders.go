package identityb2xuserflowuserflowidentityprovider

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

type ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.IdentityProviderBaseCollectionResponse
}

type ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteResult struct {
	Items []models.IdentityProviderBaseCollectionResponse
}

// ListIdentityB2xUserFlowByIdUserFlowIdentityProviders ...
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProviders(ctx context.Context, id IdentityB2xUserFlowId) (result ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userFlowIdentityProviders", id.ID()),
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

// ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteMatchingPredicate(ctx, id, models.IdentityProviderBaseCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate models.IdentityProviderBaseCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteResult, err error) {
	items := make([]models.IdentityProviderBaseCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdUserFlowIdentityProviders(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdUserFlowIdentityProvidersCompleteResult{
		Items: items,
	}
	return
}
