package customauthenticationextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCustomAuthenticationExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCustomAuthenticationExtensionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCustomAuthenticationExtensionOperationOptions() UpdateCustomAuthenticationExtensionOperationOptions {
	return UpdateCustomAuthenticationExtensionOperationOptions{}
}

func (o UpdateCustomAuthenticationExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCustomAuthenticationExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCustomAuthenticationExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCustomAuthenticationExtension - Update customAuthenticationExtension. Update the properties of a
// customAuthenticationExtension object.
func (c CustomAuthenticationExtensionClient) UpdateCustomAuthenticationExtension(ctx context.Context, id stable.IdentityCustomAuthenticationExtensionId, input stable.CustomAuthenticationExtension, options UpdateCustomAuthenticationExtensionOperationOptions) (result UpdateCustomAuthenticationExtensionOperationResponse, err error) {
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
