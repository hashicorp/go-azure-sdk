package hcrpreports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByConfigurationProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ReportList
}

// ListByConfigurationProfileAssignments ...
func (c HCRPReportsClient) ListByConfigurationProfileAssignments(ctx context.Context, id ConfigurationProfileAssignmentId) (result ListByConfigurationProfileAssignmentsOperationResponse, err error) {
	req, err := c.preparerForListByConfigurationProfileAssignments(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hcrpreports.HCRPReportsClient", "ListByConfigurationProfileAssignments", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hcrpreports.HCRPReportsClient", "ListByConfigurationProfileAssignments", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByConfigurationProfileAssignments(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hcrpreports.HCRPReportsClient", "ListByConfigurationProfileAssignments", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByConfigurationProfileAssignments prepares the ListByConfigurationProfileAssignments request.
func (c HCRPReportsClient) preparerForListByConfigurationProfileAssignments(ctx context.Context, id ConfigurationProfileAssignmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reports", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByConfigurationProfileAssignments handles the response to the ListByConfigurationProfileAssignments request. The method always
// closes the http.Response Body.
func (c HCRPReportsClient) responderForListByConfigurationProfileAssignments(resp *http.Response) (result ListByConfigurationProfileAssignmentsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
