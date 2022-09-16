package workspacemanagedsqlserver

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableSqlPoolsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecoverableSqlPool
}

// RecoverableSqlPoolsGet ...
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsGet(ctx context.Context, id RecoverableSqlPoolId) (result RecoverableSqlPoolsGetOperationResponse, err error) {
	req, err := c.preparerForRecoverableSqlPoolsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecoverableSqlPoolsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecoverableSqlPoolsGet prepares the RecoverableSqlPoolsGet request.
func (c WorkspaceManagedSqlServerClient) preparerForRecoverableSqlPoolsGet(ctx context.Context, id RecoverableSqlPoolId) (*http.Request, error) {
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

// responderForRecoverableSqlPoolsGet handles the response to the RecoverableSqlPoolsGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerClient) responderForRecoverableSqlPoolsGet(resp *http.Response) (result RecoverableSqlPoolsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
