package deviceconfigurationsallmanageddevicecertificatestate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceConfigurationsAllManagedDeviceCertificateStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ManagedAllDeviceCertificateState
}

// UpdateDeviceConfigurationsAllManagedDeviceCertificateState ...
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) UpdateDeviceConfigurationsAllManagedDeviceCertificateState(ctx context.Context, id DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId, input beta.ManagedAllDeviceCertificateState) (result UpdateDeviceConfigurationsAllManagedDeviceCertificateStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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

	var model beta.ManagedAllDeviceCertificateState
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
