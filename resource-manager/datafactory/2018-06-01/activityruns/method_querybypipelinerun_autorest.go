package activityruns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryByPipelineRunOperationResponse struct {
	HttpResponse *http.Response
	Model        *ActivityRunsQueryResponse
}

// QueryByPipelineRun ...
func (c ActivityrunsClient) QueryByPipelineRun(ctx context.Context, id PipelineRunId, input RunFilterParameters) (result QueryByPipelineRunOperationResponse, err error) {
	req, err := c.preparerForQueryByPipelineRun(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "activityruns.ActivityrunsClient", "QueryByPipelineRun", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "activityruns.ActivityrunsClient", "QueryByPipelineRun", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForQueryByPipelineRun(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "activityruns.ActivityrunsClient", "QueryByPipelineRun", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForQueryByPipelineRun prepares the QueryByPipelineRun request.
func (c ActivityrunsClient) preparerForQueryByPipelineRun(ctx context.Context, id PipelineRunId, input RunFilterParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/queryActivityruns", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForQueryByPipelineRun handles the response to the QueryByPipelineRun request. The method always
// closes the http.Response Body.
func (c ActivityrunsClient) responderForQueryByPipelineRun(resp *http.Response) (result QueryByPipelineRunOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
