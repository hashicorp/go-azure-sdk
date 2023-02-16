package subscription

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSubscriptionGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionContract
}

// UserSubscriptionGet ...
func (c SubscriptionClient) UserSubscriptionGet(ctx context.Context, id Subscriptions2Id) (result UserSubscriptionGetOperationResponse, err error) {
	req, err := c.preparerForUserSubscriptionGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.SubscriptionClient", "UserSubscriptionGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.SubscriptionClient", "UserSubscriptionGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUserSubscriptionGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.SubscriptionClient", "UserSubscriptionGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUserSubscriptionGet prepares the UserSubscriptionGet request.
func (c SubscriptionClient) preparerForUserSubscriptionGet(ctx context.Context, id Subscriptions2Id) (*http.Request, error) {
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

// responderForUserSubscriptionGet handles the response to the UserSubscriptionGet request. The method always
// closes the http.Response Body.
func (c SubscriptionClient) responderForUserSubscriptionGet(resp *http.Response) (result UserSubscriptionGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
