package environmentcontainer

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEnvironmentContainersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EnvironmentContainerResource
}

// RegistryEnvironmentContainersGet ...
func (c EnvironmentContainerClient) RegistryEnvironmentContainersGet(ctx context.Context, id RegistryEnvironmentId) (result RegistryEnvironmentContainersGetOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentContainersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryEnvironmentContainersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryEnvironmentContainersGet prepares the RegistryEnvironmentContainersGet request.
func (c EnvironmentContainerClient) preparerForRegistryEnvironmentContainersGet(ctx context.Context, id RegistryEnvironmentId) (*http.Request, error) {
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

// responderForRegistryEnvironmentContainersGet handles the response to the RegistryEnvironmentContainersGet request. The method always
// closes the http.Response Body.
func (c EnvironmentContainerClient) responderForRegistryEnvironmentContainersGet(resp *http.Response) (result RegistryEnvironmentContainersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
