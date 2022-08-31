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

type StopOperationResponse struct {
	HttpResponse *http.Response
}

type StopOperationOptions struct {
	ClientRequestId *string
}

func DefaultStopOperationOptions() StopOperationOptions {
	return StopOperationOptions{}
}

func (o StopOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o StopOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Stop ...
func (c JobClient) Stop(ctx context.Context, id JobId, options StopOperationOptions) (result StopOperationResponse, err error) {
	req, err := c.preparerForStop(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Stop", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Stop", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStop(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Stop", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStop prepares the Stop request.
func (c JobClient) preparerForStop(ctx context.Context, id JobId, options StopOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStop handles the response to the Stop request. The method always
// closes the http.Response Body.
func (c JobClient) responderForStop(resp *http.Response) (result StopOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
