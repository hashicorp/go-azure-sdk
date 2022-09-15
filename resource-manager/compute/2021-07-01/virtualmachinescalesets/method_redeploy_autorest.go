package virtualmachinescalesets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedeployOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Redeploy ...
func (c VirtualMachineScaleSetsClient) Redeploy(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceIDs) (result RedeployOperationResponse, err error) {
	req, err := c.preparerForRedeploy(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "Redeploy", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRedeploy(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "Redeploy", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RedeployThenPoll performs Redeploy then polls until it's completed
func (c VirtualMachineScaleSetsClient) RedeployThenPoll(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceIDs) error {
	result, err := c.Redeploy(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Redeploy: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Redeploy: %+v", err)
	}

	return nil
}

// preparerForRedeploy prepares the Redeploy request.
func (c VirtualMachineScaleSetsClient) preparerForRedeploy(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceIDs) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/redeploy", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRedeploy sends the Redeploy request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetsClient) senderForRedeploy(ctx context.Context, req *http.Request) (future RedeployOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
