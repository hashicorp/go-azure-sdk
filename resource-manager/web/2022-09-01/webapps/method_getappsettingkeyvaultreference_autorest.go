package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppSettingKeyVaultReferenceOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiKVReference
}

// GetAppSettingKeyVaultReference ...
func (c WebAppsClient) GetAppSettingKeyVaultReference(ctx context.Context, id AppSettingId) (result GetAppSettingKeyVaultReferenceOperationResponse, err error) {
	req, err := c.preparerForGetAppSettingKeyVaultReference(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReference", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReference", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAppSettingKeyVaultReference(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingKeyVaultReference", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAppSettingKeyVaultReference prepares the GetAppSettingKeyVaultReference request.
func (c WebAppsClient) preparerForGetAppSettingKeyVaultReference(ctx context.Context, id AppSettingId) (*http.Request, error) {
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

// responderForGetAppSettingKeyVaultReference handles the response to the GetAppSettingKeyVaultReference request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAppSettingKeyVaultReference(resp *http.Response) (result GetAppSettingKeyVaultReferenceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
