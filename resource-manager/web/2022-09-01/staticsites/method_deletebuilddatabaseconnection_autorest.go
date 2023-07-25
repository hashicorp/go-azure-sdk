package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteBuildDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteBuildDatabaseConnection ...
func (c StaticSitesClient) DeleteBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId) (result DeleteBuildDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForDeleteBuildDatabaseConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteBuildDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteBuildDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteBuildDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteBuildDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteBuildDatabaseConnection prepares the DeleteBuildDatabaseConnection request.
func (c StaticSitesClient) preparerForDeleteBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId) (*http.Request, error) {
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

// responderForDeleteBuildDatabaseConnection handles the response to the DeleteBuildDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForDeleteBuildDatabaseConnection(resp *http.Response) (result DeleteBuildDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
