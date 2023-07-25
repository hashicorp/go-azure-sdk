package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetContinuousWebJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContinuousWebJob
}

// GetContinuousWebJob ...
func (c WebAppsClient) GetContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (result GetContinuousWebJobOperationResponse, err error) {
	req, err := c.preparerForGetContinuousWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetContinuousWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetContinuousWebJob prepares the GetContinuousWebJob request.
func (c WebAppsClient) preparerForGetContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetContinuousWebJob handles the response to the GetContinuousWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetContinuousWebJob(resp *http.Response) (result GetContinuousWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
