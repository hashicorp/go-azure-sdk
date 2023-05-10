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

type CancelAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CancelAtSubscriptionScope ...
func (c DeploymentsClient) CancelAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (result CancelAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForCancelAtSubscriptionScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCancelAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCancelAtSubscriptionScope prepares the CancelAtSubscriptionScope request.
func (c DeploymentsClient) preparerForCancelAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancel", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCancelAtSubscriptionScope handles the response to the CancelAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCancelAtSubscriptionScope(resp *http.Response) (result CancelAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
