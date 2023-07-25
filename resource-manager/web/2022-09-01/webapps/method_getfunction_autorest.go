package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetFunctionOperationResponse struct {
	HttpResponse *http.Response
	Model        *FunctionEnvelope
}

// GetFunction ...
func (c WebAppsClient) GetFunction(ctx context.Context, id FunctionId) (result GetFunctionOperationResponse, err error) {
	req, err := c.preparerForGetFunction(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunction", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunction", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFunction(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunction", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFunction prepares the GetFunction request.
func (c WebAppsClient) preparerForGetFunction(ctx context.Context, id FunctionId) (*http.Request, error) {
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

// responderForGetFunction handles the response to the GetFunction request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetFunction(resp *http.Response) (result GetFunctionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
