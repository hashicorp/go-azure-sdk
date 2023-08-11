package restorables

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDatabaseAccountsGetByLocationOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableDatabaseAccountGetResult
}

// RestorableDatabaseAccountsGetByLocation ...
func (c RestorablesClient) RestorableDatabaseAccountsGetByLocation(ctx context.Context, id RestorableDatabaseAccountId) (result RestorableDatabaseAccountsGetByLocationOperationResponse, err error) {
	req, err := c.preparerForRestorableDatabaseAccountsGetByLocation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableDatabaseAccountsGetByLocation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableDatabaseAccountsGetByLocation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableDatabaseAccountsGetByLocation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableDatabaseAccountsGetByLocation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableDatabaseAccountsGetByLocation prepares the RestorableDatabaseAccountsGetByLocation request.
func (c RestorablesClient) preparerForRestorableDatabaseAccountsGetByLocation(ctx context.Context, id RestorableDatabaseAccountId) (*http.Request, error) {
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

// responderForRestorableDatabaseAccountsGetByLocation handles the response to the RestorableDatabaseAccountsGetByLocation request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableDatabaseAccountsGetByLocation(resp *http.Response) (result RestorableDatabaseAccountsGetByLocationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
