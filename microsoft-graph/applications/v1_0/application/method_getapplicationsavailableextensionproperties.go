package application

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

type GetApplicationsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionProperty
}

type GetApplicationsAvailableExtensionPropertiesCompleteResult struct {
	Items []models.ExtensionProperty
}

// GetApplicationsAvailableExtensionProperties ...
func (c ApplicationClient) GetApplicationsAvailableExtensionProperties(ctx context.Context) (result GetApplicationsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/applications/getAvailableExtensionProperties",
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

// GetApplicationsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c ApplicationClient) GetApplicationsAvailableExtensionPropertiesComplete(ctx context.Context) (GetApplicationsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, models.ExtensionPropertyOperationPredicate{})
}

// GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationClient) GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, predicate models.ExtensionPropertyOperationPredicate) (result GetApplicationsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]models.ExtensionProperty, 0)

	resp, err := c.GetApplicationsAvailableExtensionProperties(ctx)
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

	result = GetApplicationsAvailableExtensionPropertiesCompleteResult{
		Items: items,
	}
	return
}
