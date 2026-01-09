package schemaextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSchemaExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSchemaExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSchemaExtensionOperationOptions() UpdateSchemaExtensionOperationOptions {
	return UpdateSchemaExtensionOperationOptions{}
}

func (o UpdateSchemaExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSchemaExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSchemaExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSchemaExtension - Update schemaExtension. Update properties in the definition of the specified schemaExtension.
// Additive updates to the extension can only be made when the extension is in the InDevelopment or Available status.
// This means custom properties or target resource types cannot be removed from the definition, but new custom
// properties can be added and the description of the extension changed. The update applies to all the resources that
// are included in the targetTypes property of the extension. These resources are among the supporting resource types.
// For delegated flows, the signed-in user can update a schema extension as long as the owner property of the extension
// is set to the appId of an application the signed-in user owns. That application can be the one that initially created
// the extension, or some other application owned by the signed-in user. This criteria for the owner property allows a
// signed-in user to make updates through other applications they don't own, such as Microsoft Graph Explorer. When
// using Graph Explorer to update a schemaExtension resource, include the owner property in the PATCH request body.
func (c SchemaExtensionClient) UpdateSchemaExtension(ctx context.Context, id beta.SchemaExtensionId, input beta.SchemaExtension, options UpdateSchemaExtensionOperationOptions) (result UpdateSchemaExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
