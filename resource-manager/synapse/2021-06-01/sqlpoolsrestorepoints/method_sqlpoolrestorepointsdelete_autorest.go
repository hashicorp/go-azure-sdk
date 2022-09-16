package sqlpoolsrestorepoints

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolRestorePointsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolRestorePointsDelete ...
func (c SqlPoolsRestorePointsClient) SqlPoolRestorePointsDelete(ctx context.Context, id RestorePointId) (result SqlPoolRestorePointsDeleteOperationResponse, err error) {
	req, err := c.preparerForSqlPoolRestorePointsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolRestorePointsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolRestorePointsDelete prepares the SqlPoolRestorePointsDelete request.
func (c SqlPoolsRestorePointsClient) preparerForSqlPoolRestorePointsDelete(ctx context.Context, id RestorePointId) (*http.Request, error) {
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

// responderForSqlPoolRestorePointsDelete handles the response to the SqlPoolRestorePointsDelete request. The method always
// closes the http.Response Body.
func (c SqlPoolsRestorePointsClient) responderForSqlPoolRestorePointsDelete(resp *http.Response) (result SqlPoolRestorePointsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
