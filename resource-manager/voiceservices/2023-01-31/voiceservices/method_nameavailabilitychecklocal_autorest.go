package voiceservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NameAvailabilityCheckLocalOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckNameAvailabilityResponse
}

// NameAvailabilityCheckLocal ...
func (c VoiceservicesClient) NameAvailabilityCheckLocal(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (result NameAvailabilityCheckLocalOperationResponse, err error) {
	req, err := c.preparerForNameAvailabilityCheckLocal(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "voiceservices.VoiceservicesClient", "NameAvailabilityCheckLocal", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "voiceservices.VoiceservicesClient", "NameAvailabilityCheckLocal", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNameAvailabilityCheckLocal(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "voiceservices.VoiceservicesClient", "NameAvailabilityCheckLocal", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNameAvailabilityCheckLocal prepares the NameAvailabilityCheckLocal request.
func (c VoiceservicesClient) preparerForNameAvailabilityCheckLocal(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (*http.Request, error) {
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

// responderForNameAvailabilityCheckLocal handles the response to the NameAvailabilityCheckLocal request. The method always
// closes the http.Response Body.
func (c VoiceservicesClient) responderForNameAvailabilityCheckLocal(resp *http.Response) (result NameAvailabilityCheckLocalOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
