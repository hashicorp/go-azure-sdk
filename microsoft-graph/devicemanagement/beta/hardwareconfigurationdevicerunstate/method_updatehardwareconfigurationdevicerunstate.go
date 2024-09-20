package hardwareconfigurationdevicerunstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateHardwareConfigurationDeviceRunStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateHardwareConfigurationDeviceRunStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateHardwareConfigurationDeviceRunStateOperationOptions() UpdateHardwareConfigurationDeviceRunStateOperationOptions {
	return UpdateHardwareConfigurationDeviceRunStateOperationOptions{}
}

func (o UpdateHardwareConfigurationDeviceRunStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateHardwareConfigurationDeviceRunStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateHardwareConfigurationDeviceRunStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateHardwareConfigurationDeviceRunState - Update the navigation property deviceRunStates in deviceManagement
func (c HardwareConfigurationDeviceRunStateClient) UpdateHardwareConfigurationDeviceRunState(ctx context.Context, id beta.DeviceManagementHardwareConfigurationIdDeviceRunStateId, input beta.HardwareConfigurationDeviceState, options UpdateHardwareConfigurationDeviceRunStateOperationOptions) (result UpdateHardwareConfigurationDeviceRunStateOperationResponse, err error) {
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
