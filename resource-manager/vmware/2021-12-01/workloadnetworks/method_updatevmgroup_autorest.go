package workloadnetworks

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

type UpdateVMGroupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateVMGroup ...
func (c WorkloadNetworksClient) UpdateVMGroup(ctx context.Context, id VmGroupId, input WorkloadNetworkVMGroup) (result UpdateVMGroupOperationResponse, err error) {
	req, err := c.preparerForUpdateVMGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateVMGroup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateVMGroup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateVMGroup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateVMGroupThenPoll performs UpdateVMGroup then polls until it's completed
func (c WorkloadNetworksClient) UpdateVMGroupThenPoll(ctx context.Context, id VmGroupId, input WorkloadNetworkVMGroup) error {
	result, err := c.UpdateVMGroup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateVMGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateVMGroup: %+v", err)
	}

	return nil
}

// preparerForUpdateVMGroup prepares the UpdateVMGroup request.
func (c WorkloadNetworksClient) preparerForUpdateVMGroup(ctx context.Context, id VmGroupId, input WorkloadNetworkVMGroup) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateVMGroup sends the UpdateVMGroup request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdateVMGroup(ctx context.Context, req *http.Request) (future UpdateVMGroupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
