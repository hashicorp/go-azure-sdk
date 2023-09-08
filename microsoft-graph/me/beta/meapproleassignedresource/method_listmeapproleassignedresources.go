package meapproleassignedresource

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

type ListMeAppRoleAssignedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ServicePrincipalCollectionResponse
}

type ListMeAppRoleAssignedResourcesCompleteResult struct {
	Items []models.ServicePrincipalCollectionResponse
}

// ListMeAppRoleAssignedResources ...
func (c MeAppRoleAssignedResourceClient) ListMeAppRoleAssignedResources(ctx context.Context) (result ListMeAppRoleAssignedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/appRoleAssignedResources",
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

// ListMeAppRoleAssignedResourcesComplete retrieves all the results into a single object
func (c MeAppRoleAssignedResourceClient) ListMeAppRoleAssignedResourcesComplete(ctx context.Context) (ListMeAppRoleAssignedResourcesCompleteResult, error) {
	return c.ListMeAppRoleAssignedResourcesCompleteMatchingPredicate(ctx, models.ServicePrincipalCollectionResponseOperationPredicate{})
}

// ListMeAppRoleAssignedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeAppRoleAssignedResourceClient) ListMeAppRoleAssignedResourcesCompleteMatchingPredicate(ctx context.Context, predicate models.ServicePrincipalCollectionResponseOperationPredicate) (result ListMeAppRoleAssignedResourcesCompleteResult, err error) {
	items := make([]models.ServicePrincipalCollectionResponse, 0)

	resp, err := c.ListMeAppRoleAssignedResources(ctx)
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

	result = ListMeAppRoleAssignedResourcesCompleteResult{
		Items: items,
	}
	return
}
