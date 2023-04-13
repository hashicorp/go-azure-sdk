package deploymentoperations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAtScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentOperation
}

// GetAtScope ...
func (c DeploymentOperationsClient) GetAtScope(ctx context.Context, id ScopedOperationId) (result GetAtScopeOperationResponse, err error) {
	req, err := c.preparerForGetAtScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAtScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAtScope prepares the GetAtScope request.
func (c DeploymentOperationsClient) preparerForGetAtScope(ctx context.Context, id ScopedOperationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAtScope handles the response to the GetAtScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForGetAtScope(resp *http.Response) (result GetAtScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
