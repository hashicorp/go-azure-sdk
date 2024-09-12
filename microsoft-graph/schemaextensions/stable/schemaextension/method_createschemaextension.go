package schemaextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSchemaExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SchemaExtension
}

// CreateSchemaExtension - Create schemaExtension. Create a new schemaExtension definition and its associated schema
// extension property to extend a supporting resource type. Schema extensions let you add strongly-typed custom data to
// a resource. The app that creates a schema extension is the owner app. Depending on the state of the extension, the
// owner app, and only the owner app, may update or delete the extension. See examples of how to define a schema
// extension that describes a training course, use the schema extension definition to create a new group with training
// course data, and add training course data to an existing group.
func (c SchemaExtensionClient) CreateSchemaExtension(ctx context.Context, input stable.SchemaExtension) (result CreateSchemaExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/schemaExtensions",
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

	var model stable.SchemaExtension
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
