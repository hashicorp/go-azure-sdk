package identityconditionalaccestemplate

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

type ListIdentityConditionalAccesTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConditionalAccessTemplateCollectionResponse
}

type ListIdentityConditionalAccesTemplatesCompleteResult struct {
	Items []models.ConditionalAccessTemplateCollectionResponse
}

// ListIdentityConditionalAccesTemplates ...
func (c IdentityConditionalAccesTemplateClient) ListIdentityConditionalAccesTemplates(ctx context.Context) (result ListIdentityConditionalAccesTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/conditionalAccess/templates",
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
		Values *[]models.ConditionalAccessTemplateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityConditionalAccesTemplatesComplete retrieves all the results into a single object
func (c IdentityConditionalAccesTemplateClient) ListIdentityConditionalAccesTemplatesComplete(ctx context.Context) (ListIdentityConditionalAccesTemplatesCompleteResult, error) {
	return c.ListIdentityConditionalAccesTemplatesCompleteMatchingPredicate(ctx, models.ConditionalAccessTemplateCollectionResponseOperationPredicate{})
}

// ListIdentityConditionalAccesTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityConditionalAccesTemplateClient) ListIdentityConditionalAccesTemplatesCompleteMatchingPredicate(ctx context.Context, predicate models.ConditionalAccessTemplateCollectionResponseOperationPredicate) (result ListIdentityConditionalAccesTemplatesCompleteResult, err error) {
	items := make([]models.ConditionalAccessTemplateCollectionResponse, 0)

	resp, err := c.ListIdentityConditionalAccesTemplates(ctx)
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

	result = ListIdentityConditionalAccesTemplatesCompleteResult{
		Items: items,
	}
	return
}
