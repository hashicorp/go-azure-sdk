package jobcancellations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerOperationResponse struct {
	HttpResponse *http.Response
}

// Trigger ...
func (c JobCancellationsClient) Trigger(ctx context.Context, id BackupJobId) (result TriggerOperationResponse, err error) {
	req, err := c.preparerForTrigger(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobcancellations.JobCancellationsClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobcancellations.JobCancellationsClient", "Trigger", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTrigger(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobcancellations.JobCancellationsClient", "Trigger", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTrigger prepares the Trigger request.
func (c JobCancellationsClient) preparerForTrigger(ctx context.Context, id BackupJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancel", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTrigger handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (c JobCancellationsClient) responderForTrigger(resp *http.Response) (result TriggerOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
