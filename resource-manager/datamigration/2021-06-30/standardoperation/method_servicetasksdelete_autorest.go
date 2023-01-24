package standardoperation

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTasksDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type ServiceTasksDeleteOperationOptions struct {
	DeleteRunningTasks *bool
}

func DefaultServiceTasksDeleteOperationOptions() ServiceTasksDeleteOperationOptions {
	return ServiceTasksDeleteOperationOptions{}
}

func (o ServiceTasksDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ServiceTasksDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DeleteRunningTasks != nil {
		out["deleteRunningTasks"] = *o.DeleteRunningTasks
	}

	return out
}

// ServiceTasksDelete ...
func (c StandardOperationClient) ServiceTasksDelete(ctx context.Context, id ServiceTaskId, options ServiceTasksDeleteOperationOptions) (result ServiceTasksDeleteOperationResponse, err error) {
	req, err := c.preparerForServiceTasksDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServiceTasksDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServiceTasksDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServiceTasksDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServiceTasksDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServiceTasksDelete prepares the ServiceTasksDelete request.
func (c StandardOperationClient) preparerForServiceTasksDelete(ctx context.Context, id ServiceTaskId, options ServiceTasksDeleteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServiceTasksDelete handles the response to the ServiceTasksDelete request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForServiceTasksDelete(resp *http.Response) (result ServiceTasksDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
