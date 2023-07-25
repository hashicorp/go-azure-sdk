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

type ExecuteSiteDetectorSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticDetectorResponse
}

type ExecuteSiteDetectorSlotOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultExecuteSiteDetectorSlotOperationOptions() ExecuteSiteDetectorSlotOperationOptions {
	return ExecuteSiteDetectorSlotOperationOptions{}
}

func (o ExecuteSiteDetectorSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExecuteSiteDetectorSlotOperationOptions) toQueryString() map[string]interface{} {
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

// ExecuteSiteDetectorSlot ...
func (c DiagnosticsClient) ExecuteSiteDetectorSlot(ctx context.Context, id SlotDiagnosticDetectorId, options ExecuteSiteDetectorSlotOperationOptions) (result ExecuteSiteDetectorSlotOperationResponse, err error) {
	req, err := c.preparerForExecuteSiteDetectorSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetectorSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetectorSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExecuteSiteDetectorSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ExecuteSiteDetectorSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExecuteSiteDetectorSlot prepares the ExecuteSiteDetectorSlot request.
func (c DiagnosticsClient) preparerForExecuteSiteDetectorSlot(ctx context.Context, id SlotDiagnosticDetectorId, options ExecuteSiteDetectorSlotOperationOptions) (*http.Request, error) {
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

// responderForExecuteSiteDetectorSlot handles the response to the ExecuteSiteDetectorSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForExecuteSiteDetectorSlot(resp *http.Response) (result ExecuteSiteDetectorSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
