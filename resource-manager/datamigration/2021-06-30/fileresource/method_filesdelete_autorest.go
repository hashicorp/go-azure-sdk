package fileresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// FilesDelete ...
func (c FileResourceClient) FilesDelete(ctx context.Context, id FileId) (result FilesDeleteOperationResponse, err error) {
	req, err := c.preparerForFilesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fileresource.FileResourceClient", "FilesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesDelete prepares the FilesDelete request.
func (c FileResourceClient) preparerForFilesDelete(ctx context.Context, id FileId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFilesDelete handles the response to the FilesDelete request. The method always
// closes the http.Response Body.
func (c FileResourceClient) responderForFilesDelete(resp *http.Response) (result FilesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
