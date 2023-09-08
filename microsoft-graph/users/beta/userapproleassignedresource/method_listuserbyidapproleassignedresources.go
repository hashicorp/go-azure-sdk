package userapproleassignedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdAppRoleAssignedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ServicePrincipalCollectionResponse
}

type ListUserByIdAppRoleAssignedResourcesCompleteResult struct {
	Items []models.ServicePrincipalCollectionResponse
}

// ListUserByIdAppRoleAssignedResources ...
func (c UserAppRoleAssignedResourceClient) ListUserByIdAppRoleAssignedResources(ctx context.Context, id UserId) (result ListUserByIdAppRoleAssignedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appRoleAssignedResources", id.ID()),
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
		Values *[]models.ServicePrincipalCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAppRoleAssignedResourcesComplete retrieves all the results into a single object
func (c UserAppRoleAssignedResourceClient) ListUserByIdAppRoleAssignedResourcesComplete(ctx context.Context, id UserId) (ListUserByIdAppRoleAssignedResourcesCompleteResult, error) {
	return c.ListUserByIdAppRoleAssignedResourcesCompleteMatchingPredicate(ctx, id, models.ServicePrincipalCollectionResponseOperationPredicate{})
}

// ListUserByIdAppRoleAssignedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAppRoleAssignedResourceClient) ListUserByIdAppRoleAssignedResourcesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ServicePrincipalCollectionResponseOperationPredicate) (result ListUserByIdAppRoleAssignedResourcesCompleteResult, err error) {
	items := make([]models.ServicePrincipalCollectionResponse, 0)

	resp, err := c.ListUserByIdAppRoleAssignedResources(ctx, id)
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

	result = ListUserByIdAppRoleAssignedResourcesCompleteResult{
		Items: items,
	}
	return
}
