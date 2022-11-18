package testjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResumeOperationResponse struct {
	HttpResponse *http.Response
}

// Resume ...
func (c TestJobClient) Resume(ctx context.Context, id RunbookId) (result ResumeOperationResponse, err error) {
	req, err := c.preparerForResume(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Resume", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Resume", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResume(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Resume", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResume prepares the Resume request.
func (c TestJobClient) preparerForResume(ctx context.Context, id RunbookId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/draft/testJob/resume", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResume handles the response to the Resume request. The method always
// closes the http.Response Body.
func (c TestJobClient) responderForResume(resp *http.Response) (result ResumeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
