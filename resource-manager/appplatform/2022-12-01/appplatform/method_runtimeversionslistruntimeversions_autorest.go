package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuntimeVersionsListRuntimeVersionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *AvailableRuntimeVersions
}

// RuntimeVersionsListRuntimeVersions ...
func (c AppPlatformClient) RuntimeVersionsListRuntimeVersions(ctx context.Context) (result RuntimeVersionsListRuntimeVersionsOperationResponse, err error) {
	req, err := c.preparerForRuntimeVersionsListRuntimeVersions(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "RuntimeVersionsListRuntimeVersions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "RuntimeVersionsListRuntimeVersions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRuntimeVersionsListRuntimeVersions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "RuntimeVersionsListRuntimeVersions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRuntimeVersionsListRuntimeVersions prepares the RuntimeVersionsListRuntimeVersions request.
func (c AppPlatformClient) preparerForRuntimeVersionsListRuntimeVersions(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.AppPlatform/runtimeVersions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRuntimeVersionsListRuntimeVersions handles the response to the RuntimeVersionsListRuntimeVersions request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForRuntimeVersionsListRuntimeVersions(resp *http.Response) (result RuntimeVersionsListRuntimeVersionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
