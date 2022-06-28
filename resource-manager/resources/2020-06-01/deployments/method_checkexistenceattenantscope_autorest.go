package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistenceAtTenantScope ...
func (c DeploymentsClient) CheckExistenceAtTenantScope(ctx context.Context, id DeploymentId) (result CheckExistenceAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForCheckExistenceAtTenantScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtTenantScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistenceAtTenantScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtTenantScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistenceAtTenantScope prepares the CheckExistenceAtTenantScope request.
func (c DeploymentsClient) preparerForCheckExistenceAtTenantScope(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// responderForCheckExistenceAtTenantScope handles the response to the CheckExistenceAtTenantScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCheckExistenceAtTenantScope(resp *http.Response) (result CheckExistenceAtTenantScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
