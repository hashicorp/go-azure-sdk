package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDatabaseConnectionWithDetailsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// GetDatabaseConnectionWithDetails ...
func (c StaticSitesClient) GetDatabaseConnectionWithDetails(ctx context.Context, id DatabaseConnectionId) (result GetDatabaseConnectionWithDetailsOperationResponse, err error) {
	req, err := c.preparerForGetDatabaseConnectionWithDetails(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnectionWithDetails", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnectionWithDetails", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDatabaseConnectionWithDetails(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnectionWithDetails", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDatabaseConnectionWithDetails prepares the GetDatabaseConnectionWithDetails request.
func (c StaticSitesClient) preparerForGetDatabaseConnectionWithDetails(ctx context.Context, id DatabaseConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/show", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDatabaseConnectionWithDetails handles the response to the GetDatabaseConnectionWithDetails request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetDatabaseConnectionWithDetails(resp *http.Response) (result GetDatabaseConnectionWithDetailsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
