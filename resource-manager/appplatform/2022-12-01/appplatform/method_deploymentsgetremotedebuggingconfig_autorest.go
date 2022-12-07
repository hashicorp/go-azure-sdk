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

type DeploymentsGetRemoteDebuggingConfigOperationResponse struct {
	HttpResponse *http.Response
	Model        *RemoteDebugging
}

// DeploymentsGetRemoteDebuggingConfig ...
func (c AppPlatformClient) DeploymentsGetRemoteDebuggingConfig(ctx context.Context, id DeploymentId) (result DeploymentsGetRemoteDebuggingConfigOperationResponse, err error) {
	req, err := c.preparerForDeploymentsGetRemoteDebuggingConfig(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetRemoteDebuggingConfig", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetRemoteDebuggingConfig", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeploymentsGetRemoteDebuggingConfig(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetRemoteDebuggingConfig", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeploymentsGetRemoteDebuggingConfig prepares the DeploymentsGetRemoteDebuggingConfig request.
func (c AppPlatformClient) preparerForDeploymentsGetRemoteDebuggingConfig(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getRemoteDebuggingConfig", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeploymentsGetRemoteDebuggingConfig handles the response to the DeploymentsGetRemoteDebuggingConfig request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForDeploymentsGetRemoteDebuggingConfig(resp *http.Response) (result DeploymentsGetRemoteDebuggingConfigOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
