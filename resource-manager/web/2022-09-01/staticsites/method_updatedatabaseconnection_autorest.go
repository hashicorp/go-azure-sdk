package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// UpdateDatabaseConnection ...
func (c StaticSitesClient) UpdateDatabaseConnection(ctx context.Context, id DatabaseConnectionId, input DatabaseConnectionPatchRequest) (result UpdateDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForUpdateDatabaseConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateDatabaseConnection prepares the UpdateDatabaseConnection request.
func (c StaticSitesClient) preparerForUpdateDatabaseConnection(ctx context.Context, id DatabaseConnectionId, input DatabaseConnectionPatchRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateDatabaseConnection handles the response to the UpdateDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUpdateDatabaseConnection(resp *http.Response) (result UpdateDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
