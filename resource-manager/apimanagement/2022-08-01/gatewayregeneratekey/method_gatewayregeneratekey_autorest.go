package gatewayregeneratekey

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayRegenerateKeyOperationResponse struct {
	HttpResponse *http.Response
}

// GatewayRegenerateKey ...
func (c GatewayRegenerateKeyClient) GatewayRegenerateKey(ctx context.Context, id GatewayId, input GatewayKeyRegenerationRequestContract) (result GatewayRegenerateKeyOperationResponse, err error) {
	req, err := c.preparerForGatewayRegenerateKey(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewayregeneratekey.GatewayRegenerateKeyClient", "GatewayRegenerateKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewayregeneratekey.GatewayRegenerateKeyClient", "GatewayRegenerateKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewayRegenerateKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewayregeneratekey.GatewayRegenerateKeyClient", "GatewayRegenerateKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewayRegenerateKey prepares the GatewayRegenerateKey request.
func (c GatewayRegenerateKeyClient) preparerForGatewayRegenerateKey(ctx context.Context, id GatewayId, input GatewayKeyRegenerationRequestContract) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateKey", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGatewayRegenerateKey handles the response to the GatewayRegenerateKey request. The method always
// closes the http.Response Body.
func (c GatewayRegenerateKeyClient) responderForGatewayRegenerateKey(resp *http.Response) (result GatewayRegenerateKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
