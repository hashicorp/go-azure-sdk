package taskresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

// TasksUpdate ...
func (c TaskResourceClient) TasksUpdate(ctx context.Context, id TaskId, input ProjectTask) (result TasksUpdateOperationResponse, err error) {
	req, err := c.preparerForTasksUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "taskresource.TaskResourceClient", "TasksUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksUpdate prepares the TasksUpdate request.
func (c TaskResourceClient) preparerForTasksUpdate(ctx context.Context, id TaskId, input ProjectTask) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTasksUpdate handles the response to the TasksUpdate request. The method always
// closes the http.Response Body.
func (c TaskResourceClient) responderForTasksUpdate(resp *http.Response) (result TasksUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
