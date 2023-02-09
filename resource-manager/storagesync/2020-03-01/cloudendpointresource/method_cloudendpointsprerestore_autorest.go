package cloudendpointresource

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

type CloudEndpointsPreRestoreOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsPreRestore ...
func (c CloudEndpointResourceClient) CloudEndpointsPreRestore(ctx context.Context, id CloudEndpointId, input PreRestoreRequest) (result CloudEndpointsPreRestoreOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsPreRestore(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsPreRestore", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsPreRestore(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsPreRestore", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsPreRestoreThenPoll performs CloudEndpointsPreRestore then polls until it's completed
func (c CloudEndpointResourceClient) CloudEndpointsPreRestoreThenPoll(ctx context.Context, id CloudEndpointId, input PreRestoreRequest) error {
	result, err := c.CloudEndpointsPreRestore(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsPreRestore: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsPreRestore: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsPreRestore prepares the CloudEndpointsPreRestore request.
func (c CloudEndpointResourceClient) preparerForCloudEndpointsPreRestore(ctx context.Context, id CloudEndpointId, input PreRestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/prerestore", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCloudEndpointsPreRestore sends the CloudEndpointsPreRestore request. The method will close the
// http.Response Body if it receives an error.
func (c CloudEndpointResourceClient) senderForCloudEndpointsPreRestore(ctx context.Context, req *http.Request) (future CloudEndpointsPreRestoreOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
