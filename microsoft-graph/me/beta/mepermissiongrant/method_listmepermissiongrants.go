package mepermissiongrant

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

type ListMePermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ResourceSpecificPermissionGrantCollectionResponse
}

type ListMePermissionGrantsCompleteResult struct {
	Items []models.ResourceSpecificPermissionGrantCollectionResponse
}

// ListMePermissionGrants ...
func (c MePermissionGrantClient) ListMePermissionGrants(ctx context.Context) (result ListMePermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/permissionGrants",
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
		Values *[]models.ResourceSpecificPermissionGrantCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePermissionGrantsComplete retrieves all the results into a single object
func (c MePermissionGrantClient) ListMePermissionGrantsComplete(ctx context.Context) (ListMePermissionGrantsCompleteResult, error) {
	return c.ListMePermissionGrantsCompleteMatchingPredicate(ctx, models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate{})
}

// ListMePermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePermissionGrantClient) ListMePermissionGrantsCompleteMatchingPredicate(ctx context.Context, predicate models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate) (result ListMePermissionGrantsCompleteResult, err error) {
	items := make([]models.ResourceSpecificPermissionGrantCollectionResponse, 0)

	resp, err := c.ListMePermissionGrants(ctx)
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

	result = ListMePermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
