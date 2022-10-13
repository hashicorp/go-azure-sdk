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

type ServicesEnableTestEndpointOperationResponse struct {
	HttpResponse *http.Response
	Model        *TestKeys
}

// ServicesEnableTestEndpoint ...
func (c AppPlatformClient) ServicesEnableTestEndpoint(ctx context.Context, id SpringId) (result ServicesEnableTestEndpointOperationResponse, err error) {
	req, err := c.preparerForServicesEnableTestEndpoint(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesEnableTestEndpoint", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesEnableTestEndpoint", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesEnableTestEndpoint(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesEnableTestEndpoint", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesEnableTestEndpoint prepares the ServicesEnableTestEndpoint request.
func (c AppPlatformClient) preparerForServicesEnableTestEndpoint(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/enableTestEndpoint", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServicesEnableTestEndpoint handles the response to the ServicesEnableTestEndpoint request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForServicesEnableTestEndpoint(resp *http.Response) (result ServicesEnableTestEndpointOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
