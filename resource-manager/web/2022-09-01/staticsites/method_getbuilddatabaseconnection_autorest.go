package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBuildDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// GetBuildDatabaseConnection ...
func (c StaticSitesClient) GetBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId) (result GetBuildDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForGetBuildDatabaseConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBuildDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBuildDatabaseConnection prepares the GetBuildDatabaseConnection request.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId) (*http.Request, error) {
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

// responderForGetBuildDatabaseConnection handles the response to the GetBuildDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetBuildDatabaseConnection(resp *http.Response) (result GetBuildDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
