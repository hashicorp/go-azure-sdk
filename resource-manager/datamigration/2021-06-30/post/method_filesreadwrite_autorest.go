package post

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesReadWriteOperationResponse struct {
	HttpResponse *http.Response
	Model        *FileStorageInfo
}

// FilesReadWrite ...
func (c POSTClient) FilesReadWrite(ctx context.Context, id FileId) (result FilesReadWriteOperationResponse, err error) {
	req, err := c.preparerForFilesReadWrite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "FilesReadWrite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "FilesReadWrite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesReadWrite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "FilesReadWrite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesReadWrite prepares the FilesReadWrite request.
func (c POSTClient) preparerForFilesReadWrite(ctx context.Context, id FileId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/readwrite", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFilesReadWrite handles the response to the FilesReadWrite request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForFilesReadWrite(resp *http.Response) (result FilesReadWriteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
