package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlDatabasesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableSqlDatabasesListResult
}

// RestorableSqlDatabasesList ...
func (c RestorablesClient) RestorableSqlDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (result RestorableSqlDatabasesListOperationResponse, err error) {
	req, err := c.preparerForRestorableSqlDatabasesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlDatabasesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlDatabasesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableSqlDatabasesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlDatabasesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableSqlDatabasesList prepares the RestorableSqlDatabasesList request.
func (c RestorablesClient) preparerForRestorableSqlDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restorableSqlDatabases", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableSqlDatabasesList handles the response to the RestorableSqlDatabasesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableSqlDatabasesList(resp *http.Response) (result RestorableSqlDatabasesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
