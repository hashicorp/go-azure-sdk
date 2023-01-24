package customoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksCancelOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

// TasksCancel ...
func (c CustomOperationClient) TasksCancel(ctx context.Context, id TaskId) (result TasksCancelOperationResponse, err error) {
	req, err := c.preparerForTasksCancel(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "TasksCancel", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "TasksCancel", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTasksCancel(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "TasksCancel", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTasksCancel prepares the TasksCancel request.
func (c CustomOperationClient) preparerForTasksCancel(ctx context.Context, id TaskId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancel", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTasksCancel handles the response to the TasksCancel request. The method always
// closes the http.Response Body.
func (c CustomOperationClient) responderForTasksCancel(resp *http.Response) (result TasksCancelOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
