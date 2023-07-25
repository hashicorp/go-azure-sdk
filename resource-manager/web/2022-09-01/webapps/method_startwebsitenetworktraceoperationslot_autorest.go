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

type StartWebSiteNetworkTraceOperationSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type StartWebSiteNetworkTraceOperationSlotOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartWebSiteNetworkTraceOperationSlotOperationOptions() StartWebSiteNetworkTraceOperationSlotOperationOptions {
	return StartWebSiteNetworkTraceOperationSlotOperationOptions{}
}

func (o StartWebSiteNetworkTraceOperationSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartWebSiteNetworkTraceOperationSlotOperationOptions) toQueryString() map[string]interface{} {
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

// StartWebSiteNetworkTraceOperationSlot ...
func (c WebAppsClient) StartWebSiteNetworkTraceOperationSlot(ctx context.Context, id SlotId, options StartWebSiteNetworkTraceOperationSlotOperationOptions) (result StartWebSiteNetworkTraceOperationSlotOperationResponse, err error) {
	req, err := c.preparerForStartWebSiteNetworkTraceOperationSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceOperationSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartWebSiteNetworkTraceOperationSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceOperationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartWebSiteNetworkTraceOperationSlotThenPoll performs StartWebSiteNetworkTraceOperationSlot then polls until it's completed
func (c WebAppsClient) StartWebSiteNetworkTraceOperationSlotThenPoll(ctx context.Context, id SlotId, options StartWebSiteNetworkTraceOperationSlotOperationOptions) error {
	result, err := c.StartWebSiteNetworkTraceOperationSlot(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing StartWebSiteNetworkTraceOperationSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartWebSiteNetworkTraceOperationSlot: %+v", err)
	}

	return nil
}

// preparerForStartWebSiteNetworkTraceOperationSlot prepares the StartWebSiteNetworkTraceOperationSlot request.
func (c WebAppsClient) preparerForStartWebSiteNetworkTraceOperationSlot(ctx context.Context, id SlotId, options StartWebSiteNetworkTraceOperationSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/networkTrace/startOperation", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStartWebSiteNetworkTraceOperationSlot sends the StartWebSiteNetworkTraceOperationSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForStartWebSiteNetworkTraceOperationSlot(ctx context.Context, req *http.Request) (future StartWebSiteNetworkTraceOperationSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
