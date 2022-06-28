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

type CancelAtScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CancelAtScope ...
func (c DeploymentsClient) CancelAtScope(ctx context.Context, id ScopedDeploymentId) (result CancelAtScopeOperationResponse, err error) {
	req, err := c.preparerForCancelAtScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCancelAtScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCancelAtScope prepares the CancelAtScope request.
func (c DeploymentsClient) preparerForCancelAtScope(ctx context.Context, id ScopedDeploymentId) (*http.Request, error) {
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

// responderForCancelAtScope handles the response to the CancelAtScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCancelAtScope(resp *http.Response) (result CancelAtScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
