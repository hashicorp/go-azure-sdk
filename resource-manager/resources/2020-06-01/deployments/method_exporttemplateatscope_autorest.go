package deployments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportTemplateAtScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExportResult
}

// ExportTemplateAtScope ...
func (c DeploymentsClient) ExportTemplateAtScope(ctx context.Context, id ScopedDeploymentId) (result ExportTemplateAtScopeOperationResponse, err error) {
	req, err := c.preparerForExportTemplateAtScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExportTemplateAtScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExportTemplateAtScope prepares the ExportTemplateAtScope request.
func (c DeploymentsClient) preparerForExportTemplateAtScope(ctx context.Context, id ScopedDeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/exportTemplate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForExportTemplateAtScope handles the response to the ExportTemplateAtScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForExportTemplateAtScope(resp *http.Response) (result ExportTemplateAtScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
