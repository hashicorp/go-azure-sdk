package libraries

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibraryGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *LibraryResource
}

// LibraryGet ...
func (c LibrariesClient) LibraryGet(ctx context.Context, id LibraryId) (result LibraryGetOperationResponse, err error) {
	req, err := c.preparerForLibraryGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "libraries.LibrariesClient", "LibraryGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "libraries.LibrariesClient", "LibraryGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForLibraryGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "libraries.LibrariesClient", "LibraryGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForLibraryGet prepares the LibraryGet request.
func (c LibrariesClient) preparerForLibraryGet(ctx context.Context, id LibraryId) (*http.Request, error) {
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

// responderForLibraryGet handles the response to the LibraryGet request. The method always
// closes the http.Response Body.
func (c LibrariesClient) responderForLibraryGet(resp *http.Response) (result LibraryGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
