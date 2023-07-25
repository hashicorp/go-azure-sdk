package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RebootWorkerOperationResponse struct {
	HttpResponse *http.Response
}

// RebootWorker ...
func (c AppServicePlansClient) RebootWorker(ctx context.Context, id WorkerId) (result RebootWorkerOperationResponse, err error) {
	req, err := c.preparerForRebootWorker(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RebootWorker", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RebootWorker", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRebootWorker(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "RebootWorker", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRebootWorker prepares the RebootWorker request.
func (c AppServicePlansClient) preparerForRebootWorker(ctx context.Context, id WorkerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reboot", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRebootWorker handles the response to the RebootWorker request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForRebootWorker(resp *http.Response) (result RebootWorkerOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
