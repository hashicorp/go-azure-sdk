package devicecapacitycheck

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

type CheckResourceCreationFeasibilityOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type CheckResourceCreationFeasibilityOperationOptions struct {
	CapacityName *string
}

func DefaultCheckResourceCreationFeasibilityOperationOptions() CheckResourceCreationFeasibilityOperationOptions {
	return CheckResourceCreationFeasibilityOperationOptions{}
}

func (o CheckResourceCreationFeasibilityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckResourceCreationFeasibilityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o CheckResourceCreationFeasibilityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CapacityName != nil {
		out.Append("capacityName", fmt.Sprintf("%v", *o.CapacityName))
	}
	return &out
}

// CheckResourceCreationFeasibility ...
func (c DeviceCapacityCheckClient) CheckResourceCreationFeasibility(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options CheckResourceCreationFeasibilityOperationOptions) (result CheckResourceCreationFeasibilityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/deviceCapacityCheck", id.ID()),
		OptionsObject: options,
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

// CheckResourceCreationFeasibilityThenPoll performs CheckResourceCreationFeasibility then polls until it's completed
func (c DeviceCapacityCheckClient) CheckResourceCreationFeasibilityThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options CheckResourceCreationFeasibilityOperationOptions) error {
	result, err := c.CheckResourceCreationFeasibility(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing CheckResourceCreationFeasibility: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CheckResourceCreationFeasibility: %+v", err)
	}

	return nil
}
