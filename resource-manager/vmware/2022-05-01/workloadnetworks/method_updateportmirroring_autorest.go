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

type UpdatePortMirroringOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdatePortMirroring ...
func (c WorkloadNetworksClient) UpdatePortMirroring(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) (result UpdatePortMirroringOperationResponse, err error) {
	req, err := c.preparerForUpdatePortMirroring(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdatePortMirroring", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdatePortMirroring(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdatePortMirroring", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdatePortMirroringThenPoll performs UpdatePortMirroring then polls until it's completed
func (c WorkloadNetworksClient) UpdatePortMirroringThenPoll(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) error {
	result, err := c.UpdatePortMirroring(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdatePortMirroring: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdatePortMirroring: %+v", err)
	}

	return nil
}

// preparerForUpdatePortMirroring prepares the UpdatePortMirroring request.
func (c WorkloadNetworksClient) preparerForUpdatePortMirroring(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) (*http.Request, error) {
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

// senderForUpdatePortMirroring sends the UpdatePortMirroring request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdatePortMirroring(ctx context.Context, req *http.Request) (future UpdatePortMirroringOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
