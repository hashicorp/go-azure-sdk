package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeploymentOperationResponse struct {
	HttpResponse *http.Response
	Model        *Deployment
}

// GetDeployment ...
func (c WebAppsClient) GetDeployment(ctx context.Context, id DeploymentId) (result GetDeploymentOperationResponse, err error) {
	req, err := c.preparerForGetDeployment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeployment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeployment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeployment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeployment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeployment prepares the GetDeployment request.
func (c WebAppsClient) preparerForGetDeployment(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// responderForGetDeployment handles the response to the GetDeployment request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetDeployment(resp *http.Response) (result GetDeploymentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
