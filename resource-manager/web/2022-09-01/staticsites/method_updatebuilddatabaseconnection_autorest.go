package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateBuildDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// UpdateBuildDatabaseConnection ...
func (c StaticSitesClient) UpdateBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId, input DatabaseConnectionPatchRequest) (result UpdateBuildDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForUpdateBuildDatabaseConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateBuildDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateBuildDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateBuildDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateBuildDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateBuildDatabaseConnection prepares the UpdateBuildDatabaseConnection request.
func (c StaticSitesClient) preparerForUpdateBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId, input DatabaseConnectionPatchRequest) (*http.Request, error) {
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

// responderForUpdateBuildDatabaseConnection handles the response to the UpdateBuildDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUpdateBuildDatabaseConnection(resp *http.Response) (result UpdateBuildDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
