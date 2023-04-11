package regulatorycompliance

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegulatoryComplianceStandard
}

// StandardsGet ...
func (c RegulatoryComplianceClient) StandardsGet(ctx context.Context, id RegulatoryComplianceStandardId) (result StandardsGetOperationResponse, err error) {
	req, err := c.preparerForStandardsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "StandardsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "StandardsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStandardsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "StandardsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStandardsGet prepares the StandardsGet request.
func (c RegulatoryComplianceClient) preparerForStandardsGet(ctx context.Context, id RegulatoryComplianceStandardId) (*http.Request, error) {
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

// responderForStandardsGet handles the response to the StandardsGet request. The method always
// closes the http.Response Body.
func (c RegulatoryComplianceClient) responderForStandardsGet(resp *http.Response) (result StandardsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
