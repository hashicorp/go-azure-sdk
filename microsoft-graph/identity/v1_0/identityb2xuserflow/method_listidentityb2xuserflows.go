package identityb2xuserflow

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

type ListIdentityB2xUserFlowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.B2xIdentityUserFlowCollectionResponse
}

type ListIdentityB2xUserFlowsCompleteResult struct {
	Items []models.B2xIdentityUserFlowCollectionResponse
}

// ListIdentityB2xUserFlows ...
func (c IdentityB2xUserFlowClient) ListIdentityB2xUserFlows(ctx context.Context) (result ListIdentityB2xUserFlowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/b2xUserFlows",
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
		Values *[]models.B2xIdentityUserFlowCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowsComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowClient) ListIdentityB2xUserFlowsComplete(ctx context.Context) (ListIdentityB2xUserFlowsCompleteResult, error) {
	return c.ListIdentityB2xUserFlowsCompleteMatchingPredicate(ctx, models.B2xIdentityUserFlowCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowClient) ListIdentityB2xUserFlowsCompleteMatchingPredicate(ctx context.Context, predicate models.B2xIdentityUserFlowCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowsCompleteResult, err error) {
	items := make([]models.B2xIdentityUserFlowCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlows(ctx)
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

	result = ListIdentityB2xUserFlowsCompleteResult{
		Items: items,
	}
	return
}
