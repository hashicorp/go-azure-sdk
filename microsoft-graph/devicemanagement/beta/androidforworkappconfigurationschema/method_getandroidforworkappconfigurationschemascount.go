package androidforworkappconfigurationschema

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAndroidForWorkAppConfigurationSchemasCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetAndroidForWorkAppConfigurationSchemasCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetAndroidForWorkAppConfigurationSchemasCountOperationOptions() GetAndroidForWorkAppConfigurationSchemasCountOperationOptions {
	return GetAndroidForWorkAppConfigurationSchemasCountOperationOptions{}
}

func (o GetAndroidForWorkAppConfigurationSchemasCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAndroidForWorkAppConfigurationSchemasCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetAndroidForWorkAppConfigurationSchemasCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAndroidForWorkAppConfigurationSchemasCount - Get the number of the resource
func (c AndroidForWorkAppConfigurationSchemaClient) GetAndroidForWorkAppConfigurationSchemasCount(ctx context.Context, options GetAndroidForWorkAppConfigurationSchemasCountOperationOptions) (result GetAndroidForWorkAppConfigurationSchemasCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/androidForWorkAppConfigurationSchemas/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
