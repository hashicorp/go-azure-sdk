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

type CreatePortMirroringOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreatePortMirroring ...
func (c WorkloadNetworksClient) CreatePortMirroring(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) (result CreatePortMirroringOperationResponse, err error) {
	req, err := c.preparerForCreatePortMirroring(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreatePortMirroring", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreatePortMirroring(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreatePortMirroring", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreatePortMirroringThenPoll performs CreatePortMirroring then polls until it's completed
func (c WorkloadNetworksClient) CreatePortMirroringThenPoll(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) error {
	result, err := c.CreatePortMirroring(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreatePortMirroring: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreatePortMirroring: %+v", err)
	}

	return nil
}

// preparerForCreatePortMirroring prepares the CreatePortMirroring request.
func (c WorkloadNetworksClient) preparerForCreatePortMirroring(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) (*http.Request, error) {
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

// senderForCreatePortMirroring sends the CreatePortMirroring request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreatePortMirroring(ctx context.Context, req *http.Request) (future CreatePortMirroringOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
