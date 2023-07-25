package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecuteSiteAnalysisOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticAnalysis
}

type ExecuteSiteAnalysisOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultExecuteSiteAnalysisOperationOptions() ExecuteSiteAnalysisOperationOptions {
	return ExecuteSiteAnalysisOperationOptions{}
}

func (o ExecuteSiteAnalysisOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExecuteSiteAnalysisOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndTime != nil {
		out["endTime"] = *o.EndTime
	}

	if o.StartTime != nil {
		out["startTime"] = *o.StartTime
	}

	if o.TimeGrain != nil {
		out["timeGrain"] = *o.TimeGrain
	}

	return out
}

// ExecuteSiteAnalysis ...
func (c DiagnosticsClient) ExecuteSiteAnalysis(ctx context.Context, id AnalysisId, options ExecuteSiteAnalysisOperationOptions) (result ExecuteSiteAnalysisOperationResponse, err error) {
	req, err := c.preparerForExecuteSiteAnalysis(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysis", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysis", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExecuteSiteAnalysis(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysis", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExecuteSiteAnalysis prepares the ExecuteSiteAnalysis request.
func (c DiagnosticsClient) preparerForExecuteSiteAnalysis(ctx context.Context, id AnalysisId, options ExecuteSiteAnalysisOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/execute", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForExecuteSiteAnalysis handles the response to the ExecuteSiteAnalysis request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForExecuteSiteAnalysis(resp *http.Response) (result ExecuteSiteAnalysisOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
