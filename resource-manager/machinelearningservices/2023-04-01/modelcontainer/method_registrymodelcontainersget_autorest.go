package modelcontainer

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryModelContainersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ModelContainerResource
}

// RegistryModelContainersGet ...
func (c ModelContainerClient) RegistryModelContainersGet(ctx context.Context, id RegistryModelId) (result RegistryModelContainersGetOperationResponse, err error) {
	req, err := c.preparerForRegistryModelContainersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryModelContainersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryModelContainersGet prepares the RegistryModelContainersGet request.
func (c ModelContainerClient) preparerForRegistryModelContainersGet(ctx context.Context, id RegistryModelId) (*http.Request, error) {
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

// responderForRegistryModelContainersGet handles the response to the RegistryModelContainersGet request. The method always
// closes the http.Response Body.
func (c ModelContainerClient) responderForRegistryModelContainersGet(resp *http.Response) (result RegistryModelContainersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
