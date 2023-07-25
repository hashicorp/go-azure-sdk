package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopContinuousWebJobOperationResponse struct {
	HttpResponse *http.Response
}

// StopContinuousWebJob ...
func (c WebAppsClient) StopContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (result StopContinuousWebJobOperationResponse, err error) {
	req, err := c.preparerForStopContinuousWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopContinuousWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopContinuousWebJob prepares the StopContinuousWebJob request.
func (c WebAppsClient) preparerForStopContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStopContinuousWebJob handles the response to the StopContinuousWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopContinuousWebJob(resp *http.Response) (result StopContinuousWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
