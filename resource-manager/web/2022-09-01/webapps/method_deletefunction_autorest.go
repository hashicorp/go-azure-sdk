package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteFunctionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteFunction ...
func (c WebAppsClient) DeleteFunction(ctx context.Context, id FunctionId) (result DeleteFunctionOperationResponse, err error) {
	req, err := c.preparerForDeleteFunction(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunction", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunction", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteFunction(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunction", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteFunction prepares the DeleteFunction request.
func (c WebAppsClient) preparerForDeleteFunction(ctx context.Context, id FunctionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteFunction handles the response to the DeleteFunction request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteFunction(resp *http.Response) (result DeleteFunctionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
