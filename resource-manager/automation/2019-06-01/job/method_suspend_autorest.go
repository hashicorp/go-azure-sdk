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

type SuspendOperationResponse struct {
	HttpResponse *http.Response
}

type SuspendOperationOptions struct {
	ClientRequestId *string
}

func DefaultSuspendOperationOptions() SuspendOperationOptions {
	return SuspendOperationOptions{}
}

func (o SuspendOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o SuspendOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Suspend ...
func (c JobClient) Suspend(ctx context.Context, id JobId, options SuspendOperationOptions) (result SuspendOperationResponse, err error) {
	req, err := c.preparerForSuspend(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Suspend", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Suspend", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSuspend(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "Suspend", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSuspend prepares the Suspend request.
func (c JobClient) preparerForSuspend(ctx context.Context, id JobId, options SuspendOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/suspend", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSuspend handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (c JobClient) responderForSuspend(resp *http.Response) (result SuspendOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
