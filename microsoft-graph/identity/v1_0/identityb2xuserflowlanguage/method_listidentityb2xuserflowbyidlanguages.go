package identityb2xuserflowlanguage

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

type ListIdentityB2xUserFlowByIdLanguagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserFlowLanguageConfigurationCollectionResponse
}

type ListIdentityB2xUserFlowByIdLanguagesCompleteResult struct {
	Items []models.UserFlowLanguageConfigurationCollectionResponse
}

// ListIdentityB2xUserFlowByIdLanguages ...
func (c IdentityB2xUserFlowLanguageClient) ListIdentityB2xUserFlowByIdLanguages(ctx context.Context, id IdentityB2xUserFlowId) (result ListIdentityB2xUserFlowByIdLanguagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/languages", id.ID()),
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
		Values *[]models.UserFlowLanguageConfigurationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityB2xUserFlowByIdLanguagesComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowLanguageClient) ListIdentityB2xUserFlowByIdLanguagesComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListIdentityB2xUserFlowByIdLanguagesCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdLanguagesCompleteMatchingPredicate(ctx, id, models.UserFlowLanguageConfigurationCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdLanguagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowLanguageClient) ListIdentityB2xUserFlowByIdLanguagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate models.UserFlowLanguageConfigurationCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdLanguagesCompleteResult, err error) {
	items := make([]models.UserFlowLanguageConfigurationCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdLanguages(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdLanguagesCompleteResult{
		Items: items,
	}
	return
}
