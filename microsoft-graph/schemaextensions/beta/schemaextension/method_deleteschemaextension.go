package schemaextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSchemaExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSchemaExtensionOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteSchemaExtensionOperationOptions() DeleteSchemaExtensionOperationOptions {
	return DeleteSchemaExtensionOperationOptions{}
}

func (o DeleteSchemaExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSchemaExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteSchemaExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSchemaExtension - Delete schemaExtension. Delete the definition of a schema extension. In app-only scenarios,
// only the app that created the schema extension (owner app) can delete the schema extension definition, and only when
// the extension is in the InDevelopment state. In delegated scenarios, the owner of the owner app can delete the schema
// extension definition, and only when the extension is in the InDevelopment state. Deleting a schema extension
// definition before deleting the data associated with the extension in the target resources makes the data
// inaccessible. To recover the data, you can recreate the schema extension definition with the same configuration, but
// only if you used the verified domain for the schema extension id.
func (c SchemaExtensionClient) DeleteSchemaExtension(ctx context.Context, id beta.SchemaExtensionId, options DeleteSchemaExtensionOperationOptions) (result DeleteSchemaExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
