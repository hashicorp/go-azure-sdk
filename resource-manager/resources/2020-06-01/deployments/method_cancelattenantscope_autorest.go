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

type CancelAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
}

// CancelAtTenantScope ...
func (c DeploymentsClient) CancelAtTenantScope(ctx context.Context, id DeploymentId) (result CancelAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForCancelAtTenantScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtTenantScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCancelAtTenantScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CancelAtTenantScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCancelAtTenantScope prepares the CancelAtTenantScope request.
func (c DeploymentsClient) preparerForCancelAtTenantScope(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// responderForCancelAtTenantScope handles the response to the CancelAtTenantScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCancelAtTenantScope(resp *http.Response) (result CancelAtTenantScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
