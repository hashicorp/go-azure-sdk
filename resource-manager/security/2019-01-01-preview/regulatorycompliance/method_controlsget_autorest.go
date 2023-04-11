package regulatorycompliance

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegulatoryComplianceControl
}

// ControlsGet ...
func (c RegulatoryComplianceClient) ControlsGet(ctx context.Context, id RegulatoryComplianceControlId) (result ControlsGetOperationResponse, err error) {
	req, err := c.preparerForControlsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForControlsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForControlsGet prepares the ControlsGet request.
func (c RegulatoryComplianceClient) preparerForControlsGet(ctx context.Context, id RegulatoryComplianceControlId) (*http.Request, error) {
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

// responderForControlsGet handles the response to the ControlsGet request. The method always
// closes the http.Response Body.
func (c RegulatoryComplianceClient) responderForControlsGet(resp *http.Response) (result ControlsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
