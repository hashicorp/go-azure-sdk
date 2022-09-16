package sqlpoolssqlpooltables

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolTablesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Resource
}

// SqlPoolTablesGet ...
func (c SqlPoolsSqlPoolTablesClient) SqlPoolTablesGet(ctx context.Context, id TableId) (result SqlPoolTablesGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolTablesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooltables.SqlPoolsSqlPoolTablesClient", "SqlPoolTablesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooltables.SqlPoolsSqlPoolTablesClient", "SqlPoolTablesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolTablesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooltables.SqlPoolsSqlPoolTablesClient", "SqlPoolTablesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolTablesGet prepares the SqlPoolTablesGet request.
func (c SqlPoolsSqlPoolTablesClient) preparerForSqlPoolTablesGet(ctx context.Context, id TableId) (*http.Request, error) {
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

// responderForSqlPoolTablesGet handles the response to the SqlPoolTablesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSqlPoolTablesClient) responderForSqlPoolTablesGet(resp *http.Response) (result SqlPoolTablesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
