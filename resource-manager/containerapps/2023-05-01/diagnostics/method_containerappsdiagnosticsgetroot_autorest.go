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

type ContainerAppsDiagnosticsGetRootOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContainerApp
}

// ContainerAppsDiagnosticsGetRoot ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsGetRoot(ctx context.Context, id ContainerAppId) (result ContainerAppsDiagnosticsGetRootOperationResponse, err error) {
	req, err := c.preparerForContainerAppsDiagnosticsGetRoot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRoot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRoot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContainerAppsDiagnosticsGetRoot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRoot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContainerAppsDiagnosticsGetRoot prepares the ContainerAppsDiagnosticsGetRoot request.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsGetRoot(ctx context.Context, id ContainerAppId) (*http.Request, error) {
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

// responderForContainerAppsDiagnosticsGetRoot handles the response to the ContainerAppsDiagnosticsGetRoot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForContainerAppsDiagnosticsGetRoot(resp *http.Response) (result ContainerAppsDiagnosticsGetRootOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
