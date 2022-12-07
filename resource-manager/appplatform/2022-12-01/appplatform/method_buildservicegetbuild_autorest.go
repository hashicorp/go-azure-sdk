package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildServiceGetBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *Build
}

// BuildServiceGetBuild ...
func (c AppPlatformClient) BuildServiceGetBuild(ctx context.Context, id BuildId) (result BuildServiceGetBuildOperationResponse, err error) {
	req, err := c.preparerForBuildServiceGetBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuild", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuild", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceGetBuild(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuild", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceGetBuild prepares the BuildServiceGetBuild request.
func (c AppPlatformClient) preparerForBuildServiceGetBuild(ctx context.Context, id BuildId) (*http.Request, error) {
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

// responderForBuildServiceGetBuild handles the response to the BuildServiceGetBuild request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceGetBuild(resp *http.Response) (result BuildServiceGetBuildOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
