package registrymanagement

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistriesUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegistryTrackedResource
}

// RegistriesUpdate ...
func (c RegistryManagementClient) RegistriesUpdate(ctx context.Context, id RegistryId, input PartialRegistryPartialTrackedResource) (result RegistriesUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistriesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistriesUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistriesUpdate prepares the RegistriesUpdate request.
func (c RegistryManagementClient) preparerForRegistriesUpdate(ctx context.Context, id RegistryId, input PartialRegistryPartialTrackedResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegistriesUpdate handles the response to the RegistriesUpdate request. The method always
// closes the http.Response Body.
func (c RegistryManagementClient) responderForRegistriesUpdate(resp *http.Response) (result RegistriesUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
