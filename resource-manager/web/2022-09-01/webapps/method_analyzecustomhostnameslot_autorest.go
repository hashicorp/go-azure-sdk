package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyzeCustomHostnameSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomHostnameAnalysisResult
}

type AnalyzeCustomHostnameSlotOperationOptions struct {
	HostName *string
}

func DefaultAnalyzeCustomHostnameSlotOperationOptions() AnalyzeCustomHostnameSlotOperationOptions {
	return AnalyzeCustomHostnameSlotOperationOptions{}
}

func (o AnalyzeCustomHostnameSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AnalyzeCustomHostnameSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.HostName != nil {
		out["hostName"] = *o.HostName
	}

	return out
}

// AnalyzeCustomHostnameSlot ...
func (c WebAppsClient) AnalyzeCustomHostnameSlot(ctx context.Context, id SlotId, options AnalyzeCustomHostnameSlotOperationOptions) (result AnalyzeCustomHostnameSlotOperationResponse, err error) {
	req, err := c.preparerForAnalyzeCustomHostnameSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostnameSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostnameSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAnalyzeCustomHostnameSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostnameSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAnalyzeCustomHostnameSlot prepares the AnalyzeCustomHostnameSlot request.
func (c WebAppsClient) preparerForAnalyzeCustomHostnameSlot(ctx context.Context, id SlotId, options AnalyzeCustomHostnameSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/analyzeCustomHostname", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAnalyzeCustomHostnameSlot handles the response to the AnalyzeCustomHostnameSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForAnalyzeCustomHostnameSlot(resp *http.Response) (result AnalyzeCustomHostnameSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
