package configurationsetting

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateConfigurationSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceManagementConfigurationSettingDefinition
}

type CreateConfigurationSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateConfigurationSettingOperationOptions() CreateConfigurationSettingOperationOptions {
	return CreateConfigurationSettingOperationOptions{}
}

func (o CreateConfigurationSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConfigurationSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConfigurationSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConfigurationSetting - Create new navigation property to configurationSettings for deviceManagement
func (c ConfigurationSettingClient) CreateConfigurationSetting(ctx context.Context, input beta.DeviceManagementConfigurationSettingDefinition, options CreateConfigurationSettingOperationOptions) (result CreateConfigurationSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/configurationSettings",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDeviceManagementConfigurationSettingDefinitionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
