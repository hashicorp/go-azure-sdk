package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeployWorkflowArtifactsOperationResponse struct {
	HttpResponse *http.Response
}

// DeployWorkflowArtifacts ...
func (c WebAppsClient) DeployWorkflowArtifacts(ctx context.Context, id SiteId, input WorkflowArtifacts) (result DeployWorkflowArtifactsOperationResponse, err error) {
	req, err := c.preparerForDeployWorkflowArtifacts(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifacts", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifacts", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeployWorkflowArtifacts(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifacts", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeployWorkflowArtifacts prepares the DeployWorkflowArtifacts request.
func (c WebAppsClient) preparerForDeployWorkflowArtifacts(ctx context.Context, id SiteId, input WorkflowArtifacts) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deployWorkflowArtifacts", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeployWorkflowArtifacts handles the response to the DeployWorkflowArtifacts request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeployWorkflowArtifacts(resp *http.Response) (result DeployWorkflowArtifactsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
