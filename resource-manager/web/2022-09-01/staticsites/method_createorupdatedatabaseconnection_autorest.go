package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateDatabaseConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// CreateOrUpdateDatabaseConnection ...
func (c StaticSitesClient) CreateOrUpdateDatabaseConnection(ctx context.Context, id DatabaseConnectionId, input DatabaseConnection) (result CreateOrUpdateDatabaseConnectionOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateDatabaseConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateDatabaseConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateDatabaseConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateDatabaseConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateDatabaseConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateDatabaseConnection prepares the CreateOrUpdateDatabaseConnection request.
func (c StaticSitesClient) preparerForCreateOrUpdateDatabaseConnection(ctx context.Context, id DatabaseConnectionId, input DatabaseConnection) (*http.Request, error) {
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

// responderForCreateOrUpdateDatabaseConnection handles the response to the CreateOrUpdateDatabaseConnection request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateDatabaseConnection(resp *http.Response) (result CreateOrUpdateDatabaseConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
