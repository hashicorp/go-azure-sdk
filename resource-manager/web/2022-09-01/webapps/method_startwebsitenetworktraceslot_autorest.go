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

type StartWebSiteNetworkTraceSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *string
}

type StartWebSiteNetworkTraceSlotOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartWebSiteNetworkTraceSlotOperationOptions() StartWebSiteNetworkTraceSlotOperationOptions {
	return StartWebSiteNetworkTraceSlotOperationOptions{}
}

func (o StartWebSiteNetworkTraceSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o StartWebSiteNetworkTraceSlotOperationOptions) toQueryString() map[string]interface{} {
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

// StartWebSiteNetworkTraceSlot ...
func (c WebAppsClient) StartWebSiteNetworkTraceSlot(ctx context.Context, id SlotId, options StartWebSiteNetworkTraceSlotOperationOptions) (result StartWebSiteNetworkTraceSlotOperationResponse, err error) {
	req, err := c.preparerForStartWebSiteNetworkTraceSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartWebSiteNetworkTraceSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartWebSiteNetworkTraceSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartWebSiteNetworkTraceSlot prepares the StartWebSiteNetworkTraceSlot request.
func (c WebAppsClient) preparerForStartWebSiteNetworkTraceSlot(ctx context.Context, id SlotId, options StartWebSiteNetworkTraceSlotOperationOptions) (*http.Request, error) {
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

// responderForStartWebSiteNetworkTraceSlot handles the response to the StartWebSiteNetworkTraceSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStartWebSiteNetworkTraceSlot(resp *http.Response) (result StartWebSiteNetworkTraceSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
