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

type RestartOperationResponse struct {
	HttpResponse *http.Response
}

type RestartOperationOptions struct {
	SoftRestart *bool
	Synchronous *bool
}

func DefaultRestartOperationOptions() RestartOperationOptions {
	return RestartOperationOptions{}
}

func (o RestartOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestartOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.SoftRestart != nil {
		out["softRestart"] = *o.SoftRestart
	}

	if o.Synchronous != nil {
		out["synchronous"] = *o.Synchronous
	}

	return out
}

// Restart ...
func (c WebAppsClient) Restart(ctx context.Context, id SiteId, options RestartOperationOptions) (result RestartOperationResponse, err error) {
	req, err := c.preparerForRestart(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Restart", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Restart", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestart(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Restart", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestart prepares the Restart request.
func (c WebAppsClient) preparerForRestart(ctx context.Context, id SiteId, options RestartOperationOptions) (*http.Request, error) {
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

// responderForRestart handles the response to the Restart request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRestart(resp *http.Response) (result RestartOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
