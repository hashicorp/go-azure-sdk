package certificateordersdiagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppServiceCertificateOrderDetectorResponseOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorResponse
}

type GetAppServiceCertificateOrderDetectorResponseOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultGetAppServiceCertificateOrderDetectorResponseOperationOptions() GetAppServiceCertificateOrderDetectorResponseOperationOptions {
	return GetAppServiceCertificateOrderDetectorResponseOperationOptions{}
}

func (o GetAppServiceCertificateOrderDetectorResponseOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAppServiceCertificateOrderDetectorResponseOperationOptions) toQueryString() map[string]interface{} {
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

// GetAppServiceCertificateOrderDetectorResponse ...
func (c CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponse(ctx context.Context, id DetectorId, options GetAppServiceCertificateOrderDetectorResponseOperationOptions) (result GetAppServiceCertificateOrderDetectorResponseOperationResponse, err error) {
	req, err := c.preparerForGetAppServiceCertificateOrderDetectorResponse(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificateordersdiagnostics.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificateordersdiagnostics.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAppServiceCertificateOrderDetectorResponse(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificateordersdiagnostics.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAppServiceCertificateOrderDetectorResponse prepares the GetAppServiceCertificateOrderDetectorResponse request.
func (c CertificateOrdersDiagnosticsClient) preparerForGetAppServiceCertificateOrderDetectorResponse(ctx context.Context, id DetectorId, options GetAppServiceCertificateOrderDetectorResponseOperationOptions) (*http.Request, error) {
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

// responderForGetAppServiceCertificateOrderDetectorResponse handles the response to the GetAppServiceCertificateOrderDetectorResponse request. The method always
// closes the http.Response Body.
func (c CertificateOrdersDiagnosticsClient) responderForGetAppServiceCertificateOrderDetectorResponse(resp *http.Response) (result GetAppServiceCertificateOrderDetectorResponseOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
