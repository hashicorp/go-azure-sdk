package identityb2xuserflowlanguagedefaultpage

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

type ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserFlowLanguagePageCollectionResponse
}

type ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteResult struct {
	Items []models.UserFlowLanguagePageCollectionResponse
}

// ListIdentityB2xUserFlowByIdLanguageByIdDefaultPages ...
func (c IdentityB2xUserFlowLanguageDefaultPageClient) ListIdentityB2xUserFlowByIdLanguageByIdDefaultPages(ctx context.Context, id IdentityB2xUserFlowLanguageId) (result ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/defaultPages", id.ID()),
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

// ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesComplete retrieves all the results into a single object
func (c IdentityB2xUserFlowLanguageDefaultPageClient) ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesComplete(ctx context.Context, id IdentityB2xUserFlowLanguageId) (ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteResult, error) {
	return c.ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteMatchingPredicate(ctx, id, models.UserFlowLanguagePageCollectionResponseOperationPredicate{})
}

// ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityB2xUserFlowLanguageDefaultPageClient) ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowLanguageId, predicate models.UserFlowLanguagePageCollectionResponseOperationPredicate) (result ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteResult, err error) {
	items := make([]models.UserFlowLanguagePageCollectionResponse, 0)

	resp, err := c.ListIdentityB2xUserFlowByIdLanguageByIdDefaultPages(ctx, id)
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

	result = ListIdentityB2xUserFlowByIdLanguageByIdDefaultPagesCompleteResult{
		Items: items,
	}
	return
}
