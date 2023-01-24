package patch

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectFile
}

// FilesUpdate ...
func (c PATCHClient) FilesUpdate(ctx context.Context, id FileId, input ProjectFile) (result FilesUpdateOperationResponse, err error) {
	req, err := c.preparerForFilesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "patch.PATCHClient", "FilesUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "patch.PATCHClient", "FilesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "patch.PATCHClient", "FilesUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesUpdate prepares the FilesUpdate request.
func (c PATCHClient) preparerForFilesUpdate(ctx context.Context, id FileId, input ProjectFile) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFilesUpdate handles the response to the FilesUpdate request. The method always
// closes the http.Response Body.
func (c PATCHClient) responderForFilesUpdate(resp *http.Response) (result FilesUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
