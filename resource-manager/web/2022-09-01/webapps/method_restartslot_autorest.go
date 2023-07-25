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

type RestartSlotOperationResponse struct {
	HttpResponse *http.Response
}

type RestartSlotOperationOptions struct {
	SoftRestart *bool
	Synchronous *bool
}

func DefaultRestartSlotOperationOptions() RestartSlotOperationOptions {
	return RestartSlotOperationOptions{}
}

func (o RestartSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestartSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.SoftRestart != nil {
		out["softRestart"] = *o.SoftRestart
	}

	if o.Synchronous != nil {
		out["synchronous"] = *o.Synchronous
	}

	return out
}

// RestartSlot ...
func (c WebAppsClient) RestartSlot(ctx context.Context, id SlotId, options RestartSlotOperationOptions) (result RestartSlotOperationResponse, err error) {
	req, err := c.preparerForRestartSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestartSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestartSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestartSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestartSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestartSlot prepares the RestartSlot request.
func (c WebAppsClient) preparerForRestartSlot(ctx context.Context, id SlotId, options RestartSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestartSlot handles the response to the RestartSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRestartSlot(resp *http.Response) (result RestartSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
