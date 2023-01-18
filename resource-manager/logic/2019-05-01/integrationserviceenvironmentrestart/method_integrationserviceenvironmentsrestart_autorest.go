package integrationserviceenvironmentrestart

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentsRestartOperationResponse struct {
	HttpResponse *http.Response
}

// IntegrationServiceEnvironmentsRestart ...
func (c IntegrationServiceEnvironmentRestartClient) IntegrationServiceEnvironmentsRestart(ctx context.Context, id IntegrationServiceEnvironmentId) (result IntegrationServiceEnvironmentsRestartOperationResponse, err error) {
	req, err := c.preparerForIntegrationServiceEnvironmentsRestart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentrestart.IntegrationServiceEnvironmentRestartClient", "IntegrationServiceEnvironmentsRestart", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentrestart.IntegrationServiceEnvironmentRestartClient", "IntegrationServiceEnvironmentsRestart", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIntegrationServiceEnvironmentsRestart(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentrestart.IntegrationServiceEnvironmentRestartClient", "IntegrationServiceEnvironmentsRestart", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIntegrationServiceEnvironmentsRestart prepares the IntegrationServiceEnvironmentsRestart request.
func (c IntegrationServiceEnvironmentRestartClient) preparerForIntegrationServiceEnvironmentsRestart(ctx context.Context, id IntegrationServiceEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIntegrationServiceEnvironmentsRestart handles the response to the IntegrationServiceEnvironmentsRestart request. The method always
// closes the http.Response Body.
func (c IntegrationServiceEnvironmentRestartClient) responderForIntegrationServiceEnvironmentsRestart(resp *http.Response) (result IntegrationServiceEnvironmentsRestartOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
