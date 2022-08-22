package gatewaylistkeys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayListKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *GatewayKeysContract
}

// GatewayListKeys ...
func (c GatewayListKeysClient) GatewayListKeys(ctx context.Context, id GatewayId) (result GatewayListKeysOperationResponse, err error) {
	req, err := c.preparerForGatewayListKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaylistkeys.GatewayListKeysClient", "GatewayListKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaylistkeys.GatewayListKeysClient", "GatewayListKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewayListKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaylistkeys.GatewayListKeysClient", "GatewayListKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewayListKeys prepares the GatewayListKeys request.
func (c GatewayListKeysClient) preparerForGatewayListKeys(ctx context.Context, id GatewayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGatewayListKeys handles the response to the GatewayListKeys request. The method always
// closes the http.Response Body.
func (c GatewayListKeysClient) responderForGatewayListKeys(resp *http.Response) (result GatewayListKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
