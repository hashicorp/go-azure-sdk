package subscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckZonePeersOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckZonePeersResult
}

// CheckZonePeers ...
func (c SubscriptionsClient) CheckZonePeers(ctx context.Context, id commonids.SubscriptionId, input CheckZonePeersRequest) (result CheckZonePeersOperationResponse, err error) {
	req, err := c.preparerForCheckZonePeers(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CheckZonePeers", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CheckZonePeers", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckZonePeers(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CheckZonePeers", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckZonePeers prepares the CheckZonePeers request.
func (c SubscriptionsClient) preparerForCheckZonePeers(ctx context.Context, id commonids.SubscriptionId, input CheckZonePeersRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Resources/checkZonePeers", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckZonePeers handles the response to the CheckZonePeers request. The method always
// closes the http.Response Body.
func (c SubscriptionsClient) responderForCheckZonePeers(resp *http.Response) (result CheckZonePeersOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
