package directorycustomsecurityattributedefinitionallowedvalue

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

type ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AllowedValueCollectionResponse
}

type ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteResult struct {
	Items []models.AllowedValueCollectionResponse
}

// ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValues ...
func (c DirectoryCustomSecurityAttributeDefinitionAllowedValueClient) ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValues(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId) (result ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/allowedValues", id.ID()),
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
		Values *[]models.AllowedValueCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesComplete retrieves all the results into a single object
func (c DirectoryCustomSecurityAttributeDefinitionAllowedValueClient) ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesComplete(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId) (ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteResult, error) {
	return c.ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteMatchingPredicate(ctx, id, models.AllowedValueCollectionResponseOperationPredicate{})
}

// ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryCustomSecurityAttributeDefinitionAllowedValueClient) ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteMatchingPredicate(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId, predicate models.AllowedValueCollectionResponseOperationPredicate) (result ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteResult, err error) {
	items := make([]models.AllowedValueCollectionResponse, 0)

	resp, err := c.ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValues(ctx, id)
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

	result = ListDirectoryCustomSecurityAttributeDefinitionByIdAllowedValuesCompleteResult{
		Items: items,
	}
	return
}
