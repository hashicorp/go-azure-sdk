package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDiagnosticsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HostingEnvironmentDiagnostics
}

// ListDiagnostics ...
func (c AppServiceEnvironmentsClient) ListDiagnostics(ctx context.Context, id HostingEnvironmentId) (result ListDiagnosticsOperationResponse, err error) {
	req, err := c.preparerForListDiagnostics(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListDiagnostics", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListDiagnostics", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListDiagnostics(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListDiagnostics", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListDiagnostics prepares the ListDiagnostics request.
func (c AppServiceEnvironmentsClient) preparerForListDiagnostics(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnostics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListDiagnostics handles the response to the ListDiagnostics request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListDiagnostics(resp *http.Response) (result ListDiagnosticsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
