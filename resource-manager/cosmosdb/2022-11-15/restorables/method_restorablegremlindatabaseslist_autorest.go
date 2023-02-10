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

type RestorableGremlinDatabasesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableGremlinDatabasesListResult
}

// RestorableGremlinDatabasesList ...
func (c RestorablesClient) RestorableGremlinDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (result RestorableGremlinDatabasesListOperationResponse, err error) {
	req, err := c.preparerForRestorableGremlinDatabasesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinDatabasesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinDatabasesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableGremlinDatabasesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinDatabasesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableGremlinDatabasesList prepares the RestorableGremlinDatabasesList request.
func (c RestorablesClient) preparerForRestorableGremlinDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restorableGremlinDatabases", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableGremlinDatabasesList handles the response to the RestorableGremlinDatabasesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableGremlinDatabasesList(resp *http.Response) (result RestorableGremlinDatabasesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
