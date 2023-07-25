package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDiagnosticCategorySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticCategory
}

// GetSiteDiagnosticCategorySlot ...
func (c DiagnosticsClient) GetSiteDiagnosticCategorySlot(ctx context.Context, id SlotDiagnosticId) (result GetSiteDiagnosticCategorySlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteDiagnosticCategorySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategorySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategorySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDiagnosticCategorySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDiagnosticCategorySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDiagnosticCategorySlot prepares the GetSiteDiagnosticCategorySlot request.
func (c DiagnosticsClient) preparerForGetSiteDiagnosticCategorySlot(ctx context.Context, id SlotDiagnosticId) (*http.Request, error) {
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

// responderForGetSiteDiagnosticCategorySlot handles the response to the GetSiteDiagnosticCategorySlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDiagnosticCategorySlot(resp *http.Response) (result GetSiteDiagnosticCategorySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
