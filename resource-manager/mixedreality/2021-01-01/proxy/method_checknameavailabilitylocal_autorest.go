package proxy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityLocalOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckNameAvailabilityResponse
}

// CheckNameAvailabilityLocal ...
func (c ProxyClient) CheckNameAvailabilityLocal(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (result CheckNameAvailabilityLocalOperationResponse, err error) {
	req, err := c.preparerForCheckNameAvailabilityLocal(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "CheckNameAvailabilityLocal", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "CheckNameAvailabilityLocal", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckNameAvailabilityLocal(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "CheckNameAvailabilityLocal", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckNameAvailabilityLocal prepares the CheckNameAvailabilityLocal request.
func (c ProxyClient) preparerForCheckNameAvailabilityLocal(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (*http.Request, error) {
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

// responderForCheckNameAvailabilityLocal handles the response to the CheckNameAvailabilityLocal request. The method always
// closes the http.Response Body.
func (c ProxyClient) responderForCheckNameAvailabilityLocal(resp *http.Response) (result CheckNameAvailabilityLocalOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
