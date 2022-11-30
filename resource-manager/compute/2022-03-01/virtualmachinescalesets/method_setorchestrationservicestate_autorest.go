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

type SetOrchestrationServiceStateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SetOrchestrationServiceState ...
func (c VirtualMachineScaleSetsClient) SetOrchestrationServiceState(ctx context.Context, id VirtualMachineScaleSetId, input OrchestrationServiceStateInput) (result SetOrchestrationServiceStateOperationResponse, err error) {
	req, err := c.preparerForSetOrchestrationServiceState(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "SetOrchestrationServiceState", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSetOrchestrationServiceState(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "SetOrchestrationServiceState", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SetOrchestrationServiceStateThenPoll performs SetOrchestrationServiceState then polls until it's completed
func (c VirtualMachineScaleSetsClient) SetOrchestrationServiceStateThenPoll(ctx context.Context, id VirtualMachineScaleSetId, input OrchestrationServiceStateInput) error {
	result, err := c.SetOrchestrationServiceState(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SetOrchestrationServiceState: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SetOrchestrationServiceState: %+v", err)
	}

	return nil
}

// preparerForSetOrchestrationServiceState prepares the SetOrchestrationServiceState request.
func (c VirtualMachineScaleSetsClient) preparerForSetOrchestrationServiceState(ctx context.Context, id VirtualMachineScaleSetId, input OrchestrationServiceStateInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/setOrchestrationServiceState", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSetOrchestrationServiceState sends the SetOrchestrationServiceState request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetsClient) senderForSetOrchestrationServiceState(ctx context.Context, req *http.Request) (future SetOrchestrationServiceStateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
