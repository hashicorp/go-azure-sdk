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

type ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefs ...
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefs(ctx context.Context, id IdentityB2xUserFlowId) (result ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userFlowIdentityProviders/$ref", id.ID()),
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
		Values *[]models.StringCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowUserFlowIdentityProviderClient) ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate models.StringCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefs(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdUserFlowIdentityProviderRefsCompleteResult{
		Items: items,
	}
	return
}
