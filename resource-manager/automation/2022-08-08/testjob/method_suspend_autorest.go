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

type SuspendOperationResponse struct {
	HttpResponse *http.Response
}

// Suspend ...
func (c TestJobClient) Suspend(ctx context.Context, id RunbookId) (result SuspendOperationResponse, err error) {
	req, err := c.preparerForSuspend(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Suspend", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Suspend", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSuspend(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjob.TestJobClient", "Suspend", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSuspend prepares the Suspend request.
func (c TestJobClient) preparerForSuspend(ctx context.Context, id RunbookId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/draft/testJob/suspend", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSuspend handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (c TestJobClient) responderForSuspend(resp *http.Response) (result SuspendOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
