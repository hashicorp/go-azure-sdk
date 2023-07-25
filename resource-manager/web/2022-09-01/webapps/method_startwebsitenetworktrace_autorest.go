package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartWebSiteNetworkTraceOperationResponse struct {
	HttpResponse *http.Response
	Model        *string
}

type StartWebSiteNetworkTraceOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartWebSiteNetworkTraceOperationOptions() StartWebSiteNetworkTraceOperationOptions {
	return StartWebSiteNetworkTraceOperationOptions{}
}

func (o StartWebSiteNetworkTraceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartWebSiteNetworkTraceOperationOptions) toQueryString() map[string]interface{} {
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

// StartWebSiteNetworkTrace ...
func (c WebAppsClient) StartWebSiteNetworkTrace(ctx context.Context, id SiteId, options StartWebSiteNetworkTraceOperationOptions) (result StartWebSiteNetworkTraceOperationResponse, err error) {
	req, err := c.preparerForStartWebSiteNetworkTrace(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTrace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTrace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartWebSiteNetworkTrace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTrace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartWebSiteNetworkTrace prepares the StartWebSiteNetworkTrace request.
func (c WebAppsClient) preparerForStartWebSiteNetworkTrace(ctx context.Context, id SiteId, options StartWebSiteNetworkTraceOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/networkTrace/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStartWebSiteNetworkTrace handles the response to the StartWebSiteNetworkTrace request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStartWebSiteNetworkTrace(resp *http.Response) (result StartWebSiteNetworkTraceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
