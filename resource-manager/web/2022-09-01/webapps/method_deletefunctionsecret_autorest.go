package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteFunctionSecretOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteFunctionSecret ...
func (c WebAppsClient) DeleteFunctionSecret(ctx context.Context, id KeyId) (result DeleteFunctionSecretOperationResponse, err error) {
	req, err := c.preparerForDeleteFunctionSecret(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecret", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecret", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteFunctionSecret(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecret", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteFunctionSecret prepares the DeleteFunctionSecret request.
func (c WebAppsClient) preparerForDeleteFunctionSecret(ctx context.Context, id KeyId) (*http.Request, error) {
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

// responderForDeleteFunctionSecret handles the response to the DeleteFunctionSecret request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteFunctionSecret(resp *http.Response) (result DeleteFunctionSecretOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
