package customoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicesNestedCheckNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *NameAvailabilityResponse
}

// ServicesNestedCheckNameAvailability ...
func (c CustomOperationClient) ServicesNestedCheckNameAvailability(ctx context.Context, id ServiceId, input NameAvailabilityRequest) (result ServicesNestedCheckNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForServicesNestedCheckNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "ServicesNestedCheckNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "ServicesNestedCheckNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesNestedCheckNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customoperation.CustomOperationClient", "ServicesNestedCheckNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesNestedCheckNameAvailability prepares the ServicesNestedCheckNameAvailability request.
func (c CustomOperationClient) preparerForServicesNestedCheckNameAvailability(ctx context.Context, id ServiceId, input NameAvailabilityRequest) (*http.Request, error) {
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

// responderForServicesNestedCheckNameAvailability handles the response to the ServicesNestedCheckNameAvailability request. The method always
// closes the http.Response Body.
func (c CustomOperationClient) responderForServicesNestedCheckNameAvailability(resp *http.Response) (result ServicesNestedCheckNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
