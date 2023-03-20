package recoveryservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilitiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *CapabilitiesResponse
}

// Capabilities ...
func (c RecoveryServicesClient) Capabilities(ctx context.Context, id LocationId, input ResourceCapabilities) (result CapabilitiesOperationResponse, err error) {
	req, err := c.preparerForCapabilities(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.RecoveryServicesClient", "Capabilities", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.RecoveryServicesClient", "Capabilities", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCapabilities(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.RecoveryServicesClient", "Capabilities", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCapabilities prepares the Capabilities request.
func (c RecoveryServicesClient) preparerForCapabilities(ctx context.Context, id LocationId, input ResourceCapabilities) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capabilities", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCapabilities handles the response to the Capabilities request. The method always
// closes the http.Response Body.
func (c RecoveryServicesClient) responderForCapabilities(resp *http.Response) (result CapabilitiesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
