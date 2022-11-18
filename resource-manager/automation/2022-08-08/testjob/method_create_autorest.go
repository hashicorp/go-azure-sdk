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

type CreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *TestJob
}

// Create ...
func (c TestJobClient) Create(ctx context.Context, id RunbookId, input TestJobCreateParameters) (result CreateOperationResponse, err error) {
	req, err := c.preparerForCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Create", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Create", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Create", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreate prepares the Create request.
func (c TestJobClient) preparerForCreate(ctx context.Context, id RunbookId, input TestJobCreateParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/draft/testJob", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreate handles the response to the Create request. The method always
// closes the http.Response Body.
func (c TestJobClient) responderForCreate(resp *http.Response) (result CreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
