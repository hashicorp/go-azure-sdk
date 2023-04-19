package componentversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryComponentVersionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ComponentVersionResource
}

// RegistryComponentVersionsGet ...
func (c ComponentVersionClient) RegistryComponentVersionsGet(ctx context.Context, id RegistryComponentVersionId) (result RegistryComponentVersionsGetOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentVersionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryComponentVersionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryComponentVersionsGet prepares the RegistryComponentVersionsGet request.
func (c ComponentVersionClient) preparerForRegistryComponentVersionsGet(ctx context.Context, id RegistryComponentVersionId) (*http.Request, error) {
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

// responderForRegistryComponentVersionsGet handles the response to the RegistryComponentVersionsGet request. The method always
// closes the http.Response Body.
func (c ComponentVersionClient) responderForRegistryComponentVersionsGet(resp *http.Response) (result RegistryComponentVersionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
