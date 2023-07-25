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

type DeployWorkflowArtifactsSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeployWorkflowArtifactsSlot ...
func (c WebAppsClient) DeployWorkflowArtifactsSlot(ctx context.Context, id SlotId, input WorkflowArtifacts) (result DeployWorkflowArtifactsSlotOperationResponse, err error) {
	req, err := c.preparerForDeployWorkflowArtifactsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifactsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifactsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeployWorkflowArtifactsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeployWorkflowArtifactsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeployWorkflowArtifactsSlot prepares the DeployWorkflowArtifactsSlot request.
func (c WebAppsClient) preparerForDeployWorkflowArtifactsSlot(ctx context.Context, id SlotId, input WorkflowArtifacts) (*http.Request, error) {
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

// responderForDeployWorkflowArtifactsSlot handles the response to the DeployWorkflowArtifactsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeployWorkflowArtifactsSlot(resp *http.Response) (result DeployWorkflowArtifactsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
