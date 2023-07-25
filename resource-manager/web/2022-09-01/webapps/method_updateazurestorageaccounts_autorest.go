package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAzureStorageAccountsOperationResponse struct {
	HttpResponse *http.Response
	Model        *AzureStoragePropertyDictionaryResource
}

// UpdateAzureStorageAccounts ...
func (c WebAppsClient) UpdateAzureStorageAccounts(ctx context.Context, id SiteId, input AzureStoragePropertyDictionaryResource) (result UpdateAzureStorageAccountsOperationResponse, err error) {
	req, err := c.preparerForUpdateAzureStorageAccounts(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAzureStorageAccounts", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAzureStorageAccounts", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAzureStorageAccounts(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAzureStorageAccounts", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAzureStorageAccounts prepares the UpdateAzureStorageAccounts request.
func (c WebAppsClient) preparerForUpdateAzureStorageAccounts(ctx context.Context, id SiteId, input AzureStoragePropertyDictionaryResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/azurestorageaccounts", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateAzureStorageAccounts handles the response to the UpdateAzureStorageAccounts request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateAzureStorageAccounts(resp *http.Response) (result UpdateAzureStorageAccountsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
