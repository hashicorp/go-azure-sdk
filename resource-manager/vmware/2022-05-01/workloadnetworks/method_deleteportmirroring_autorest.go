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

type DeletePortMirroringOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeletePortMirroring ...
func (c WorkloadNetworksClient) DeletePortMirroring(ctx context.Context, id PortMirroringProfileId) (result DeletePortMirroringOperationResponse, err error) {
	req, err := c.preparerForDeletePortMirroring(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeletePortMirroring", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeletePortMirroring(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeletePortMirroring", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeletePortMirroringThenPoll performs DeletePortMirroring then polls until it's completed
func (c WorkloadNetworksClient) DeletePortMirroringThenPoll(ctx context.Context, id PortMirroringProfileId) error {
	result, err := c.DeletePortMirroring(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeletePortMirroring: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeletePortMirroring: %+v", err)
	}

	return nil
}

// preparerForDeletePortMirroring prepares the DeletePortMirroring request.
func (c WorkloadNetworksClient) preparerForDeletePortMirroring(ctx context.Context, id PortMirroringProfileId) (*http.Request, error) {
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

// senderForDeletePortMirroring sends the DeletePortMirroring request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeletePortMirroring(ctx context.Context, req *http.Request) (future DeletePortMirroringOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
