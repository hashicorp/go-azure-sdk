package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsDiagnosticsGetDetectorOperationResponse struct {
	HttpResponse *http.Response
	Model        *Diagnostics
}

// ContainerAppsDiagnosticsGetDetector ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsGetDetector(ctx context.Context, id DetectorId) (result ContainerAppsDiagnosticsGetDetectorOperationResponse, err error) {
	req, err := c.preparerForContainerAppsDiagnosticsGetDetector(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetDetector", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetDetector", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContainerAppsDiagnosticsGetDetector(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetDetector", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContainerAppsDiagnosticsGetDetector prepares the ContainerAppsDiagnosticsGetDetector request.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsGetDetector(ctx context.Context, id DetectorId) (*http.Request, error) {
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

// responderForContainerAppsDiagnosticsGetDetector handles the response to the ContainerAppsDiagnosticsGetDetector request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForContainerAppsDiagnosticsGetDetector(resp *http.Response) (result ContainerAppsDiagnosticsGetDetectorOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
