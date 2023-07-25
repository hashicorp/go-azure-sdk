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

type StartContinuousWebJobOperationResponse struct {
	HttpResponse *http.Response
}

// StartContinuousWebJob ...
func (c WebAppsClient) StartContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (result StartContinuousWebJobOperationResponse, err error) {
	req, err := c.preparerForStartContinuousWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartContinuousWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartContinuousWebJob prepares the StartContinuousWebJob request.
func (c WebAppsClient) preparerForStartContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStartContinuousWebJob handles the response to the StartContinuousWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStartContinuousWebJob(resp *http.Response) (result StartContinuousWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
