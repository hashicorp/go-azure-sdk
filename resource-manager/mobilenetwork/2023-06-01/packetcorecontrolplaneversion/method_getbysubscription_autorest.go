package packetcorecontrolplaneversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *PacketCoreControlPlaneVersion
}

// GetBySubscription ...
func (c PacketCoreControlPlaneVersionClient) GetBySubscription(ctx context.Context, id ProviderPacketCoreControlPlaneVersionId) (result GetBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForGetBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplaneversion.PacketCoreControlPlaneVersionClient", "GetBySubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplaneversion.PacketCoreControlPlaneVersionClient", "GetBySubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBySubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplaneversion.PacketCoreControlPlaneVersionClient", "GetBySubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBySubscription prepares the GetBySubscription request.
func (c PacketCoreControlPlaneVersionClient) preparerForGetBySubscription(ctx context.Context, id ProviderPacketCoreControlPlaneVersionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBySubscription handles the response to the GetBySubscription request. The method always
// closes the http.Response Body.
func (c PacketCoreControlPlaneVersionClient) responderForGetBySubscription(resp *http.Response) (result GetBySubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
