package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDetectorResponseSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorResponse
}

type GetSiteDetectorResponseSlotOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultGetSiteDetectorResponseSlotOperationOptions() GetSiteDetectorResponseSlotOperationOptions {
	return GetSiteDetectorResponseSlotOperationOptions{}
}

func (o GetSiteDetectorResponseSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetSiteDetectorResponseSlotOperationOptions) toQueryString() map[string]interface{} {
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

// GetSiteDetectorResponseSlot ...
func (c DiagnosticsClient) GetSiteDetectorResponseSlot(ctx context.Context, id SlotDetectorId, options GetSiteDetectorResponseSlotOperationOptions) (result GetSiteDetectorResponseSlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteDetectorResponseSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponseSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponseSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDetectorResponseSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponseSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDetectorResponseSlot prepares the GetSiteDetectorResponseSlot request.
func (c DiagnosticsClient) preparerForGetSiteDetectorResponseSlot(ctx context.Context, id SlotDetectorId, options GetSiteDetectorResponseSlotOperationOptions) (*http.Request, error) {
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

// responderForGetSiteDetectorResponseSlot handles the response to the GetSiteDetectorResponseSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDetectorResponseSlot(resp *http.Response) (result GetSiteDetectorResponseSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
