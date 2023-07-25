package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHostSecretOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHostSecret ...
func (c WebAppsClient) DeleteHostSecret(ctx context.Context, id DefaultId) (result DeleteHostSecretOperationResponse, err error) {
	req, err := c.preparerForDeleteHostSecret(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecret", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecret", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHostSecret(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecret", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHostSecret prepares the DeleteHostSecret request.
func (c WebAppsClient) preparerForDeleteHostSecret(ctx context.Context, id DefaultId) (*http.Request, error) {
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

// responderForDeleteHostSecret handles the response to the DeleteHostSecret request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHostSecret(resp *http.Response) (result DeleteHostSecretOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
