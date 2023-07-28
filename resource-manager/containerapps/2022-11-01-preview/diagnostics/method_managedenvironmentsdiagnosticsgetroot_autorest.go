package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEnvironmentsDiagnosticsGetRootOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedEnvironment
}

// ManagedEnvironmentsDiagnosticsGetRoot ...
func (c DiagnosticsClient) ManagedEnvironmentsDiagnosticsGetRoot(ctx context.Context, id ManagedEnvironmentId) (result ManagedEnvironmentsDiagnosticsGetRootOperationResponse, err error) {
	req, err := c.preparerForManagedEnvironmentsDiagnosticsGetRoot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ManagedEnvironmentsDiagnosticsGetRoot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ManagedEnvironmentsDiagnosticsGetRoot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForManagedEnvironmentsDiagnosticsGetRoot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ManagedEnvironmentsDiagnosticsGetRoot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForManagedEnvironmentsDiagnosticsGetRoot prepares the ManagedEnvironmentsDiagnosticsGetRoot request.
func (c DiagnosticsClient) preparerForManagedEnvironmentsDiagnosticsGetRoot(ctx context.Context, id ManagedEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectorProperties/rootApi", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForManagedEnvironmentsDiagnosticsGetRoot handles the response to the ManagedEnvironmentsDiagnosticsGetRoot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForManagedEnvironmentsDiagnosticsGetRoot(resp *http.Response) (result ManagedEnvironmentsDiagnosticsGetRootOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
