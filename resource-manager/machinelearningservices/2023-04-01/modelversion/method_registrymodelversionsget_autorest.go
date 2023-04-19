package modelversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryModelVersionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ModelVersionResource
}

// RegistryModelVersionsGet ...
func (c ModelVersionClient) RegistryModelVersionsGet(ctx context.Context, id RegistryModelVersionId) (result RegistryModelVersionsGetOperationResponse, err error) {
	req, err := c.preparerForRegistryModelVersionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryModelVersionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryModelVersionsGet prepares the RegistryModelVersionsGet request.
func (c ModelVersionClient) preparerForRegistryModelVersionsGet(ctx context.Context, id RegistryModelVersionId) (*http.Request, error) {
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

// responderForRegistryModelVersionsGet handles the response to the RegistryModelVersionsGet request. The method always
// closes the http.Response Body.
func (c ModelVersionClient) responderForRegistryModelVersionsGet(resp *http.Response) (result RegistryModelVersionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
