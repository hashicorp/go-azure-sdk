package environmentversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEnvironmentVersionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EnvironmentVersionResource
}

// RegistryEnvironmentVersionsGet ...
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsGet(ctx context.Context, id RegistryEnvironmentVersionId) (result RegistryEnvironmentVersionsGetOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentVersionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryEnvironmentVersionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryEnvironmentVersionsGet prepares the RegistryEnvironmentVersionsGet request.
func (c EnvironmentVersionClient) preparerForRegistryEnvironmentVersionsGet(ctx context.Context, id RegistryEnvironmentVersionId) (*http.Request, error) {
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

// responderForRegistryEnvironmentVersionsGet handles the response to the RegistryEnvironmentVersionsGet request. The method always
// closes the http.Response Body.
func (c EnvironmentVersionClient) responderForRegistryEnvironmentVersionsGet(resp *http.Response) (result RegistryEnvironmentVersionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
