package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceAtScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistenceAtScope ...
func (c DeploymentsClient) CheckExistenceAtScope(ctx context.Context, id ScopedDeploymentId) (result CheckExistenceAtScopeOperationResponse, err error) {
	req, err := c.preparerForCheckExistenceAtScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistenceAtScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistenceAtScope prepares the CheckExistenceAtScope request.
func (c DeploymentsClient) preparerForCheckExistenceAtScope(ctx context.Context, id ScopedDeploymentId) (*http.Request, error) {
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

// responderForCheckExistenceAtScope handles the response to the CheckExistenceAtScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCheckExistenceAtScope(resp *http.Response) (result CheckExistenceAtScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
