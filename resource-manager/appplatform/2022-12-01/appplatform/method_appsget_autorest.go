package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *AppResource
}

type AppsGetOperationOptions struct {
	SyncStatus *string
}

func DefaultAppsGetOperationOptions() AppsGetOperationOptions {
	return AppsGetOperationOptions{}
}

func (o AppsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AppsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.SyncStatus != nil {
		out["syncStatus"] = *o.SyncStatus
	}

	return out
}

// AppsGet ...
func (c AppPlatformClient) AppsGet(ctx context.Context, id AppId, options AppsGetOperationOptions) (result AppsGetOperationResponse, err error) {
	req, err := c.preparerForAppsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAppsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAppsGet prepares the AppsGet request.
func (c AppPlatformClient) preparerForAppsGet(ctx context.Context, id AppId, options AppsGetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAppsGet handles the response to the AppsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForAppsGet(resp *http.Response) (result AppsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
