package diagnostics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsDiagnosticsGetRevisionOperationResponse struct {
	HttpResponse *http.Response
	Model        *Revision
}

// ContainerAppsDiagnosticsGetRevision ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsGetRevision(ctx context.Context, id RevisionsApiRevisionId) (result ContainerAppsDiagnosticsGetRevisionOperationResponse, err error) {
	req, err := c.preparerForContainerAppsDiagnosticsGetRevision(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRevision", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRevision", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContainerAppsDiagnosticsGetRevision(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsGetRevision", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContainerAppsDiagnosticsGetRevision prepares the ContainerAppsDiagnosticsGetRevision request.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsGetRevision(ctx context.Context, id RevisionsApiRevisionId) (*http.Request, error) {
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

// responderForContainerAppsDiagnosticsGetRevision handles the response to the ContainerAppsDiagnosticsGetRevision request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForContainerAppsDiagnosticsGetRevision(resp *http.Response) (result ContainerAppsDiagnosticsGetRevisionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
