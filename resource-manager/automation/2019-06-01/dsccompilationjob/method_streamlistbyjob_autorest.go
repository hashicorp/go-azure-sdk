package dsccompilationjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamListByJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *JobStreamListResult
}

// StreamListByJob ...
func (c DscCompilationJobClient) StreamListByJob(ctx context.Context, id CompilationJobId) (result StreamListByJobOperationResponse, err error) {
	req, err := c.preparerForStreamListByJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "StreamListByJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "StreamListByJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStreamListByJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dsccompilationjob.DscCompilationJobClient", "StreamListByJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStreamListByJob prepares the StreamListByJob request.
func (c DscCompilationJobClient) preparerForStreamListByJob(ctx context.Context, id CompilationJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/streams", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStreamListByJob handles the response to the StreamListByJob request. The method always
// closes the http.Response Body.
func (c DscCompilationJobClient) responderForStreamListByJob(resp *http.Response) (result StreamListByJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
