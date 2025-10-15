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

type BareMetalMachinesRunDataExtractsRestrictedOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// BareMetalMachinesRunDataExtractsRestricted ...
func (c NetworkcloudsClient) BareMetalMachinesRunDataExtractsRestricted(ctx context.Context, id BareMetalMachineId, input BareMetalMachineRunDataExtractsParameters) (result BareMetalMachinesRunDataExtractsRestrictedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/runDataExtractsRestricted", id.ID()),
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

// BareMetalMachinesRunDataExtractsRestrictedThenPoll performs BareMetalMachinesRunDataExtractsRestricted then polls until it's completed
func (c NetworkcloudsClient) BareMetalMachinesRunDataExtractsRestrictedThenPoll(ctx context.Context, id BareMetalMachineId, input BareMetalMachineRunDataExtractsParameters) error {
	result, err := c.BareMetalMachinesRunDataExtractsRestricted(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BareMetalMachinesRunDataExtractsRestricted: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BareMetalMachinesRunDataExtractsRestricted: %+v", err)
	}

	return nil
}
