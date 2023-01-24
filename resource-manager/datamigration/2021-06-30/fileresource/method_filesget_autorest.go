package fileresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectFile
}

// FilesGet ...
func (c FileResourceClient) FilesGet(ctx context.Context, id FileId) (result FilesGetOperationResponse, err error) {
	req, err := c.preparerForFilesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesGet prepares the FilesGet request.
func (c FileResourceClient) preparerForFilesGet(ctx context.Context, id FileId) (*http.Request, error) {
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

// responderForFilesGet handles the response to the FilesGet request. The method always
// closes the http.Response Body.
func (c FileResourceClient) responderForFilesGet(resp *http.Response) (result FilesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
