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

type DeploymentsGetLogFileUrlOperationResponse struct {
	HttpResponse *http.Response
	Model        *LogFileUrlResponse
}

// DeploymentsGetLogFileUrl ...
func (c AppPlatformClient) DeploymentsGetLogFileUrl(ctx context.Context, id DeploymentId) (result DeploymentsGetLogFileUrlOperationResponse, err error) {
	req, err := c.preparerForDeploymentsGetLogFileUrl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetLogFileUrl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetLogFileUrl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeploymentsGetLogFileUrl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGetLogFileUrl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeploymentsGetLogFileUrl prepares the DeploymentsGetLogFileUrl request.
func (c AppPlatformClient) preparerForDeploymentsGetLogFileUrl(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getLogFileUrl", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeploymentsGetLogFileUrl handles the response to the DeploymentsGetLogFileUrl request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForDeploymentsGetLogFileUrl(resp *http.Response) (result DeploymentsGetLogFileUrlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
