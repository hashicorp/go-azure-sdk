package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDiagnosticCategoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticCategory
}

// GetSiteDiagnosticCategory ...
func (c DiagnosticsClient) GetSiteDiagnosticCategory(ctx context.Context, id DiagnosticId) (result GetSiteDiagnosticCategoryOperationResponse, err error) {
	req, err := c.preparerForGetSiteDiagnosticCategory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDiagnosticCategory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDiagnosticCategory prepares the GetSiteDiagnosticCategory request.
func (c DiagnosticsClient) preparerForGetSiteDiagnosticCategory(ctx context.Context, id DiagnosticId) (*http.Request, error) {
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

// responderForGetSiteDiagnosticCategory handles the response to the GetSiteDiagnosticCategory request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDiagnosticCategory(resp *http.Response) (result GetSiteDiagnosticCategoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
