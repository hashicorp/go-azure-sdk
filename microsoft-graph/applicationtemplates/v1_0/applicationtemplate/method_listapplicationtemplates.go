package applicationtemplate

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

type ListApplicationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApplicationTemplateCollectionResponse
}

type ListApplicationTemplatesCompleteResult struct {
	Items []models.ApplicationTemplateCollectionResponse
}

// ListApplicationTemplates ...
func (c ApplicationTemplateClient) ListApplicationTemplates(ctx context.Context) (result ListApplicationTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/applicationTemplates",
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
		Values *[]models.ApplicationTemplateCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationTemplatesComplete retrieves all the results into a single object
func (c ApplicationTemplateClient) ListApplicationTemplatesComplete(ctx context.Context) (ListApplicationTemplatesCompleteResult, error) {
	return c.ListApplicationTemplatesCompleteMatchingPredicate(ctx, models.ApplicationTemplateCollectionResponseOperationPredicate{})
}

// ListApplicationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTemplateClient) ListApplicationTemplatesCompleteMatchingPredicate(ctx context.Context, predicate models.ApplicationTemplateCollectionResponseOperationPredicate) (result ListApplicationTemplatesCompleteResult, err error) {
	items := make([]models.ApplicationTemplateCollectionResponse, 0)

	resp, err := c.ListApplicationTemplates(ctx)
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

	result = ListApplicationTemplatesCompleteResult{
		Items: items,
	}
	return
}
