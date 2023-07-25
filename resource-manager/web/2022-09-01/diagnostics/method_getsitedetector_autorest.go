package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDetectorOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorDefinitionResource
}

// GetSiteDetector ...
func (c DiagnosticsClient) GetSiteDetector(ctx context.Context, id DiagnosticDetectorId) (result GetSiteDetectorOperationResponse, err error) {
	req, err := c.preparerForGetSiteDetector(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetector", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetector", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDetector(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetector", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDetector prepares the GetSiteDetector request.
func (c DiagnosticsClient) preparerForGetSiteDetector(ctx context.Context, id DiagnosticDetectorId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSiteDetector handles the response to the GetSiteDetector request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDetector(resp *http.Response) (result GetSiteDetectorOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
