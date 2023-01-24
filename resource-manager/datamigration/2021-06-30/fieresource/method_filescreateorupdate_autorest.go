package fieresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectFile
}

// FilesCreateOrUpdate ...
func (c FieResourceClient) FilesCreateOrUpdate(ctx context.Context, id FileId, input ProjectFile) (result FilesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForFilesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fieresource.FieResourceClient", "FilesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "fieresource.FieResourceClient", "FilesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFilesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fieresource.FieResourceClient", "FilesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFilesCreateOrUpdate prepares the FilesCreateOrUpdate request.
func (c FieResourceClient) preparerForFilesCreateOrUpdate(ctx context.Context, id FileId, input ProjectFile) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFilesCreateOrUpdate handles the response to the FilesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c FieResourceClient) responderForFilesCreateOrUpdate(resp *http.Response) (result FilesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
