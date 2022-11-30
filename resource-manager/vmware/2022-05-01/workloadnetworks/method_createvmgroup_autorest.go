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

type CreateVMGroupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateVMGroup ...
func (c WorkloadNetworksClient) CreateVMGroup(ctx context.Context, id VMGroupId, input WorkloadNetworkVMGroup) (result CreateVMGroupOperationResponse, err error) {
	req, err := c.preparerForCreateVMGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateVMGroup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateVMGroup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateVMGroup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateVMGroupThenPoll performs CreateVMGroup then polls until it's completed
func (c WorkloadNetworksClient) CreateVMGroupThenPoll(ctx context.Context, id VMGroupId, input WorkloadNetworkVMGroup) error {
	result, err := c.CreateVMGroup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateVMGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateVMGroup: %+v", err)
	}

	return nil
}

// preparerForCreateVMGroup prepares the CreateVMGroup request.
func (c WorkloadNetworksClient) preparerForCreateVMGroup(ctx context.Context, id VMGroupId, input WorkloadNetworkVMGroup) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateVMGroup sends the CreateVMGroup request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreateVMGroup(ctx context.Context, req *http.Request) (future CreateVMGroupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
