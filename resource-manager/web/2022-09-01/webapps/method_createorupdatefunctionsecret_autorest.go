package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateFunctionSecretOperationResponse struct {
	HttpResponse *http.Response
	Model        *KeyInfo
}

// CreateOrUpdateFunctionSecret ...
func (c WebAppsClient) CreateOrUpdateFunctionSecret(ctx context.Context, id KeyId, input KeyInfo) (result CreateOrUpdateFunctionSecretOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateFunctionSecret(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateFunctionSecret", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateFunctionSecret", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateFunctionSecret(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateFunctionSecret", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateFunctionSecret prepares the CreateOrUpdateFunctionSecret request.
func (c WebAppsClient) preparerForCreateOrUpdateFunctionSecret(ctx context.Context, id KeyId, input KeyInfo) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateFunctionSecret handles the response to the CreateOrUpdateFunctionSecret request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateFunctionSecret(resp *http.Response) (result CreateOrUpdateFunctionSecretOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
