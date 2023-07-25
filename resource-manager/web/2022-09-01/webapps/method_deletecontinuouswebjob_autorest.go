package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteContinuousWebJobOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteContinuousWebJob ...
func (c WebAppsClient) DeleteContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (result DeleteContinuousWebJobOperationResponse, err error) {
	req, err := c.preparerForDeleteContinuousWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteContinuousWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteContinuousWebJob prepares the DeleteContinuousWebJob request.
func (c WebAppsClient) preparerForDeleteContinuousWebJob(ctx context.Context, id ContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteContinuousWebJob handles the response to the DeleteContinuousWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteContinuousWebJob(resp *http.Response) (result DeleteContinuousWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
