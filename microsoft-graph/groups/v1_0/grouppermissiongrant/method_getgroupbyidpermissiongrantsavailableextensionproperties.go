package grouppermissiongrant

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

type GetGroupByIdPermissionGrantsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionProperty
}

type GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteResult struct {
	Items []models.ExtensionProperty
}

// GetGroupByIdPermissionGrantsAvailableExtensionProperties ...
func (c GroupPermissionGrantClient) GetGroupByIdPermissionGrantsAvailableExtensionProperties(ctx context.Context, id GroupId) (result GetGroupByIdPermissionGrantsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissionGrants/getAvailableExtensionProperties", id.ID()),
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

// GetGroupByIdPermissionGrantsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c GroupPermissionGrantClient) GetGroupByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx context.Context, id GroupId) (GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, id, models.ExtensionPropertyOperationPredicate{})
}

// GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPermissionGrantClient) GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ExtensionPropertyOperationPredicate) (result GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]models.ExtensionProperty, 0)

	resp, err := c.GetGroupByIdPermissionGrantsAvailableExtensionProperties(ctx, id)
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

	result = GetGroupByIdPermissionGrantsAvailableExtensionPropertiesCompleteResult{
		Items: items,
	}
	return
}
