package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestartWebAppsOperationResponse struct {
	HttpResponse *http.Response
}

type RestartWebAppsOperationOptions struct {
	SoftRestart *bool
}

func DefaultRestartWebAppsOperationOptions() RestartWebAppsOperationOptions {
	return RestartWebAppsOperationOptions{}
}

func (o RestartWebAppsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestartWebAppsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.SoftRestart != nil {
		out["softRestart"] = *o.SoftRestart
	}

	return out
}

// RestartWebApps ...
func (c AppServicePlansClient) RestartWebApps(ctx context.Context, id ServerFarmId, options RestartWebAppsOperationOptions) (result RestartWebAppsOperationResponse, err error) {
	req, err := c.preparerForRestartWebApps(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RestartWebApps", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RestartWebApps", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestartWebApps(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RestartWebApps", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestartWebApps prepares the RestartWebApps request.
func (c AppServicePlansClient) preparerForRestartWebApps(ctx context.Context, id ServerFarmId, options RestartWebAppsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restartSites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestartWebApps handles the response to the RestartWebApps request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForRestartWebApps(resp *http.Response) (result RestartWebAppsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
