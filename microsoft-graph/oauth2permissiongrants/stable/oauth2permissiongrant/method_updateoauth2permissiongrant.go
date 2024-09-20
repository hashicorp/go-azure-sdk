package oauth2permissiongrant

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOAuth2PermissionGrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOAuth2PermissionGrantOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateOAuth2PermissionGrantOperationOptions() UpdateOAuth2PermissionGrantOperationOptions {
	return UpdateOAuth2PermissionGrantOperationOptions{}
}

func (o UpdateOAuth2PermissionGrantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOAuth2PermissionGrantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOAuth2PermissionGrantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOAuth2PermissionGrant - Update a delegated permission grant (oAuth2PermissionGrant). Update the properties of
// oAuth2PermissionGrant object, representing a delegated permission grant. An oAuth2PermissionGrant can be updated to
// change which delegated permissions are granted, by adding or removing items from the list in scopes.
func (c OAuth2PermissionGrantClient) UpdateOAuth2PermissionGrant(ctx context.Context, id stable.OAuth2PermissionGrantId, input stable.OAuth2PermissionGrant, options UpdateOAuth2PermissionGrantOperationOptions) (result UpdateOAuth2PermissionGrantOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
