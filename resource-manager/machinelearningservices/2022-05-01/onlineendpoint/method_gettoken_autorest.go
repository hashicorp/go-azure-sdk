package onlineendpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *EndpointAuthToken
}

// GetToken ...
func (c OnlineEndpointClient) GetToken(ctx context.Context, id OnlineEndpointId) (result GetTokenOperationResponse, err error) {
	req, err := c.preparerForGetToken(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onlineendpoint.OnlineEndpointClient", "GetToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "onlineendpoint.OnlineEndpointClient", "GetToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onlineendpoint.OnlineEndpointClient", "GetToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetToken prepares the GetToken request.
func (c OnlineEndpointClient) preparerForGetToken(ctx context.Context, id OnlineEndpointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/token", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetToken handles the response to the GetToken request. The method always
// closes the http.Response Body.
func (c OnlineEndpointClient) responderForGetToken(resp *http.Response) (result GetTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
