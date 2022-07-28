package checkfrontdoornameavailability

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FrontDoorNameAvailabilityCheckOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckNameAvailabilityOutput
}

// FrontDoorNameAvailabilityCheck ...
func (c CheckFrontDoorNameAvailabilityClient) FrontDoorNameAvailabilityCheck(ctx context.Context, input CheckNameAvailabilityInput) (result FrontDoorNameAvailabilityCheckOperationResponse, err error) {
	req, err := c.preparerForFrontDoorNameAvailabilityCheck(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient", "FrontDoorNameAvailabilityCheck", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient", "FrontDoorNameAvailabilityCheck", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFrontDoorNameAvailabilityCheck(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient", "FrontDoorNameAvailabilityCheck", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFrontDoorNameAvailabilityCheck prepares the FrontDoorNameAvailabilityCheck request.
func (c CheckFrontDoorNameAvailabilityClient) preparerForFrontDoorNameAvailabilityCheck(ctx context.Context, input CheckNameAvailabilityInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Network/checkFrontDoorNameAvailability"),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFrontDoorNameAvailabilityCheck handles the response to the FrontDoorNameAvailabilityCheck request. The method always
// closes the http.Response Body.
func (c CheckFrontDoorNameAvailabilityClient) responderForFrontDoorNameAvailabilityCheck(resp *http.Response) (result FrontDoorNameAvailabilityCheckOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
