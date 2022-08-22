package notificationrecipientuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByNotificationOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecipientUserCollection
}

// ListByNotification ...
func (c NotificationRecipientUserClient) ListByNotification(ctx context.Context, id NotificationId) (result ListByNotificationOperationResponse, err error) {
	req, err := c.preparerForListByNotification(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationrecipientuser.NotificationRecipientUserClient", "ListByNotification", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationrecipientuser.NotificationRecipientUserClient", "ListByNotification", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByNotification(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationrecipientuser.NotificationRecipientUserClient", "ListByNotification", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByNotification prepares the ListByNotification request.
func (c NotificationRecipientUserClient) preparerForListByNotification(ctx context.Context, id NotificationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recipientUsers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByNotification handles the response to the ListByNotification request. The method always
// closes the http.Response Body.
func (c NotificationRecipientUserClient) responderForListByNotification(resp *http.Response) (result ListByNotificationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
