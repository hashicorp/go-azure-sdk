package triggers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEventSubscriptionStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggerSubscriptionOperationStatus
}

// GetEventSubscriptionStatus ...
func (c TriggersClient) GetEventSubscriptionStatus(ctx context.Context, id TriggerId) (result GetEventSubscriptionStatusOperationResponse, err error) {
	req, err := c.preparerForGetEventSubscriptionStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "GetEventSubscriptionStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "GetEventSubscriptionStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetEventSubscriptionStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "GetEventSubscriptionStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetEventSubscriptionStatus prepares the GetEventSubscriptionStatus request.
func (c TriggersClient) preparerForGetEventSubscriptionStatus(ctx context.Context, id TriggerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getEventSubscriptionStatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetEventSubscriptionStatus handles the response to the GetEventSubscriptionStatus request. The method always
// closes the http.Response Body.
func (c TriggersClient) responderForGetEventSubscriptionStatus(resp *http.Response) (result GetEventSubscriptionStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
