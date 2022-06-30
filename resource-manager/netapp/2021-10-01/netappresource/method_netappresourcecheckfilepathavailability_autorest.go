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

type NetAppResourceCheckFilePathAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckAvailabilityResponse
}

// NetAppResourceCheckFilePathAvailability ...
func (c NetAppResourceClient) NetAppResourceCheckFilePathAvailability(ctx context.Context, id LocationId, input FilePathAvailabilityRequest) (result NetAppResourceCheckFilePathAvailabilityOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceCheckFilePathAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckFilePathAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckFilePathAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceCheckFilePathAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckFilePathAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceCheckFilePathAvailability prepares the NetAppResourceCheckFilePathAvailability request.
func (c NetAppResourceClient) preparerForNetAppResourceCheckFilePathAvailability(ctx context.Context, id LocationId, input FilePathAvailabilityRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/checkFilePathAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceCheckFilePathAvailability handles the response to the NetAppResourceCheckFilePathAvailability request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceCheckFilePathAvailability(resp *http.Response) (result NetAppResourceCheckFilePathAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
