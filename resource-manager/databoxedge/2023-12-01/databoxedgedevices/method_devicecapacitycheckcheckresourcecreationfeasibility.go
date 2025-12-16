package databoxedgedevices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityCheckCheckResourceCreationFeasibilityOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions struct {
	CapacityName *string
}

func DefaultDeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions() DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions {
	return DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions{}
}

func (o DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CapacityName != nil {
		out.Append("capacityName", fmt.Sprintf("%v", *o.CapacityName))
	}
	return &out
}

// DeviceCapacityCheckCheckResourceCreationFeasibility ...
func (c DataBoxEdgeDevicesClient) DeviceCapacityCheckCheckResourceCreationFeasibility(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions) (result DeviceCapacityCheckCheckResourceCreationFeasibilityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceCapacityCheck", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DeviceCapacityCheckCheckResourceCreationFeasibilityThenPoll performs DeviceCapacityCheckCheckResourceCreationFeasibility then polls until it's completed
func (c DataBoxEdgeDevicesClient) DeviceCapacityCheckCheckResourceCreationFeasibilityThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options DeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions) error {
	result, err := c.DeviceCapacityCheckCheckResourceCreationFeasibility(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing DeviceCapacityCheckCheckResourceCreationFeasibility: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeviceCapacityCheckCheckResourceCreationFeasibility: %+v", err)
	}

	return nil
}
