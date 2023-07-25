package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// GetDatabaseConnection ...
func (c StaticSitesClient) GetDatabaseConnection(ctx context.Context, id DatabaseConnectionId) (result GetDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForGetDatabaseConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDatabaseConnection prepares the GetDatabaseConnection request.
func (c StaticSitesClient) preparerForGetDatabaseConnection(ctx context.Context, id DatabaseConnectionId) (*http.Request, error) {
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

// responderForGetDatabaseConnection handles the response to the GetDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetDatabaseConnection(resp *http.Response) (result GetDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
