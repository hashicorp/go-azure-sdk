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

type BuildServiceListSupportedBuildpacksOperationResponse struct {
	HttpResponse *http.Response
	Model        *SupportedBuildpacksCollection
}

// BuildServiceListSupportedBuildpacks ...
func (c AppPlatformClient) BuildServiceListSupportedBuildpacks(ctx context.Context, id BuildServiceId) (result BuildServiceListSupportedBuildpacksOperationResponse, err error) {
	req, err := c.preparerForBuildServiceListSupportedBuildpacks(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedBuildpacks", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedBuildpacks", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceListSupportedBuildpacks(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedBuildpacks", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceListSupportedBuildpacks prepares the BuildServiceListSupportedBuildpacks request.
func (c AppPlatformClient) preparerForBuildServiceListSupportedBuildpacks(ctx context.Context, id BuildServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/supportedBuildPacks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBuildServiceListSupportedBuildpacks handles the response to the BuildServiceListSupportedBuildpacks request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceListSupportedBuildpacks(resp *http.Response) (result BuildServiceListSupportedBuildpacksOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
