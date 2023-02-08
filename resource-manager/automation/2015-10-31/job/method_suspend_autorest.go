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

// Suspend ...
func (c JobClient) Suspend(ctx context.Context, id JobId) (result SuspendOperationResponse, err error) {
	req, err := c.preparerForSuspend(ctx, id)
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
func (c JobClient) preparerForSuspend(ctx context.Context, id JobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
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
