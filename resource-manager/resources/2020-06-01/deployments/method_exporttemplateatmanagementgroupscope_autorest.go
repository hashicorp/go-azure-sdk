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

type ExportTemplateAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExportResult
}

// ExportTemplateAtManagementGroupScope ...
func (c DeploymentsClient) ExportTemplateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (result ExportTemplateAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForExportTemplateAtManagementGroupScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExportTemplateAtManagementGroupScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtManagementGroupScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExportTemplateAtManagementGroupScope prepares the ExportTemplateAtManagementGroupScope request.
func (c DeploymentsClient) preparerForExportTemplateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (*http.Request, error) {
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

// responderForExportTemplateAtManagementGroupScope handles the response to the ExportTemplateAtManagementGroupScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForExportTemplateAtManagementGroupScope(resp *http.Response) (result ExportTemplateAtManagementGroupScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
