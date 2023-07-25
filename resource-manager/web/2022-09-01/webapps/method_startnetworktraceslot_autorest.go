package webapps

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

type StartNetworkTraceSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type StartNetworkTraceSlotOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartNetworkTraceSlotOperationOptions() StartNetworkTraceSlotOperationOptions {
	return StartNetworkTraceSlotOperationOptions{}
}

func (o StartNetworkTraceSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartNetworkTraceSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DurationInSeconds != nil {
		out["durationInSeconds"] = *o.DurationInSeconds
	}

	if o.MaxFrameLength != nil {
		out["maxFrameLength"] = *o.MaxFrameLength
	}

	if o.SasUrl != nil {
		out["sasUrl"] = *o.SasUrl
	}

	return out
}

// StartNetworkTraceSlot ...
func (c WebAppsClient) StartNetworkTraceSlot(ctx context.Context, id SlotId, options StartNetworkTraceSlotOperationOptions) (result StartNetworkTraceSlotOperationResponse, err error) {
	req, err := c.preparerForStartNetworkTraceSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartNetworkTraceSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartNetworkTraceSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartNetworkTraceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartNetworkTraceSlotThenPoll performs StartNetworkTraceSlot then polls until it's completed
func (c WebAppsClient) StartNetworkTraceSlotThenPoll(ctx context.Context, id SlotId, options StartNetworkTraceSlotOperationOptions) error {
	result, err := c.StartNetworkTraceSlot(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing StartNetworkTraceSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartNetworkTraceSlot: %+v", err)
	}

	return nil
}

// preparerForStartNetworkTraceSlot prepares the StartNetworkTraceSlot request.
func (c WebAppsClient) preparerForStartNetworkTraceSlot(ctx context.Context, id SlotId, options StartNetworkTraceSlotOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/startNetworkTrace", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStartNetworkTraceSlot sends the StartNetworkTraceSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForStartNetworkTraceSlot(ctx context.Context, req *http.Request) (future StartNetworkTraceSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
