package capabilities

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityTypesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *CapabilityType
}

// CapabilityTypesGet ...
func (c CapabilitiesClient) CapabilityTypesGet(ctx context.Context, id CapabilityTypeId) (result CapabilityTypesGetOperationResponse, err error) {
	req, err := c.preparerForCapabilityTypesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capabilities.CapabilitiesClient", "CapabilityTypesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "capabilities.CapabilitiesClient", "CapabilityTypesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCapabilityTypesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capabilities.CapabilitiesClient", "CapabilityTypesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCapabilityTypesGet prepares the CapabilityTypesGet request.
func (c CapabilitiesClient) preparerForCapabilityTypesGet(ctx context.Context, id CapabilityTypeId) (*http.Request, error) {
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

// responderForCapabilityTypesGet handles the response to the CapabilityTypesGet request. The method always
// closes the http.Response Body.
func (c CapabilitiesClient) responderForCapabilityTypesGet(resp *http.Response) (result CapabilityTypesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
