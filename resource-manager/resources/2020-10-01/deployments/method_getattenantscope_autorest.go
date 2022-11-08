package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentExtended
}

// GetAtTenantScope ...
func (c DeploymentsClient) GetAtTenantScope(ctx context.Context, id DeploymentId) (result GetAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForGetAtTenantScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "GetAtTenantScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "GetAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAtTenantScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "GetAtTenantScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAtTenantScope prepares the GetAtTenantScope request.
func (c DeploymentsClient) preparerForGetAtTenantScope(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// responderForGetAtTenantScope handles the response to the GetAtTenantScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForGetAtTenantScope(resp *http.Response) (result GetAtTenantScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
