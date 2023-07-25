package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeploymentLogOperationResponse struct {
	HttpResponse *http.Response
	Model        *Deployment
}

// ListDeploymentLog ...
func (c WebAppsClient) ListDeploymentLog(ctx context.Context, id DeploymentId) (result ListDeploymentLogOperationResponse, err error) {
	req, err := c.preparerForListDeploymentLog(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLog", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLog", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListDeploymentLog(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLog", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListDeploymentLog prepares the ListDeploymentLog request.
func (c WebAppsClient) preparerForListDeploymentLog(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/log", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListDeploymentLog handles the response to the ListDeploymentLog request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListDeploymentLog(resp *http.Response) (result ListDeploymentLogOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
