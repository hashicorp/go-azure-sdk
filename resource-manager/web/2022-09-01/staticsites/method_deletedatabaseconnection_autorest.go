package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteDatabaseConnection ...
func (c StaticSitesClient) DeleteDatabaseConnection(ctx context.Context, id DatabaseConnectionId) (result DeleteDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForDeleteDatabaseConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteDatabaseConnection prepares the DeleteDatabaseConnection request.
func (c StaticSitesClient) preparerForDeleteDatabaseConnection(ctx context.Context, id DatabaseConnectionId) (*http.Request, error) {
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

// responderForDeleteDatabaseConnection handles the response to the DeleteDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForDeleteDatabaseConnection(resp *http.Response) (result DeleteDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
