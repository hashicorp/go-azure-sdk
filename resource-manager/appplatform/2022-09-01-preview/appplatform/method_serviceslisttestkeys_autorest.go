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

type ServicesListTestKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *TestKeys
}

// ServicesListTestKeys ...
func (c AppPlatformClient) ServicesListTestKeys(ctx context.Context, id SpringId) (result ServicesListTestKeysOperationResponse, err error) {
	req, err := c.preparerForServicesListTestKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesListTestKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesListTestKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesListTestKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesListTestKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesListTestKeys prepares the ServicesListTestKeys request.
func (c AppPlatformClient) preparerForServicesListTestKeys(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listTestKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServicesListTestKeys handles the response to the ServicesListTestKeys request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForServicesListTestKeys(resp *http.Response) (result ServicesListTestKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
