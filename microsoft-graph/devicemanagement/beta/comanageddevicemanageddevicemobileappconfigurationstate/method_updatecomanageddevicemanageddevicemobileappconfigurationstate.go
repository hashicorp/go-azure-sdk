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

type UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions() UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions {
	return UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions{}
}

func (o UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagedDeviceManagedDeviceMobileAppConfigurationState - Update the navigation property
// managedDeviceMobileAppConfigurationStates in deviceManagement
func (c ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient) UpdateComanagedDeviceManagedDeviceMobileAppConfigurationState(ctx context.Context, id beta.DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId, input beta.ManagedDeviceMobileAppConfigurationState, options UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationOptions) (result UpdateComanagedDeviceManagedDeviceMobileAppConfigurationStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
