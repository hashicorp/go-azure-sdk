package extensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionMetadataListOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExtensionValueListResult
}

// ExtensionMetadataList ...
func (c ExtensionsClient) ExtensionMetadataList(ctx context.Context, id ExtensionTypeId) (result ExtensionMetadataListOperationResponse, err error) {
	req, err := c.preparerForExtensionMetadataList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExtensionMetadataList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "extensions.ExtensionsClient", "ExtensionMetadataList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExtensionMetadataList prepares the ExtensionMetadataList request.
func (c ExtensionsClient) preparerForExtensionMetadataList(ctx context.Context, id ExtensionTypeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForExtensionMetadataList handles the response to the ExtensionMetadataList request. The method always
// closes the http.Response Body.
func (c ExtensionsClient) responderForExtensionMetadataList(resp *http.Response) (result ExtensionMetadataListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
