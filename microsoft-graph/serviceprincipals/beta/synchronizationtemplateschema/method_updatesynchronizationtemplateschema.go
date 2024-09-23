package synchronizationtemplateschema

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

type UpdateSynchronizationTemplateSchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSynchronizationTemplateSchemaOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSynchronizationTemplateSchemaOperationOptions() UpdateSynchronizationTemplateSchemaOperationOptions {
	return UpdateSynchronizationTemplateSchemaOperationOptions{}
}

func (o UpdateSynchronizationTemplateSchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSynchronizationTemplateSchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSynchronizationTemplateSchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSynchronizationTemplateSchema - Update the navigation property schema in servicePrincipals
func (c SynchronizationTemplateSchemaClient) UpdateSynchronizationTemplateSchema(ctx context.Context, id beta.ServicePrincipalIdSynchronizationTemplateId, input beta.SynchronizationSchema, options UpdateSynchronizationTemplateSchemaOperationOptions) (result UpdateSynchronizationTemplateSchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/schema", id.ID()),
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
