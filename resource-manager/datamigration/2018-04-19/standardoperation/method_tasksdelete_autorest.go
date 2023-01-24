package standardoperation

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type TasksDeleteOperationOptions struct {
	DeleteRunningTasks *bool
}

func DefaultTasksDeleteOperationOptions() TasksDeleteOperationOptions {
	return TasksDeleteOperationOptions{}
}

func (o TasksDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o TasksDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DeleteRunningTasks != nil {
		out["deleteRunningTasks"] = *o.DeleteRunningTasks
	}

	return out
}

// TasksDelete ...
func (c StandardOperationClient) TasksDelete(ctx context.Context, id TaskId, options TasksDeleteOperationOptions) (result TasksDeleteOperationResponse, err error) {
	req, err := c.preparerForTasksDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "TasksDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "TasksDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "TasksDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksDelete prepares the TasksDelete request.
func (c StandardOperationClient) preparerForTasksDelete(ctx context.Context, id TaskId, options TasksDeleteOperationOptions) (*http.Request, error) {
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

// responderForTasksDelete handles the response to the TasksDelete request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForTasksDelete(resp *http.Response) (result TasksDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
