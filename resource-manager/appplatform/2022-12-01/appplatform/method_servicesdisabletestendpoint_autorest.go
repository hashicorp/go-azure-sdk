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

type ServicesDisableTestEndpointOperationResponse struct {
	HttpResponse *http.Response
}

// ServicesDisableTestEndpoint ...
func (c AppPlatformClient) ServicesDisableTestEndpoint(ctx context.Context, id SpringId) (result ServicesDisableTestEndpointOperationResponse, err error) {
	req, err := c.preparerForServicesDisableTestEndpoint(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesDisableTestEndpoint", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesDisableTestEndpoint", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesDisableTestEndpoint(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesDisableTestEndpoint", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesDisableTestEndpoint prepares the ServicesDisableTestEndpoint request.
func (c AppPlatformClient) preparerForServicesDisableTestEndpoint(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disableTestEndpoint", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServicesDisableTestEndpoint handles the response to the ServicesDisableTestEndpoint request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForServicesDisableTestEndpoint(resp *http.Response) (result ServicesDisableTestEndpointOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
