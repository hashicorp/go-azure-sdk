package usersession

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendMessageOperationResponse struct {
	HttpResponse *http.Response
}

// SendMessage ...
func (c UserSessionClient) SendMessage(ctx context.Context, id UserSessionId, input SendMessage) (result SendMessageOperationResponse, err error) {
	req, err := c.preparerForSendMessage(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "SendMessage", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "SendMessage", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSendMessage(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "SendMessage", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSendMessage prepares the SendMessage request.
func (c UserSessionClient) preparerForSendMessage(ctx context.Context, id UserSessionId, input SendMessage) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sendMessage", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSendMessage handles the response to the SendMessage request. The method always
// closes the http.Response Body.
func (c UserSessionClient) responderForSendMessage(resp *http.Response) (result SendMessageOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
