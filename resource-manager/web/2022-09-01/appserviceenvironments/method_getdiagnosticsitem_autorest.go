package appserviceenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDiagnosticsItemOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostingEnvironmentDiagnostics
}

// GetDiagnosticsItem ...
func (c AppServiceEnvironmentsClient) GetDiagnosticsItem(ctx context.Context, id HostingEnvironmentDiagnosticId) (result GetDiagnosticsItemOperationResponse, err error) {
	req, err := c.preparerForGetDiagnosticsItem(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetDiagnosticsItem", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetDiagnosticsItem", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDiagnosticsItem(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetDiagnosticsItem", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDiagnosticsItem prepares the GetDiagnosticsItem request.
func (c AppServiceEnvironmentsClient) preparerForGetDiagnosticsItem(ctx context.Context, id HostingEnvironmentDiagnosticId) (*http.Request, error) {
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

// responderForGetDiagnosticsItem handles the response to the GetDiagnosticsItem request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetDiagnosticsItem(resp *http.Response) (result GetDiagnosticsItemOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
