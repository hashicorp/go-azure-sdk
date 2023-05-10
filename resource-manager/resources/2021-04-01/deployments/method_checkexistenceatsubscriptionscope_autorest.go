package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistenceAtSubscriptionScope ...
func (c DeploymentsClient) CheckExistenceAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (result CheckExistenceAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForCheckExistenceAtSubscriptionScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistenceAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CheckExistenceAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistenceAtSubscriptionScope prepares the CheckExistenceAtSubscriptionScope request.
func (c DeploymentsClient) preparerForCheckExistenceAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (*http.Request, error) {
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

// responderForCheckExistenceAtSubscriptionScope handles the response to the CheckExistenceAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCheckExistenceAtSubscriptionScope(resp *http.Response) (result CheckExistenceAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
