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

type NetAppResourceCheckQuotaAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckAvailabilityResponse
}

// NetAppResourceCheckQuotaAvailability ...
func (c NetAppResourceClient) NetAppResourceCheckQuotaAvailability(ctx context.Context, id LocationId, input QuotaAvailabilityRequest) (result NetAppResourceCheckQuotaAvailabilityOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceCheckQuotaAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckQuotaAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckQuotaAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceCheckQuotaAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceCheckQuotaAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceCheckQuotaAvailability prepares the NetAppResourceCheckQuotaAvailability request.
func (c NetAppResourceClient) preparerForNetAppResourceCheckQuotaAvailability(ctx context.Context, id LocationId, input QuotaAvailabilityRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/checkQuotaAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceCheckQuotaAvailability handles the response to the NetAppResourceCheckQuotaAvailability request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceCheckQuotaAvailability(resp *http.Response) (result NetAppResourceCheckQuotaAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
