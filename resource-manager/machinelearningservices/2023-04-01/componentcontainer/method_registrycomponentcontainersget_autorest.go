package componentcontainer

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryComponentContainersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ComponentContainerResource
}

// RegistryComponentContainersGet ...
func (c ComponentContainerClient) RegistryComponentContainersGet(ctx context.Context, id RegistryComponentId) (result RegistryComponentContainersGetOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentContainersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryComponentContainersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryComponentContainersGet prepares the RegistryComponentContainersGet request.
func (c ComponentContainerClient) preparerForRegistryComponentContainersGet(ctx context.Context, id RegistryComponentId) (*http.Request, error) {
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

// responderForRegistryComponentContainersGet handles the response to the RegistryComponentContainersGet request. The method always
// closes the http.Response Body.
func (c ComponentContainerClient) responderForRegistryComponentContainersGet(resp *http.Response) (result RegistryComponentContainersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
