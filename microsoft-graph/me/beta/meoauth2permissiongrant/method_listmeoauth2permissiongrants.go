package meoauth2permissiongrant

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

type ListMeOauth2PermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OAuth2PermissionGrantCollectionResponse
}

type ListMeOauth2PermissionGrantsCompleteResult struct {
	Items []models.OAuth2PermissionGrantCollectionResponse
}

// ListMeOauth2PermissionGrants ...
func (c MeOauth2PermissionGrantClient) ListMeOauth2PermissionGrants(ctx context.Context) (result ListMeOauth2PermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/oauth2PermissionGrants",
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
		Values *[]models.OAuth2PermissionGrantCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOauth2PermissionGrantsComplete retrieves all the results into a single object
func (c MeOauth2PermissionGrantClient) ListMeOauth2PermissionGrantsComplete(ctx context.Context) (ListMeOauth2PermissionGrantsCompleteResult, error) {
	return c.ListMeOauth2PermissionGrantsCompleteMatchingPredicate(ctx, models.OAuth2PermissionGrantCollectionResponseOperationPredicate{})
}

// ListMeOauth2PermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOauth2PermissionGrantClient) ListMeOauth2PermissionGrantsCompleteMatchingPredicate(ctx context.Context, predicate models.OAuth2PermissionGrantCollectionResponseOperationPredicate) (result ListMeOauth2PermissionGrantsCompleteResult, err error) {
	items := make([]models.OAuth2PermissionGrantCollectionResponse, 0)

	resp, err := c.ListMeOauth2PermissionGrants(ctx)
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

	result = ListMeOauth2PermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
