package oauth2permissiongrant

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOAuth2PermissionGrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OAuth2PermissionGrant
}

// CreateOAuth2PermissionGrant - Create oAuth2PermissionGrant (a delegated permission grant). Create a delegated
// permission grant, represented by an oAuth2PermissionGrant object. A delegated permission grant authorizes a client
// service principal (representing a client application) to access a resource service principal (representing an API),
// on behalf of a signed-in user, for the level of access limited by the delegated permissions which were granted.
func (c OAuth2PermissionGrantClient) CreateOAuth2PermissionGrant(ctx context.Context, input beta.OAuth2PermissionGrant) (result CreateOAuth2PermissionGrantOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/oauth2PermissionGrants",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.OAuth2PermissionGrant
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
