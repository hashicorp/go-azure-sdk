package registrymanagement

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistriesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegistryTrackedResource
}

// RegistriesGet ...
func (c RegistryManagementClient) RegistriesGet(ctx context.Context, id RegistryId) (result RegistriesGetOperationResponse, err error) {
	req, err := c.preparerForRegistriesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistriesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistriesGet prepares the RegistriesGet request.
func (c RegistryManagementClient) preparerForRegistriesGet(ctx context.Context, id RegistryId) (*http.Request, error) {
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

// responderForRegistriesGet handles the response to the RegistriesGet request. The method always
// closes the http.Response Body.
func (c RegistryManagementClient) responderForRegistriesGet(resp *http.Response) (result RegistriesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
