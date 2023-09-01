package networkclouds

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

type BareMetalMachinesPowerOffOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// BareMetalMachinesPowerOff ...
func (c NetworkcloudsClient) BareMetalMachinesPowerOff(ctx context.Context, id BareMetalMachineId, input BareMetalMachinePowerOffParameters) (result BareMetalMachinesPowerOffOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/powerOff", id.ID()),
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

// BareMetalMachinesPowerOffThenPoll performs BareMetalMachinesPowerOff then polls until it's completed
func (c NetworkcloudsClient) BareMetalMachinesPowerOffThenPoll(ctx context.Context, id BareMetalMachineId, input BareMetalMachinePowerOffParameters) error {
	result, err := c.BareMetalMachinesPowerOff(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BareMetalMachinesPowerOff: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BareMetalMachinesPowerOff: %+v", err)
	}

	return nil
}
