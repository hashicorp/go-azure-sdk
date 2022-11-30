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

type DeleteVMGroupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteVMGroup ...
func (c WorkloadNetworksClient) DeleteVMGroup(ctx context.Context, id VMGroupId) (result DeleteVMGroupOperationResponse, err error) {
	req, err := c.preparerForDeleteVMGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteVMGroup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteVMGroup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteVMGroup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteVMGroupThenPoll performs DeleteVMGroup then polls until it's completed
func (c WorkloadNetworksClient) DeleteVMGroupThenPoll(ctx context.Context, id VMGroupId) error {
	result, err := c.DeleteVMGroup(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteVMGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteVMGroup: %+v", err)
	}

	return nil
}

// preparerForDeleteVMGroup prepares the DeleteVMGroup request.
func (c WorkloadNetworksClient) preparerForDeleteVMGroup(ctx context.Context, id VMGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteVMGroup sends the DeleteVMGroup request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeleteVMGroup(ctx context.Context, req *http.Request) (future DeleteVMGroupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
