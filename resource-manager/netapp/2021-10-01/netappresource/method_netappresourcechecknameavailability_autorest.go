package netappresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceCheckNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckAvailabilityResponse
}

// NetAppResourceCheckNameAvailability ...
func (c NetAppResourceClient) NetAppResourceCheckNameAvailability(ctx context.Context, id LocationId, input ResourceNameAvailabilityRequest) (result NetAppResourceCheckNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceCheckNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceCheckNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceCheckNameAvailability prepares the NetAppResourceCheckNameAvailability request.
func (c NetAppResourceClient) preparerForNetAppResourceCheckNameAvailability(ctx context.Context, id LocationId, input ResourceNameAvailabilityRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/checkNameAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceCheckNameAvailability handles the response to the NetAppResourceCheckNameAvailability request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceCheckNameAvailability(resp *http.Response) (result NetAppResourceCheckNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
