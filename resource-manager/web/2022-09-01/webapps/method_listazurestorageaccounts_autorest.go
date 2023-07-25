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

type ListAzureStorageAccountsOperationResponse struct {
	HttpResponse *http.Response
	Model        *AzureStoragePropertyDictionaryResource
}

// ListAzureStorageAccounts ...
func (c WebAppsClient) ListAzureStorageAccounts(ctx context.Context, id SiteId) (result ListAzureStorageAccountsOperationResponse, err error) {
	req, err := c.preparerForListAzureStorageAccounts(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccounts", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccounts", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListAzureStorageAccounts(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccounts", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListAzureStorageAccounts prepares the ListAzureStorageAccounts request.
func (c WebAppsClient) preparerForListAzureStorageAccounts(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/azurestorageaccounts/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListAzureStorageAccounts handles the response to the ListAzureStorageAccounts request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListAzureStorageAccounts(resp *http.Response) (result ListAzureStorageAccountsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
