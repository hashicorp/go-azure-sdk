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

type ExportTemplateAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExportResult
}

// ExportTemplateAtTenantScope ...
func (c DeploymentsClient) ExportTemplateAtTenantScope(ctx context.Context, id DeploymentId) (result ExportTemplateAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForExportTemplateAtTenantScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtTenantScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExportTemplateAtTenantScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtTenantScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExportTemplateAtTenantScope prepares the ExportTemplateAtTenantScope request.
func (c DeploymentsClient) preparerForExportTemplateAtTenantScope(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// responderForExportTemplateAtTenantScope handles the response to the ExportTemplateAtTenantScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForExportTemplateAtTenantScope(resp *http.Response) (result ExportTemplateAtTenantScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
