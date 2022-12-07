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

type BuildServiceGetBuildResultLogOperationResponse struct {
	HttpResponse *http.Response
	Model        *BuildResultLog
}

// BuildServiceGetBuildResultLog ...
func (c AppPlatformClient) BuildServiceGetBuildResultLog(ctx context.Context, id ResultId) (result BuildServiceGetBuildResultLogOperationResponse, err error) {
	req, err := c.preparerForBuildServiceGetBuildResultLog(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildResultLog", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildResultLog", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildServiceGetBuildResultLog(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceGetBuildResultLog", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildServiceGetBuildResultLog prepares the BuildServiceGetBuildResultLog request.
func (c AppPlatformClient) preparerForBuildServiceGetBuildResultLog(ctx context.Context, id ResultId) (*http.Request, error) {
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

// responderForBuildServiceGetBuildResultLog handles the response to the BuildServiceGetBuildResultLog request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceGetBuildResultLog(resp *http.Response) (result BuildServiceGetBuildResultLogOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
