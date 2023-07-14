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

type ListForResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Permission
}

type ListForResourceCompleteResult struct {
	Items []Permission
}

// ListForResource ...
func (c PermissionsClient) ListForResource(ctx context.Context, id commonids.ScopeId) (result ListForResourceOperationResponse, err error) {
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

// ListForResourceComplete retrieves all the results into a single object
func (c PermissionsClient) ListForResourceComplete(ctx context.Context, id commonids.ScopeId) (ListForResourceCompleteResult, error) {
	return c.ListForResourceCompleteMatchingPredicate(ctx, id, PermissionOperationPredicate{})
}

// ListForResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PermissionsClient) ListForResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate PermissionOperationPredicate) (result ListForResourceCompleteResult, err error) {
	items := make([]Permission, 0)

	resp, err := c.ListForResource(ctx, id)
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

	result = ListForResourceCompleteResult{
		Items: items,
	}
	return
}
