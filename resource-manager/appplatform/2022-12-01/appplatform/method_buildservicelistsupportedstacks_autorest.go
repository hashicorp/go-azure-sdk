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

type BuildServiceListSupportedStacksOperationResponse struct {
	HttpResponse *http.Response
	Model        *SupportedStacksCollection
}

// BuildServiceListSupportedStacks ...
func (c AppPlatformClient) BuildServiceListSupportedStacks(ctx context.Context, id BuildServiceId) (result BuildServiceListSupportedStacksOperationResponse, err error) {
	req, err := c.preparerForBuildServiceListSupportedStacks(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedStacks", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedStacks", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceListSupportedStacks(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListSupportedStacks", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceListSupportedStacks prepares the BuildServiceListSupportedStacks request.
func (c AppPlatformClient) preparerForBuildServiceListSupportedStacks(ctx context.Context, id BuildServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/supportedStacks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBuildServiceListSupportedStacks handles the response to the BuildServiceListSupportedStacks request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceListSupportedStacks(resp *http.Response) (result BuildServiceListSupportedStacksOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
