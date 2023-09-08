package directoryobject

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

type GetDirectoryObjectsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionProperty
}

type GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult struct {
	Items []models.ExtensionProperty
}

// GetDirectoryObjectsAvailableExtensionProperties ...
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionProperties(ctx context.Context) (result GetDirectoryObjectsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directoryObjects/getAvailableExtensionProperties",
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
		Values *[]models.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDirectoryObjectsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionPropertiesComplete(ctx context.Context) (GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, models.ExtensionPropertyOperationPredicate{})
}

// GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, predicate models.ExtensionPropertyOperationPredicate) (result GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]models.ExtensionProperty, 0)

	resp, err := c.GetDirectoryObjectsAvailableExtensionProperties(ctx)
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

	result = GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult{
		Items: items,
	}
	return
}
