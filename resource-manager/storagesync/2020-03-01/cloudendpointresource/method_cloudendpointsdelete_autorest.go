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

type CloudEndpointsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsDelete ...
func (c CloudEndpointResourceClient) CloudEndpointsDelete(ctx context.Context, id CloudEndpointId) (result CloudEndpointsDeleteOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsDeleteThenPoll performs CloudEndpointsDelete then polls until it's completed
func (c CloudEndpointResourceClient) CloudEndpointsDeleteThenPoll(ctx context.Context, id CloudEndpointId) error {
	result, err := c.CloudEndpointsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsDelete: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsDelete prepares the CloudEndpointsDelete request.
func (c CloudEndpointResourceClient) preparerForCloudEndpointsDelete(ctx context.Context, id CloudEndpointId) (*http.Request, error) {
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

// senderForCloudEndpointsDelete sends the CloudEndpointsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c CloudEndpointResourceClient) senderForCloudEndpointsDelete(ctx context.Context, req *http.Request) (future CloudEndpointsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
