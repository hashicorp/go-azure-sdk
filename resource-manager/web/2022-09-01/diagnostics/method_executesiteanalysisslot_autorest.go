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

type ExecuteSiteAnalysisSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticAnalysis
}

type ExecuteSiteAnalysisSlotOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultExecuteSiteAnalysisSlotOperationOptions() ExecuteSiteAnalysisSlotOperationOptions {
	return ExecuteSiteAnalysisSlotOperationOptions{}
}

func (o ExecuteSiteAnalysisSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExecuteSiteAnalysisSlotOperationOptions) toQueryString() map[string]interface{} {
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

// ExecuteSiteAnalysisSlot ...
func (c DiagnosticsClient) ExecuteSiteAnalysisSlot(ctx context.Context, id DiagnosticAnalysisId, options ExecuteSiteAnalysisSlotOperationOptions) (result ExecuteSiteAnalysisSlotOperationResponse, err error) {
	req, err := c.preparerForExecuteSiteAnalysisSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysisSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysisSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExecuteSiteAnalysisSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteAnalysisSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExecuteSiteAnalysisSlot prepares the ExecuteSiteAnalysisSlot request.
func (c DiagnosticsClient) preparerForExecuteSiteAnalysisSlot(ctx context.Context, id DiagnosticAnalysisId, options ExecuteSiteAnalysisSlotOperationOptions) (*http.Request, error) {
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

// responderForExecuteSiteAnalysisSlot handles the response to the ExecuteSiteAnalysisSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForExecuteSiteAnalysisSlot(resp *http.Response) (result ExecuteSiteAnalysisSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
