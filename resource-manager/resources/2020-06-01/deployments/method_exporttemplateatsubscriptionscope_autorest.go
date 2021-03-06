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

type ExportTemplateAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExportResult
}

// ExportTemplateAtSubscriptionScope ...
func (c DeploymentsClient) ExportTemplateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (result ExportTemplateAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForExportTemplateAtSubscriptionScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExportTemplateAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ExportTemplateAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExportTemplateAtSubscriptionScope prepares the ExportTemplateAtSubscriptionScope request.
func (c DeploymentsClient) preparerForExportTemplateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (*http.Request, error) {
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

// responderForExportTemplateAtSubscriptionScope handles the response to the ExportTemplateAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForExportTemplateAtSubscriptionScope(resp *http.Response) (result ExportTemplateAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
