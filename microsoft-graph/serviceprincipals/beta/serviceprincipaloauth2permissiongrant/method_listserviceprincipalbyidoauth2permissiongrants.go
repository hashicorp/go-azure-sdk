package serviceprincipaloauth2permissiongrant

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

type ListServicePrincipalByIdOauth2PermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OAuth2PermissionGrantCollectionResponse
}

type ListServicePrincipalByIdOauth2PermissionGrantsCompleteResult struct {
	Items []models.OAuth2PermissionGrantCollectionResponse
}

// ListServicePrincipalByIdOauth2PermissionGrants ...
func (c ServicePrincipalOauth2PermissionGrantClient) ListServicePrincipalByIdOauth2PermissionGrants(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdOauth2PermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/oauth2PermissionGrants", id.ID()),
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

// ListServicePrincipalByIdOauth2PermissionGrantsComplete retrieves all the results into a single object
func (c ServicePrincipalOauth2PermissionGrantClient) ListServicePrincipalByIdOauth2PermissionGrantsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdOauth2PermissionGrantsCompleteResult, error) {
	return c.ListServicePrincipalByIdOauth2PermissionGrantsCompleteMatchingPredicate(ctx, id, models.OAuth2PermissionGrantCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdOauth2PermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalOauth2PermissionGrantClient) ListServicePrincipalByIdOauth2PermissionGrantsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.OAuth2PermissionGrantCollectionResponseOperationPredicate) (result ListServicePrincipalByIdOauth2PermissionGrantsCompleteResult, err error) {
	items := make([]models.OAuth2PermissionGrantCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdOauth2PermissionGrants(ctx, id)
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

	result = ListServicePrincipalByIdOauth2PermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
