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

type GetRunbookContentOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetRunbookContent ...
func (c JobClient) GetRunbookContent(ctx context.Context, id JobId) (result GetRunbookContentOperationResponse, err error) {
	req, err := c.preparerForGetRunbookContent(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetRunbookContent", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetRunbookContent", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRunbookContent(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetRunbookContent", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRunbookContent prepares the GetRunbookContent request.
func (c JobClient) preparerForGetRunbookContent(ctx context.Context, id JobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/runbookContent", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetRunbookContent handles the response to the GetRunbookContent request. The method always
// closes the http.Response Body.
func (c JobClient) responderForGetRunbookContent(resp *http.Response) (result GetRunbookContentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
