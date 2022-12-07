package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewaysGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *GatewayResource
}

// GatewaysGet ...
func (c AppPlatformClient) GatewaysGet(ctx context.Context, id GatewayId) (result GatewaysGetOperationResponse, err error) {
	req, err := c.preparerForGatewaysGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewaysGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewaysGet prepares the GatewaysGet request.
func (c AppPlatformClient) preparerForGatewaysGet(ctx context.Context, id GatewayId) (*http.Request, error) {
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

// responderForGatewaysGet handles the response to the GatewaysGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewaysGet(resp *http.Response) (result GatewaysGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
