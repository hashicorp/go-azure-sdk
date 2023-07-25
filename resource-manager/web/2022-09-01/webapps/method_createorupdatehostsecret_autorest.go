package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateHostSecretOperationResponse struct {
	HttpResponse *http.Response
	Model        *KeyInfo
}

// CreateOrUpdateHostSecret ...
func (c WebAppsClient) CreateOrUpdateHostSecret(ctx context.Context, id DefaultId, input KeyInfo) (result CreateOrUpdateHostSecretOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateHostSecret(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostSecret", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostSecret", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateHostSecret(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostSecret", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateHostSecret prepares the CreateOrUpdateHostSecret request.
func (c WebAppsClient) preparerForCreateOrUpdateHostSecret(ctx context.Context, id DefaultId, input KeyInfo) (*http.Request, error) {
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

// responderForCreateOrUpdateHostSecret handles the response to the CreateOrUpdateHostSecret request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateHostSecret(resp *http.Response) (result CreateOrUpdateHostSecretOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
