package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistenceAtManagementGroupScope ...
func (c DeploymentsClient) CheckExistenceAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (result CheckExistenceAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForCheckExistenceAtManagementGroupScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistenceAtManagementGroupScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtManagementGroupScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistenceAtManagementGroupScope prepares the CheckExistenceAtManagementGroupScope request.
func (c DeploymentsClient) preparerForCheckExistenceAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckExistenceAtManagementGroupScope handles the response to the CheckExistenceAtManagementGroupScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCheckExistenceAtManagementGroupScope(resp *http.Response) (result CheckExistenceAtManagementGroupScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
