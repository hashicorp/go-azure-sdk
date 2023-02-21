package checknameavailability

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DigitalTwinsCheckNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckNameResult
}

// DigitalTwinsCheckNameAvailability ...
func (c CheckNameAvailabilityClient) DigitalTwinsCheckNameAvailability(ctx context.Context, id LocationId, input CheckNameRequest) (result DigitalTwinsCheckNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForDigitalTwinsCheckNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "DigitalTwinsCheckNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "DigitalTwinsCheckNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDigitalTwinsCheckNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "DigitalTwinsCheckNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDigitalTwinsCheckNameAvailability prepares the DigitalTwinsCheckNameAvailability request.
func (c CheckNameAvailabilityClient) preparerForDigitalTwinsCheckNameAvailability(ctx context.Context, id LocationId, input CheckNameRequest) (*http.Request, error) {
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

// responderForDigitalTwinsCheckNameAvailability handles the response to the DigitalTwinsCheckNameAvailability request. The method always
// closes the http.Response Body.
func (c CheckNameAvailabilityClient) responderForDigitalTwinsCheckNameAvailability(resp *http.Response) (result DigitalTwinsCheckNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
