package sqlpoolssqlpoolcolumns

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolColumnsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlPoolColumn
}

// SqlPoolColumnsGet ...
func (c SqlPoolsSqlPoolColumnsClient) SqlPoolColumnsGet(ctx context.Context, id ColumnId) (result SqlPoolColumnsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolColumnsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolcolumns.SqlPoolsSqlPoolColumnsClient", "SqlPoolColumnsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolcolumns.SqlPoolsSqlPoolColumnsClient", "SqlPoolColumnsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolColumnsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolcolumns.SqlPoolsSqlPoolColumnsClient", "SqlPoolColumnsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolColumnsGet prepares the SqlPoolColumnsGet request.
func (c SqlPoolsSqlPoolColumnsClient) preparerForSqlPoolColumnsGet(ctx context.Context, id ColumnId) (*http.Request, error) {
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

// responderForSqlPoolColumnsGet handles the response to the SqlPoolColumnsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSqlPoolColumnsClient) responderForSqlPoolColumnsGet(resp *http.Response) (result SqlPoolColumnsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
