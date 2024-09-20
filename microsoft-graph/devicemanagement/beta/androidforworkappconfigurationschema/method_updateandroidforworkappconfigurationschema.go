package androidforworkappconfigurationschema

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAndroidForWorkAppConfigurationSchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAndroidForWorkAppConfigurationSchemaOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAndroidForWorkAppConfigurationSchemaOperationOptions() UpdateAndroidForWorkAppConfigurationSchemaOperationOptions {
	return UpdateAndroidForWorkAppConfigurationSchemaOperationOptions{}
}

func (o UpdateAndroidForWorkAppConfigurationSchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAndroidForWorkAppConfigurationSchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAndroidForWorkAppConfigurationSchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAndroidForWorkAppConfigurationSchema - Update the navigation property androidForWorkAppConfigurationSchemas in
// deviceManagement
func (c AndroidForWorkAppConfigurationSchemaClient) UpdateAndroidForWorkAppConfigurationSchema(ctx context.Context, id beta.DeviceManagementAndroidForWorkAppConfigurationSchemaId, input beta.AndroidForWorkAppConfigurationSchema, options UpdateAndroidForWorkAppConfigurationSchemaOperationOptions) (result UpdateAndroidForWorkAppConfigurationSchemaOperationResponse, err error) {
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
