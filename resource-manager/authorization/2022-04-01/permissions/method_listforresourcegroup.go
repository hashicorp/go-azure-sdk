package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListForResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Permission
}

type ListForResourceGroupCompleteResult struct {
	Items []Permission
}

// ListForResourceGroup ...
func (c PermissionsClient) ListForResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ListForResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Authorization/permissions", id.ID()),
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
		Values *[]Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListForResourceGroupComplete retrieves all the results into a single object
func (c PermissionsClient) ListForResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ListForResourceGroupCompleteResult, error) {
	return c.ListForResourceGroupCompleteMatchingPredicate(ctx, id, PermissionOperationPredicate{})
}

// ListForResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsClient) ListForResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate PermissionOperationPredicate) (result ListForResourceGroupCompleteResult, err error) {
	items := make([]Permission, 0)

	resp, err := c.ListForResourceGroup(ctx, id)
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

	result = ListForResourceGroupCompleteResult{
		Items: items,
	}
	return
}
