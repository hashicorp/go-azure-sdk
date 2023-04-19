package codeversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryCodeVersionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *CodeVersionResource
}

// RegistryCodeVersionsGet ...
func (c CodeVersionClient) RegistryCodeVersionsGet(ctx context.Context, id RegistryCodeVersionId) (result RegistryCodeVersionsGetOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeVersionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryCodeVersionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryCodeVersionsGet prepares the RegistryCodeVersionsGet request.
func (c CodeVersionClient) preparerForRegistryCodeVersionsGet(ctx context.Context, id RegistryCodeVersionId) (*http.Request, error) {
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

// responderForRegistryCodeVersionsGet handles the response to the RegistryCodeVersionsGet request. The method always
// closes the http.Response Body.
func (c CodeVersionClient) responderForRegistryCodeVersionsGet(resp *http.Response) (result RegistryCodeVersionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
