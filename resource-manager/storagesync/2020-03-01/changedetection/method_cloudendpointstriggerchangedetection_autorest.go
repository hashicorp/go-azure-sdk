package changedetection

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

type CloudEndpointsTriggerChangeDetectionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsTriggerChangeDetection ...
func (c ChangeDetectionClient) CloudEndpointsTriggerChangeDetection(ctx context.Context, id CloudEndpointId, input TriggerChangeDetectionParameters) (result CloudEndpointsTriggerChangeDetectionOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsTriggerChangeDetection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "changedetection.ChangeDetectionClient", "CloudEndpointsTriggerChangeDetection", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsTriggerChangeDetection(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "changedetection.ChangeDetectionClient", "CloudEndpointsTriggerChangeDetection", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsTriggerChangeDetectionThenPoll performs CloudEndpointsTriggerChangeDetection then polls until it's completed
func (c ChangeDetectionClient) CloudEndpointsTriggerChangeDetectionThenPoll(ctx context.Context, id CloudEndpointId, input TriggerChangeDetectionParameters) error {
	result, err := c.CloudEndpointsTriggerChangeDetection(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsTriggerChangeDetection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsTriggerChangeDetection: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsTriggerChangeDetection prepares the CloudEndpointsTriggerChangeDetection request.
func (c ChangeDetectionClient) preparerForCloudEndpointsTriggerChangeDetection(ctx context.Context, id CloudEndpointId, input TriggerChangeDetectionParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/triggerChangeDetection", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCloudEndpointsTriggerChangeDetection sends the CloudEndpointsTriggerChangeDetection request. The method will close the
// http.Response Body if it receives an error.
func (c ChangeDetectionClient) senderForCloudEndpointsTriggerChangeDetection(ctx context.Context, req *http.Request) (future CloudEndpointsTriggerChangeDetectionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
