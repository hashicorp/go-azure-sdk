package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteDetectorSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *DetectorDefinitionResource
}

// GetSiteDetectorSlot ...
func (c DiagnosticsClient) GetSiteDetectorSlot(ctx context.Context, id SlotDiagnosticDetectorId) (result GetSiteDetectorSlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteDetectorSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteDetectorSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "GetSiteDetectorSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteDetectorSlot prepares the GetSiteDetectorSlot request.
func (c DiagnosticsClient) preparerForGetSiteDetectorSlot(ctx context.Context, id SlotDiagnosticDetectorId) (*http.Request, error) {
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

// responderForGetSiteDetectorSlot handles the response to the GetSiteDetectorSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForGetSiteDetectorSlot(resp *http.Response) (result GetSiteDetectorSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
