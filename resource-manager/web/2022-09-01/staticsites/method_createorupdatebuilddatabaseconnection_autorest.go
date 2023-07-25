package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateBuildDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// CreateOrUpdateBuildDatabaseConnection ...
func (c StaticSitesClient) CreateOrUpdateBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId, input DatabaseConnection) (result CreateOrUpdateBuildDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateBuildDatabaseConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBuildDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBuildDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateBuildDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBuildDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateBuildDatabaseConnection prepares the CreateOrUpdateBuildDatabaseConnection request.
func (c StaticSitesClient) preparerForCreateOrUpdateBuildDatabaseConnection(ctx context.Context, id BuildDatabaseConnectionId, input DatabaseConnection) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateBuildDatabaseConnection handles the response to the CreateOrUpdateBuildDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateBuildDatabaseConnection(resp *http.Response) (result CreateOrUpdateBuildDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
