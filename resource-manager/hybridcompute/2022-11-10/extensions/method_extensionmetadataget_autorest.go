package extensions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionMetadataGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExtensionValue
}

// ExtensionMetadataGet ...
func (c ExtensionsClient) ExtensionMetadataGet(ctx context.Context, id VersionId) (result ExtensionMetadataGetOperationResponse, err error) {
	req, err := c.preparerForExtensionMetadataGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExtensionMetadataGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExtensionMetadataGet prepares the ExtensionMetadataGet request.
func (c ExtensionsClient) preparerForExtensionMetadataGet(ctx context.Context, id VersionId) (*http.Request, error) {
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

// responderForExtensionMetadataGet handles the response to the ExtensionMetadataGet request. The method always
// closes the http.Response Body.
func (c ExtensionsClient) responderForExtensionMetadataGet(resp *http.Response) (result ExtensionMetadataGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
