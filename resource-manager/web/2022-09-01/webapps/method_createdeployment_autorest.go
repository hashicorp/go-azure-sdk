package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeploymentOperationResponse struct {
	HttpResponse *http.Response
	Model        *Deployment
}

// CreateDeployment ...
func (c WebAppsClient) CreateDeployment(ctx context.Context, id DeploymentId, input Deployment) (result CreateDeploymentOperationResponse, err error) {
	req, err := c.preparerForCreateDeployment(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateDeployment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateDeployment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateDeployment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateDeployment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateDeployment prepares the CreateDeployment request.
func (c WebAppsClient) preparerForCreateDeployment(ctx context.Context, id DeploymentId, input Deployment) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateDeployment handles the response to the CreateDeployment request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateDeployment(resp *http.Response) (result CreateDeploymentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
