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

type AnalyzeCustomHostnameOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomHostnameAnalysisResult
}

type AnalyzeCustomHostnameOperationOptions struct {
	HostName *string
}

func DefaultAnalyzeCustomHostnameOperationOptions() AnalyzeCustomHostnameOperationOptions {
	return AnalyzeCustomHostnameOperationOptions{}
}

func (o AnalyzeCustomHostnameOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AnalyzeCustomHostnameOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.HostName != nil {
		out["hostName"] = *o.HostName
	}

	return out
}

// AnalyzeCustomHostname ...
func (c WebAppsClient) AnalyzeCustomHostname(ctx context.Context, id SiteId, options AnalyzeCustomHostnameOperationOptions) (result AnalyzeCustomHostnameOperationResponse, err error) {
	req, err := c.preparerForAnalyzeCustomHostname(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostname", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostname", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAnalyzeCustomHostname(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AnalyzeCustomHostname", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAnalyzeCustomHostname prepares the AnalyzeCustomHostname request.
func (c WebAppsClient) preparerForAnalyzeCustomHostname(ctx context.Context, id SiteId, options AnalyzeCustomHostnameOperationOptions) (*http.Request, error) {
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

// responderForAnalyzeCustomHostname handles the response to the AnalyzeCustomHostname request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForAnalyzeCustomHostname(resp *http.Response) (result AnalyzeCustomHostnameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
