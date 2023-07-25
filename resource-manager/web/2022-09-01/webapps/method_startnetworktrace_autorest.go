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

type StartNetworkTraceOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type StartNetworkTraceOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartNetworkTraceOperationOptions() StartNetworkTraceOperationOptions {
	return StartNetworkTraceOperationOptions{}
}

func (o StartNetworkTraceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartNetworkTraceOperationOptions) toQueryString() map[string]interface{} {
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

// StartNetworkTrace ...
func (c WebAppsClient) StartNetworkTrace(ctx context.Context, id SiteId, options StartNetworkTraceOperationOptions) (result StartNetworkTraceOperationResponse, err error) {
	req, err := c.preparerForStartNetworkTrace(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartNetworkTrace", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartNetworkTrace(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartNetworkTrace", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartNetworkTraceThenPoll performs StartNetworkTrace then polls until it's completed
func (c WebAppsClient) StartNetworkTraceThenPoll(ctx context.Context, id SiteId, options StartNetworkTraceOperationOptions) error {
	result, err := c.StartNetworkTrace(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing StartNetworkTrace: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartNetworkTrace: %+v", err)
	}

	return nil
}

// preparerForStartNetworkTrace prepares the StartNetworkTrace request.
func (c WebAppsClient) preparerForStartNetworkTrace(ctx context.Context, id SiteId, options StartNetworkTraceOperationOptions) (*http.Request, error) {
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

// senderForStartNetworkTrace sends the StartNetworkTrace request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForStartNetworkTrace(ctx context.Context, req *http.Request) (future StartNetworkTraceOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
