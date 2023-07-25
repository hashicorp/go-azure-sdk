package appserviceenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWorkerPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkerPoolResource
}

// GetWorkerPool ...
func (c AppServiceEnvironmentsClient) GetWorkerPool(ctx context.Context, id WorkerPoolId) (result GetWorkerPoolOperationResponse, err error) {
	req, err := c.preparerForGetWorkerPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetWorkerPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetWorkerPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetWorkerPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetWorkerPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetWorkerPool prepares the GetWorkerPool request.
func (c AppServiceEnvironmentsClient) preparerForGetWorkerPool(ctx context.Context, id WorkerPoolId) (*http.Request, error) {
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

// responderForGetWorkerPool handles the response to the GetWorkerPool request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetWorkerPool(resp *http.Response) (result GetWorkerPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
