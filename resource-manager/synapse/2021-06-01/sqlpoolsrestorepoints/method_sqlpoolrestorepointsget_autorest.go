package sqlpoolsrestorepoints

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolRestorePointsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorePoint
}

// SqlPoolRestorePointsGet ...
func (c SqlPoolsRestorePointsClient) SqlPoolRestorePointsGet(ctx context.Context, id RestorePointId) (result SqlPoolRestorePointsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolRestorePointsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolRestorePointsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsrestorepoints.SqlPoolsRestorePointsClient", "SqlPoolRestorePointsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolRestorePointsGet prepares the SqlPoolRestorePointsGet request.
func (c SqlPoolsRestorePointsClient) preparerForSqlPoolRestorePointsGet(ctx context.Context, id RestorePointId) (*http.Request, error) {
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

// responderForSqlPoolRestorePointsGet handles the response to the SqlPoolRestorePointsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsRestorePointsClient) responderForSqlPoolRestorePointsGet(resp *http.Response) (result SqlPoolRestorePointsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
