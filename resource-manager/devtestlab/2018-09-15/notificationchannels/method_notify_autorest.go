package notificationchannels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotifyOperationResponse struct {
	HttpResponse *http.Response
}

// Notify ...
func (c NotificationChannelsClient) Notify(ctx context.Context, id NotificationChannelId, input NotifyParameters) (result NotifyOperationResponse, err error) {
	req, err := c.preparerForNotify(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationchannels.NotificationChannelsClient", "Notify", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationchannels.NotificationChannelsClient", "Notify", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNotify(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationchannels.NotificationChannelsClient", "Notify", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNotify prepares the Notify request.
func (c NotificationChannelsClient) preparerForNotify(ctx context.Context, id NotificationChannelId, input NotifyParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notify", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNotify handles the response to the Notify request. The method always
// closes the http.Response Body.
func (c NotificationChannelsClient) responderForNotify(resp *http.Response) (result NotifyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
