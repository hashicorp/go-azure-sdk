package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetProcessOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessInfo
}

// GetProcess ...
func (c WebAppsClient) GetProcess(ctx context.Context, id ProcessId) (result GetProcessOperationResponse, err error) {
	req, err := c.preparerForGetProcess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcess prepares the GetProcess request.
func (c WebAppsClient) preparerForGetProcess(ctx context.Context, id ProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetProcess handles the response to the GetProcess request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcess(resp *http.Response) (result GetProcessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
