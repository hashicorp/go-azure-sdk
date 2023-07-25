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

type StartWebSiteNetworkTraceOperationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type StartWebSiteNetworkTraceOperationOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartWebSiteNetworkTraceOperationOperationOptions() StartWebSiteNetworkTraceOperationOperationOptions {
	return StartWebSiteNetworkTraceOperationOperationOptions{}
}

func (o StartWebSiteNetworkTraceOperationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartWebSiteNetworkTraceOperationOperationOptions) toQueryString() map[string]interface{} {
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

// StartWebSiteNetworkTraceOperation ...
func (c WebAppsClient) StartWebSiteNetworkTraceOperation(ctx context.Context, id SiteId, options StartWebSiteNetworkTraceOperationOperationOptions) (result StartWebSiteNetworkTraceOperationOperationResponse, err error) {
	req, err := c.preparerForStartWebSiteNetworkTraceOperation(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceOperation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartWebSiteNetworkTraceOperation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceOperation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartWebSiteNetworkTraceOperationThenPoll performs StartWebSiteNetworkTraceOperation then polls until it's completed
func (c WebAppsClient) StartWebSiteNetworkTraceOperationThenPoll(ctx context.Context, id SiteId, options StartWebSiteNetworkTraceOperationOperationOptions) error {
	result, err := c.StartWebSiteNetworkTraceOperation(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing StartWebSiteNetworkTraceOperation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartWebSiteNetworkTraceOperation: %+v", err)
	}

	return nil
}

// preparerForStartWebSiteNetworkTraceOperation prepares the StartWebSiteNetworkTraceOperation request.
func (c WebAppsClient) preparerForStartWebSiteNetworkTraceOperation(ctx context.Context, id SiteId, options StartWebSiteNetworkTraceOperationOperationOptions) (*http.Request, error) {
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

// senderForStartWebSiteNetworkTraceOperation sends the StartWebSiteNetworkTraceOperation request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForStartWebSiteNetworkTraceOperation(ctx context.Context, req *http.Request) (future StartWebSiteNetworkTraceOperationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
