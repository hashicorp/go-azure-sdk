package dsccompilationjob

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStreamOperationResponse struct {
	HttpResponse *http.Response
	Model        *JobStream
}

// GetStream ...
func (c DscCompilationJobClient) GetStream(ctx context.Context, id CompilationJobStreamId) (result GetStreamOperationResponse, err error) {
	req, err := c.preparerForGetStream(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "GetStream", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "GetStream", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetStream(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "GetStream", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetStream prepares the GetStream request.
func (c DscCompilationJobClient) preparerForGetStream(ctx context.Context, id CompilationJobStreamId) (*http.Request, error) {
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

// responderForGetStream handles the response to the GetStream request. The method always
// closes the http.Response Body.
func (c DscCompilationJobClient) responderForGetStream(resp *http.Response) (result GetStreamOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
