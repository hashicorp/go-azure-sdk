package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDetectorResponseOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorResponse
}

type GetSiteDetectorResponseOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultGetSiteDetectorResponseOperationOptions() GetSiteDetectorResponseOperationOptions {
	return GetSiteDetectorResponseOperationOptions{}
}

func (o GetSiteDetectorResponseOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetSiteDetectorResponseOperationOptions) toQueryString() map[string]interface{} {
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

// GetSiteDetectorResponse ...
func (c DiagnosticsClient) GetSiteDetectorResponse(ctx context.Context, id DetectorId, options GetSiteDetectorResponseOperationOptions) (result GetSiteDetectorResponseOperationResponse, err error) {
	req, err := c.preparerForGetSiteDetectorResponse(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponse", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponse", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDetectorResponse(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorResponse", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDetectorResponse prepares the GetSiteDetectorResponse request.
func (c DiagnosticsClient) preparerForGetSiteDetectorResponse(ctx context.Context, id DetectorId, options GetSiteDetectorResponseOperationOptions) (*http.Request, error) {
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

// responderForGetSiteDetectorResponse handles the response to the GetSiteDetectorResponse request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDetectorResponse(resp *http.Response) (result GetSiteDetectorResponseOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
