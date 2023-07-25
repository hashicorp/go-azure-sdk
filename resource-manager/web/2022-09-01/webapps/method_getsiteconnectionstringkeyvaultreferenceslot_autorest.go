package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteConnectionStringKeyVaultReferenceSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiKVReference
}

// GetSiteConnectionStringKeyVaultReferenceSlot ...
func (c WebAppsClient) GetSiteConnectionStringKeyVaultReferenceSlot(ctx context.Context, id ConfigReferenceConnectionStringId) (result GetSiteConnectionStringKeyVaultReferenceSlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteConnectionStringKeyVaultReferenceSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferenceSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferenceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteConnectionStringKeyVaultReferenceSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferenceSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteConnectionStringKeyVaultReferenceSlot prepares the GetSiteConnectionStringKeyVaultReferenceSlot request.
func (c WebAppsClient) preparerForGetSiteConnectionStringKeyVaultReferenceSlot(ctx context.Context, id ConfigReferenceConnectionStringId) (*http.Request, error) {
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

// responderForGetSiteConnectionStringKeyVaultReferenceSlot handles the response to the GetSiteConnectionStringKeyVaultReferenceSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSiteConnectionStringKeyVaultReferenceSlot(resp *http.Response) (result GetSiteConnectionStringKeyVaultReferenceSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
