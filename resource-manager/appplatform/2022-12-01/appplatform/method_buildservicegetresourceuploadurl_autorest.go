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

type BuildServiceGetResourceUploadUrlOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceUploadDefinition
}

// BuildServiceGetResourceUploadUrl ...
func (c AppPlatformClient) BuildServiceGetResourceUploadUrl(ctx context.Context, id BuildServiceId) (result BuildServiceGetResourceUploadUrlOperationResponse, err error) {
	req, err := c.preparerForBuildServiceGetResourceUploadUrl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetResourceUploadUrl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetResourceUploadUrl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceGetResourceUploadUrl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetResourceUploadUrl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceGetResourceUploadUrl prepares the BuildServiceGetResourceUploadUrl request.
func (c AppPlatformClient) preparerForBuildServiceGetResourceUploadUrl(ctx context.Context, id BuildServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getResourceUploadUrl", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBuildServiceGetResourceUploadUrl handles the response to the BuildServiceGetResourceUploadUrl request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceGetResourceUploadUrl(resp *http.Response) (result BuildServiceGetResourceUploadUrlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
