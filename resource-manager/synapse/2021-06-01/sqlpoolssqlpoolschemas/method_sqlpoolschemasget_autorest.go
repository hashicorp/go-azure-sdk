package sqlpoolssqlpoolschemas

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSchemasGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Resource
}

// SqlPoolSchemasGet ...
func (c SqlPoolsSqlPoolSchemasClient) SqlPoolSchemasGet(ctx context.Context, id SchemaId) (result SqlPoolSchemasGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSchemasGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolschemas.SqlPoolsSqlPoolSchemasClient", "SqlPoolSchemasGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolschemas.SqlPoolsSqlPoolSchemasClient", "SqlPoolSchemasGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSchemasGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpoolschemas.SqlPoolsSqlPoolSchemasClient", "SqlPoolSchemasGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSchemasGet prepares the SqlPoolSchemasGet request.
func (c SqlPoolsSqlPoolSchemasClient) preparerForSqlPoolSchemasGet(ctx context.Context, id SchemaId) (*http.Request, error) {
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

// responderForSqlPoolSchemasGet handles the response to the SqlPoolSchemasGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSqlPoolSchemasClient) responderForSqlPoolSchemasGet(resp *http.Response) (result SqlPoolSchemasGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
