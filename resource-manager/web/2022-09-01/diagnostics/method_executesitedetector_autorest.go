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

type ExecuteSiteDetectorOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticDetectorResponse
}

type ExecuteSiteDetectorOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultExecuteSiteDetectorOperationOptions() ExecuteSiteDetectorOperationOptions {
	return ExecuteSiteDetectorOperationOptions{}
}

func (o ExecuteSiteDetectorOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExecuteSiteDetectorOperationOptions) toQueryString() map[string]interface{} {
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

// ExecuteSiteDetector ...
func (c DiagnosticsClient) ExecuteSiteDetector(ctx context.Context, id DiagnosticDetectorId, options ExecuteSiteDetectorOperationOptions) (result ExecuteSiteDetectorOperationResponse, err error) {
	req, err := c.preparerForExecuteSiteDetector(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetector", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetector", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExecuteSiteDetector(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetector", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExecuteSiteDetector prepares the ExecuteSiteDetector request.
func (c DiagnosticsClient) preparerForExecuteSiteDetector(ctx context.Context, id DiagnosticDetectorId, options ExecuteSiteDetectorOperationOptions) (*http.Request, error) {
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

// responderForExecuteSiteDetector handles the response to the ExecuteSiteDetector request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForExecuteSiteDetector(resp *http.Response) (result ExecuteSiteDetectorOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
