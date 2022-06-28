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

type ExportTemplateOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExportResult
}

// ExportTemplate ...
func (c DeploymentsClient) ExportTemplate(ctx context.Context, id ResourceGroupProviderDeploymentId) (result ExportTemplateOperationResponse, err error) {
	req, err := c.preparerForExportTemplate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExportTemplate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExportTemplate prepares the ExportTemplate request.
func (c DeploymentsClient) preparerForExportTemplate(ctx context.Context, id ResourceGroupProviderDeploymentId) (*http.Request, error) {
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

// responderForExportTemplate handles the response to the ExportTemplate request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForExportTemplate(resp *http.Response) (result ExportTemplateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
