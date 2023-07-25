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

type GetBuildDatabaseConnectionWithDetailsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DatabaseConnection
}

// GetBuildDatabaseConnectionWithDetails ...
func (c StaticSitesClient) GetBuildDatabaseConnectionWithDetails(ctx context.Context, id BuildDatabaseConnectionId) (result GetBuildDatabaseConnectionWithDetailsOperationResponse, err error) {
	req, err := c.preparerForGetBuildDatabaseConnectionWithDetails(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionWithDetails", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionWithDetails", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBuildDatabaseConnectionWithDetails(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionWithDetails", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBuildDatabaseConnectionWithDetails prepares the GetBuildDatabaseConnectionWithDetails request.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnectionWithDetails(ctx context.Context, id BuildDatabaseConnectionId) (*http.Request, error) {
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

// responderForGetBuildDatabaseConnectionWithDetails handles the response to the GetBuildDatabaseConnectionWithDetails request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetBuildDatabaseConnectionWithDetails(resp *http.Response) (result GetBuildDatabaseConnectionWithDetailsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
