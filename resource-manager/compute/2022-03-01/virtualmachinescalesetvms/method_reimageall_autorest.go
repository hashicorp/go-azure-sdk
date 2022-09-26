package virtualmachinescalesetvms

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

type ReimageAllOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ReimageAll ...
func (c VirtualMachineScaleSetVMsClient) ReimageAll(ctx context.Context, id VirtualMachineId) (result ReimageAllOperationResponse, err error) {
	req, err := c.preparerForReimageAll(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "ReimageAll", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForReimageAll(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "ReimageAll", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ReimageAllThenPoll performs ReimageAll then polls until it's completed
func (c VirtualMachineScaleSetVMsClient) ReimageAllThenPoll(ctx context.Context, id VirtualMachineId) error {
	result, err := c.ReimageAll(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ReimageAll: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ReimageAll: %+v", err)
	}

	return nil
}

// preparerForReimageAll prepares the ReimageAll request.
func (c VirtualMachineScaleSetVMsClient) preparerForReimageAll(ctx context.Context, id VirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reimageall", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForReimageAll sends the ReimageAll request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetVMsClient) senderForReimageAll(ctx context.Context, req *http.Request) (future ReimageAllOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
