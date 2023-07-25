package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceProcessOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessInfo
}

// GetInstanceProcess ...
func (c WebAppsClient) GetInstanceProcess(ctx context.Context, id InstanceProcessId) (result GetInstanceProcessOperationResponse, err error) {
	req, err := c.preparerForGetInstanceProcess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceProcess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceProcess prepares the GetInstanceProcess request.
func (c WebAppsClient) preparerForGetInstanceProcess(ctx context.Context, id InstanceProcessId) (*http.Request, error) {
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

// responderForGetInstanceProcess handles the response to the GetInstanceProcess request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceProcess(resp *http.Response) (result GetInstanceProcessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
