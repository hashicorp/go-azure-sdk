package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildServiceGetSupportedBuildpackOperationResponse struct {
	HttpResponse *http.Response
	Model        *SupportedBuildpackResource
}

// BuildServiceGetSupportedBuildpack ...
func (c AppPlatformClient) BuildServiceGetSupportedBuildpack(ctx context.Context, id SupportedBuildPackId) (result BuildServiceGetSupportedBuildpackOperationResponse, err error) {
	req, err := c.preparerForBuildServiceGetSupportedBuildpack(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetSupportedBuildpack", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetSupportedBuildpack", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceGetSupportedBuildpack(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetSupportedBuildpack", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceGetSupportedBuildpack prepares the BuildServiceGetSupportedBuildpack request.
func (c AppPlatformClient) preparerForBuildServiceGetSupportedBuildpack(ctx context.Context, id SupportedBuildPackId) (*http.Request, error) {
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

// responderForBuildServiceGetSupportedBuildpack handles the response to the BuildServiceGetSupportedBuildpack request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceGetSupportedBuildpack(resp *http.Response) (result BuildServiceGetSupportedBuildpackOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
