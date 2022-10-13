package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildServiceGetBuildServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *BuildService
}

// BuildServiceGetBuildService ...
func (c AppPlatformClient) BuildServiceGetBuildService(ctx context.Context, id BuildServiceId) (result BuildServiceGetBuildServiceOperationResponse, err error) {
	req, err := c.preparerForBuildServiceGetBuildService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceGetBuildService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceGetBuildService prepares the BuildServiceGetBuildService request.
func (c AppPlatformClient) preparerForBuildServiceGetBuildService(ctx context.Context, id BuildServiceId) (*http.Request, error) {
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

// responderForBuildServiceGetBuildService handles the response to the BuildServiceGetBuildService request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceGetBuildService(resp *http.Response) (result BuildServiceGetBuildServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
