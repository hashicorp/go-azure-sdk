package job

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

type ResumeOperationOptions struct {
	ClientRequestId *string
}

func DefaultResumeOperationOptions() ResumeOperationOptions {
	return ResumeOperationOptions{}
}

func (o ResumeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o ResumeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Resume ...
func (c JobClient) Resume(ctx context.Context, id JobId, options ResumeOperationOptions) (result ResumeOperationResponse, err error) {
	req, err := c.preparerForResume(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Resume", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Resume", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResume(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Resume", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResume prepares the Resume request.
func (c JobClient) preparerForResume(ctx context.Context, id JobId, options ResumeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/resume", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResume handles the response to the Resume request. The method always
// closes the http.Response Body.
func (c JobClient) responderForResume(resp *http.Response) (result ResumeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
