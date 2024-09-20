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

type CreateAndroidForWorkAppConfigurationSchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AndroidForWorkAppConfigurationSchema
}

type CreateAndroidForWorkAppConfigurationSchemaOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAndroidForWorkAppConfigurationSchemaOperationOptions() CreateAndroidForWorkAppConfigurationSchemaOperationOptions {
	return CreateAndroidForWorkAppConfigurationSchemaOperationOptions{}
}

func (o CreateAndroidForWorkAppConfigurationSchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidForWorkAppConfigurationSchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidForWorkAppConfigurationSchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidForWorkAppConfigurationSchema - Create new navigation property to androidForWorkAppConfigurationSchemas
// for deviceManagement
func (c AndroidForWorkAppConfigurationSchemaClient) CreateAndroidForWorkAppConfigurationSchema(ctx context.Context, input beta.AndroidForWorkAppConfigurationSchema, options CreateAndroidForWorkAppConfigurationSchemaOperationOptions) (result CreateAndroidForWorkAppConfigurationSchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/androidForWorkAppConfigurationSchemas",
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

	var model beta.AndroidForWorkAppConfigurationSchema
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
