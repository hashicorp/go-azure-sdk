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

type ListAzureStorageAccountsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *AzureStoragePropertyDictionaryResource
}

// ListAzureStorageAccountsSlot ...
func (c WebAppsClient) ListAzureStorageAccountsSlot(ctx context.Context, id SlotId) (result ListAzureStorageAccountsSlotOperationResponse, err error) {
	req, err := c.preparerForListAzureStorageAccountsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccountsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccountsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListAzureStorageAccountsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListAzureStorageAccountsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListAzureStorageAccountsSlot prepares the ListAzureStorageAccountsSlot request.
func (c WebAppsClient) preparerForListAzureStorageAccountsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForListAzureStorageAccountsSlot handles the response to the ListAzureStorageAccountsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListAzureStorageAccountsSlot(resp *http.Response) (result ListAzureStorageAccountsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
