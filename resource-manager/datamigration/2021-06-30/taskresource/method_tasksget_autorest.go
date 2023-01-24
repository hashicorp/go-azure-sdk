package taskresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

type TasksGetOperationOptions struct {
	Expand *string
}

func DefaultTasksGetOperationOptions() TasksGetOperationOptions {
	return TasksGetOperationOptions{}
}

func (o TasksGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o TasksGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// TasksGet ...
func (c TaskResourceClient) TasksGet(ctx context.Context, id TaskId, options TasksGetOperationOptions) (result TasksGetOperationResponse, err error) {
	req, err := c.preparerForTasksGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksGet prepares the TasksGet request.
func (c TaskResourceClient) preparerForTasksGet(ctx context.Context, id TaskId, options TasksGetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTasksGet handles the response to the TasksGet request. The method always
// closes the http.Response Body.
func (c TaskResourceClient) responderForTasksGet(resp *http.Response) (result TasksGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
