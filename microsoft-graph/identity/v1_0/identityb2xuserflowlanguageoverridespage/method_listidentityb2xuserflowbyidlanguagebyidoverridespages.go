package identityb2xuserflowlanguageoverridespage

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

type ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserFlowLanguagePageCollectionResponse
}

type ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteResult struct {
	Items []models.UserFlowLanguagePageCollectionResponse
}

// ListIdentityB2xUserFlowByIdLanguageByIdOverridesPages ...
func (c IdentityB2xUserFlowLanguageOverridesPageClient) ListIdentityB2xUserFlowByIdLanguageByIdOverridesPages(ctx context.Context, id IdentityB2xUserFlowLanguageId) (result ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/overridesPages", id.ID()),
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
		Values *[]models.UserFlowLanguagePageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowLanguageOverridesPageClient) ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesComplete(ctx context.Context, id IdentityB2xUserFlowLanguageId) (ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteMatchingPredicate(ctx, id, models.UserFlowLanguagePageCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowLanguageOverridesPageClient) ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowLanguageId, predicate models.UserFlowLanguagePageCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteResult, err error) {
	items := make([]models.UserFlowLanguagePageCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdLanguageByIdOverridesPages(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdLanguageByIdOverridesPagesCompleteResult{
		Items: items,
	}
	return
}
