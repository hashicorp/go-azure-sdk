package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildServiceBuilderListDeploymentsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentList
}

// BuildServiceBuilderListDeployments ...
func (c AppPlatformClient) BuildServiceBuilderListDeployments(ctx context.Context, id BuilderId) (result BuildServiceBuilderListDeploymentsOperationResponse, err error) {
	req, err := c.preparerForBuildServiceBuilderListDeployments(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderListDeployments", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderListDeployments", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceBuilderListDeployments(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderListDeployments", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceBuilderListDeployments prepares the BuildServiceBuilderListDeployments request.
func (c AppPlatformClient) preparerForBuildServiceBuilderListDeployments(ctx context.Context, id BuilderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listUsingDeployments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBuildServiceBuilderListDeployments handles the response to the BuildServiceBuilderListDeployments request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceBuilderListDeployments(resp *http.Response) (result BuildServiceBuilderListDeploymentsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
