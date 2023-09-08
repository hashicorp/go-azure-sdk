package oauth2permissiongrant

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

type ListOauth2PermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OAuth2PermissionGrantCollectionResponse
}

type ListOauth2PermissionGrantsCompleteResult struct {
	Items []models.OAuth2PermissionGrantCollectionResponse
}

// ListOauth2PermissionGrants ...
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrants(ctx context.Context) (result ListOauth2PermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/oauth2PermissionGrants",
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

// ListOauth2PermissionGrantsComplete retrieves all the results into a single object
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrantsComplete(ctx context.Context) (ListOauth2PermissionGrantsCompleteResult, error) {
	return c.ListOauth2PermissionGrantsCompleteMatchingPredicate(ctx, models.OAuth2PermissionGrantCollectionResponseOperationPredicate{})
}

// ListOauth2PermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrantsCompleteMatchingPredicate(ctx context.Context, predicate models.OAuth2PermissionGrantCollectionResponseOperationPredicate) (result ListOauth2PermissionGrantsCompleteResult, err error) {
	items := make([]models.OAuth2PermissionGrantCollectionResponse, 0)

	resp, err := c.ListOauth2PermissionGrants(ctx)
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

	result = ListOauth2PermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
