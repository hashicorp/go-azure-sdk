package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppSettingKeyVaultReferenceSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiKVReference
}

// GetAppSettingKeyVaultReferenceSlot ...
func (c WebAppsClient) GetAppSettingKeyVaultReferenceSlot(ctx context.Context, id ConfigReferenceAppSettingId) (result GetAppSettingKeyVaultReferenceSlotOperationResponse, err error) {
	req, err := c.preparerForGetAppSettingKeyVaultReferenceSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReferenceSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReferenceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAppSettingKeyVaultReferenceSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReferenceSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAppSettingKeyVaultReferenceSlot prepares the GetAppSettingKeyVaultReferenceSlot request.
func (c WebAppsClient) preparerForGetAppSettingKeyVaultReferenceSlot(ctx context.Context, id ConfigReferenceAppSettingId) (*http.Request, error) {
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

// responderForGetAppSettingKeyVaultReferenceSlot handles the response to the GetAppSettingKeyVaultReferenceSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAppSettingKeyVaultReferenceSlot(resp *http.Response) (result GetAppSettingKeyVaultReferenceSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
