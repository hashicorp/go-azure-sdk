package comanageddevicemanageddevicemobileappconfigurationstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ManagedDeviceMobileAppConfigurationState
}

type GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions() GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions {
	return GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions{}
}

func (o GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToOData() *odata.Query {
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

func (o GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetComanagedDeviceManagedDeviceMobileAppConfigurationState - Get managedDeviceMobileAppConfigurationStates from
// deviceManagement. Managed device mobile app configuration states for this device.
func (c ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient) GetComanagedDeviceManagedDeviceMobileAppConfigurationState(ctx context.Context, id beta.DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId, options GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) (result GetComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationResponse, err error) {
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

	var model beta.ManagedDeviceMobileAppConfigurationState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
