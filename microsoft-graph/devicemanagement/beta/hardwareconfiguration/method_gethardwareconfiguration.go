package hardwareconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHardwareConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwareConfiguration
}

type GetHardwareConfigurationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetHardwareConfigurationOperationOptions() GetHardwareConfigurationOperationOptions {
	return GetHardwareConfigurationOperationOptions{}
}

func (o GetHardwareConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHardwareConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetHardwareConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetHardwareConfiguration - Get hardwareConfigurations from deviceManagement. BIOS configuration and other settings
// provides customers the ability to configure hardware/bios settings on the enrolled Windows 10/11 Entra ID joined
// devices by uploading a configuration file generated with their OEM tool (e.g. Dell Command tool). A BIOS
// configuration policy can be assigned to multiple devices, allowing admins to remotely control a device's hardware
// properties (e.g. enable Secure Boot) from the Intune Portal. Supported for Dell only at this time.
func (c HardwareConfigurationClient) GetHardwareConfiguration(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, options GetHardwareConfigurationOperationOptions) (result GetHardwareConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.HardwareConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
