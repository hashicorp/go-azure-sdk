package manageddeviceencryptionstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateManagedDeviceEncryptionStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceEncryptionStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateManagedDeviceEncryptionStateOperationOptions() UpdateManagedDeviceEncryptionStateOperationOptions {
	return UpdateManagedDeviceEncryptionStateOperationOptions{}
}

func (o UpdateManagedDeviceEncryptionStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceEncryptionStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceEncryptionStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceEncryptionState - Update the navigation property managedDeviceEncryptionStates in deviceManagement
func (c ManagedDeviceEncryptionStateClient) UpdateManagedDeviceEncryptionState(ctx context.Context, id beta.DeviceManagementManagedDeviceEncryptionStateId, input beta.ManagedDeviceEncryptionState, options UpdateManagedDeviceEncryptionStateOperationOptions) (result UpdateManagedDeviceEncryptionStateOperationResponse, err error) {
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
