package dataversionregistry

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryDataVersionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataVersionBaseResource
}

// RegistryDataVersionsGet ...
func (c DataVersionRegistryClient) RegistryDataVersionsGet(ctx context.Context, id VersionId) (result RegistryDataVersionsGetOperationResponse, err error) {
	req, err := c.preparerForRegistryDataVersionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryDataVersionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryDataVersionsGet prepares the RegistryDataVersionsGet request.
func (c DataVersionRegistryClient) preparerForRegistryDataVersionsGet(ctx context.Context, id VersionId) (*http.Request, error) {
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

// responderForRegistryDataVersionsGet handles the response to the RegistryDataVersionsGet request. The method always
// closes the http.Response Body.
func (c DataVersionRegistryClient) responderForRegistryDataVersionsGet(resp *http.Response) (result RegistryDataVersionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
