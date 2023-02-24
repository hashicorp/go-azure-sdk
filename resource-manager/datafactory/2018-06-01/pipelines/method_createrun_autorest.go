package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateRunOperationResponse struct {
	HttpResponse *http.Response
	Model        *CreateRunResponse
}

type CreateRunOperationOptions struct {
	IsRecovery             *bool
	ReferencePipelineRunId *string
	StartActivityName      *string
	StartFromFailure       *bool
}

func DefaultCreateRunOperationOptions() CreateRunOperationOptions {
	return CreateRunOperationOptions{}
}

func (o CreateRunOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o CreateRunOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IsRecovery != nil {
		out["isRecovery"] = *o.IsRecovery
	}

	if o.ReferencePipelineRunId != nil {
		out["referencePipelineRunId"] = *o.ReferencePipelineRunId
	}

	if o.StartActivityName != nil {
		out["startActivityName"] = *o.StartActivityName
	}

	if o.StartFromFailure != nil {
		out["startFromFailure"] = *o.StartFromFailure
	}

	return out
}

// CreateRun ...
func (c PipelinesClient) CreateRun(ctx context.Context, id PipelineId, input map[string]interface{}, options CreateRunOperationOptions) (result CreateRunOperationResponse, err error) {
	req, err := c.preparerForCreateRun(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pipelines.PipelinesClient", "CreateRun", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pipelines.PipelinesClient", "CreateRun", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateRun(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pipelines.PipelinesClient", "CreateRun", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateRun prepares the CreateRun request.
func (c PipelinesClient) preparerForCreateRun(ctx context.Context, id PipelineId, input map[string]interface{}, options CreateRunOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/createRun", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateRun handles the response to the CreateRun request. The method always
// closes the http.Response Body.
func (c PipelinesClient) responderForCreateRun(resp *http.Response) (result CreateRunOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
