package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHostingEnvironmentDetectorResponseOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorResponse
}

type GetHostingEnvironmentDetectorResponseOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultGetHostingEnvironmentDetectorResponseOperationOptions() GetHostingEnvironmentDetectorResponseOperationOptions {
	return GetHostingEnvironmentDetectorResponseOperationOptions{}
}

func (o GetHostingEnvironmentDetectorResponseOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetHostingEnvironmentDetectorResponseOperationOptions) toQueryString() map[string]interface{} {
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

// GetHostingEnvironmentDetectorResponse ...
func (c DiagnosticsClient) GetHostingEnvironmentDetectorResponse(ctx context.Context, id HostingEnvironmentDetectorId, options GetHostingEnvironmentDetectorResponseOperationOptions) (result GetHostingEnvironmentDetectorResponseOperationResponse, err error) {
	req, err := c.preparerForGetHostingEnvironmentDetectorResponse(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetHostingEnvironmentDetectorResponse", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetHostingEnvironmentDetectorResponse", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHostingEnvironmentDetectorResponse(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetHostingEnvironmentDetectorResponse", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHostingEnvironmentDetectorResponse prepares the GetHostingEnvironmentDetectorResponse request.
func (c DiagnosticsClient) preparerForGetHostingEnvironmentDetectorResponse(ctx context.Context, id HostingEnvironmentDetectorId, options GetHostingEnvironmentDetectorResponseOperationOptions) (*http.Request, error) {
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

// responderForGetHostingEnvironmentDetectorResponse handles the response to the GetHostingEnvironmentDetectorResponse request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetHostingEnvironmentDetectorResponse(resp *http.Response) (result GetHostingEnvironmentDetectorResponseOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
