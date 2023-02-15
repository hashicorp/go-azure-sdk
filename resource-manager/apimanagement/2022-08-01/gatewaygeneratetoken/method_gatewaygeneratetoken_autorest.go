package gatewaygeneratetoken

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayGenerateTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *GatewayTokenContract
}

// GatewayGenerateToken ...
func (c GatewayGenerateTokenClient) GatewayGenerateToken(ctx context.Context, id GatewayId, input GatewayTokenRequestContract) (result GatewayGenerateTokenOperationResponse, err error) {
	req, err := c.preparerForGatewayGenerateToken(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaygeneratetoken.GatewayGenerateTokenClient", "GatewayGenerateToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaygeneratetoken.GatewayGenerateTokenClient", "GatewayGenerateToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewayGenerateToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "gatewaygeneratetoken.GatewayGenerateTokenClient", "GatewayGenerateToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewayGenerateToken prepares the GatewayGenerateToken request.
func (c GatewayGenerateTokenClient) preparerForGatewayGenerateToken(ctx context.Context, id GatewayId, input GatewayTokenRequestContract) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/generateToken", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGatewayGenerateToken handles the response to the GatewayGenerateToken request. The method always
// closes the http.Response Body.
func (c GatewayGenerateTokenClient) responderForGatewayGenerateToken(resp *http.Response) (result GatewayGenerateTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
