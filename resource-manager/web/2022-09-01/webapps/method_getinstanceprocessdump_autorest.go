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

type GetInstanceProcessDumpOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetInstanceProcessDump ...
func (c WebAppsClient) GetInstanceProcessDump(ctx context.Context, id InstanceProcessId) (result GetInstanceProcessDumpOperationResponse, err error) {
	req, err := c.preparerForGetInstanceProcessDump(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDump", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDump", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceProcessDump(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDump", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceProcessDump prepares the GetInstanceProcessDump request.
func (c WebAppsClient) preparerForGetInstanceProcessDump(ctx context.Context, id InstanceProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dump", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInstanceProcessDump handles the response to the GetInstanceProcessDump request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceProcessDump(resp *http.Response) (result GetInstanceProcessDumpOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
