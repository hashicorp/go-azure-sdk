package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteProcessOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteProcess ...
func (c WebAppsClient) DeleteProcess(ctx context.Context, id ProcessId) (result DeleteProcessOperationResponse, err error) {
	req, err := c.preparerForDeleteProcess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteProcess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteProcess prepares the DeleteProcess request.
func (c WebAppsClient) preparerForDeleteProcess(ctx context.Context, id ProcessId) (*http.Request, error) {
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

// responderForDeleteProcess handles the response to the DeleteProcess request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteProcess(resp *http.Response) (result DeleteProcessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
