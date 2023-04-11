package regulatorycompliance

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegulatoryComplianceAssessment
}

// AssessmentsGet ...
func (c RegulatoryComplianceClient) AssessmentsGet(ctx context.Context, id RegulatoryComplianceAssessmentId) (result AssessmentsGetOperationResponse, err error) {
	req, err := c.preparerForAssessmentsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAssessmentsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAssessmentsGet prepares the AssessmentsGet request.
func (c RegulatoryComplianceClient) preparerForAssessmentsGet(ctx context.Context, id RegulatoryComplianceAssessmentId) (*http.Request, error) {
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

// responderForAssessmentsGet handles the response to the AssessmentsGet request. The method always
// closes the http.Response Body.
func (c RegulatoryComplianceClient) responderForAssessmentsGet(resp *http.Response) (result AssessmentsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
