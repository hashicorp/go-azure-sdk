package standardoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicesCheckNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *NameAvailabilityResponse
}

// ServicesCheckNameAvailability ...
func (c StandardOperationClient) ServicesCheckNameAvailability(ctx context.Context, id LocationId, input NameAvailabilityRequest) (result ServicesCheckNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForServicesCheckNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServicesCheckNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServicesCheckNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesCheckNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ServicesCheckNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesCheckNameAvailability prepares the ServicesCheckNameAvailability request.
func (c StandardOperationClient) preparerForServicesCheckNameAvailability(ctx context.Context, id LocationId, input NameAvailabilityRequest) (*http.Request, error) {
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

// responderForServicesCheckNameAvailability handles the response to the ServicesCheckNameAvailability request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForServicesCheckNameAvailability(resp *http.Response) (result ServicesCheckNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
