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

type DisconnectOperationResponse struct {
	HttpResponse *http.Response
}

// Disconnect ...
func (c UserSessionClient) Disconnect(ctx context.Context, id UserSessionId) (result DisconnectOperationResponse, err error) {
	req, err := c.preparerForDisconnect(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "Disconnect", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "Disconnect", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisconnect(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "Disconnect", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisconnect prepares the Disconnect request.
func (c UserSessionClient) preparerForDisconnect(ctx context.Context, id UserSessionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disconnect", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisconnect handles the response to the Disconnect request. The method always
// closes the http.Response Body.
func (c UserSessionClient) responderForDisconnect(resp *http.Response) (result DisconnectOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
