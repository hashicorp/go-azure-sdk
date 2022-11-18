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

type RestorableMongodbDatabasesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableMongodbDatabasesListResult
}

// RestorableMongodbDatabasesList ...
func (c RestorablesClient) RestorableMongodbDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (result RestorableMongodbDatabasesListOperationResponse, err error) {
	req, err := c.preparerForRestorableMongodbDatabasesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbDatabasesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbDatabasesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableMongodbDatabasesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbDatabasesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableMongodbDatabasesList prepares the RestorableMongodbDatabasesList request.
func (c RestorablesClient) preparerForRestorableMongodbDatabasesList(ctx context.Context, id RestorableDatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restorableMongodbDatabases", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableMongodbDatabasesList handles the response to the RestorableMongodbDatabasesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableMongodbDatabasesList(resp *http.Response) (result RestorableMongodbDatabasesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
