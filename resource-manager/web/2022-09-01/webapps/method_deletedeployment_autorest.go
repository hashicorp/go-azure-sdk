package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDeploymentOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteDeployment ...
func (c WebAppsClient) DeleteDeployment(ctx context.Context, id DeploymentId) (result DeleteDeploymentOperationResponse, err error) {
	req, err := c.preparerForDeleteDeployment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeployment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeployment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteDeployment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeployment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteDeployment prepares the DeleteDeployment request.
func (c WebAppsClient) preparerForDeleteDeployment(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteDeployment handles the response to the DeleteDeployment request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteDeployment(resp *http.Response) (result DeleteDeploymentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
