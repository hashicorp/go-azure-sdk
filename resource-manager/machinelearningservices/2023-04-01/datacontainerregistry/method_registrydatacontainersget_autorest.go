package datacontainerregistry

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryDataContainersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataContainerResource
}

// RegistryDataContainersGet ...
func (c DataContainerRegistryClient) RegistryDataContainersGet(ctx context.Context, id DataId) (result RegistryDataContainersGetOperationResponse, err error) {
	req, err := c.preparerForRegistryDataContainersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryDataContainersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryDataContainersGet prepares the RegistryDataContainersGet request.
func (c DataContainerRegistryClient) preparerForRegistryDataContainersGet(ctx context.Context, id DataId) (*http.Request, error) {
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

// responderForRegistryDataContainersGet handles the response to the RegistryDataContainersGet request. The method always
// closes the http.Response Body.
func (c DataContainerRegistryClient) responderForRegistryDataContainersGet(resp *http.Response) (result RegistryDataContainersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
