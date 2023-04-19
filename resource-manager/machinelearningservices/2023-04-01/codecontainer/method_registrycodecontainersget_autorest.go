package codecontainer

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryCodeContainersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *CodeContainerResource
}

// RegistryCodeContainersGet ...
func (c CodeContainerClient) RegistryCodeContainersGet(ctx context.Context, id RegistryCodeId) (result RegistryCodeContainersGetOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeContainersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryCodeContainersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryCodeContainersGet prepares the RegistryCodeContainersGet request.
func (c CodeContainerClient) preparerForRegistryCodeContainersGet(ctx context.Context, id RegistryCodeId) (*http.Request, error) {
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

// responderForRegistryCodeContainersGet handles the response to the RegistryCodeContainersGet request. The method always
// closes the http.Response Body.
func (c CodeContainerClient) responderForRegistryCodeContainersGet(resp *http.Response) (result RegistryCodeContainersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
