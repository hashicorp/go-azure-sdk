package recoverablesqlpools

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerRecoverableSqlPoolsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecoverableSqlPool
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsGet ...
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsGet(ctx context.Context, id RecoverableSqlPoolId) (result WorkspaceManagedSqlServerRecoverableSqlPoolsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerRecoverableSqlPoolsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsGet prepares the WorkspaceManagedSqlServerRecoverableSqlPoolsGet request.
func (c RecoverableSqlPoolsClient) preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsGet(ctx context.Context, id RecoverableSqlPoolId) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerRecoverableSqlPoolsGet handles the response to the WorkspaceManagedSqlServerRecoverableSqlPoolsGet request. The method always
// closes the http.Response Body.
func (c RecoverableSqlPoolsClient) responderForWorkspaceManagedSqlServerRecoverableSqlPoolsGet(resp *http.Response) (result WorkspaceManagedSqlServerRecoverableSqlPoolsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
