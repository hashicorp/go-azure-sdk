package standardoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesReadOperationResponse struct {
	HttpResponse *http.Response
	Model        *FileStorageInfo
}

// FilesRead ...
func (c StandardOperationClient) FilesRead(ctx context.Context, id FileId) (result FilesReadOperationResponse, err error) {
	req, err := c.preparerForFilesRead(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "FilesRead", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "FilesRead", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesRead(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "FilesRead", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesRead prepares the FilesRead request.
func (c StandardOperationClient) preparerForFilesRead(ctx context.Context, id FileId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/read", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFilesRead handles the response to the FilesRead request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForFilesRead(resp *http.Response) (result FilesReadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
