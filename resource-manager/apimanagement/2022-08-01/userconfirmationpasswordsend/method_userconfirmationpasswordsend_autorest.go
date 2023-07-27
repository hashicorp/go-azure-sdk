package userconfirmationpasswordsend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserConfirmationPasswordSendOperationResponse struct {
	HttpResponse *http.Response
}

type UserConfirmationPasswordSendOperationOptions struct {
	AppType *AppType
}

func DefaultUserConfirmationPasswordSendOperationOptions() UserConfirmationPasswordSendOperationOptions {
	return UserConfirmationPasswordSendOperationOptions{}
}

func (o UserConfirmationPasswordSendOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o UserConfirmationPasswordSendOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.AppType != nil {
		out["appType"] = *o.AppType
	}

	return out
}

// UserConfirmationPasswordSend ...
func (c UserConfirmationPasswordSendClient) UserConfirmationPasswordSend(ctx context.Context, id UserId, options UserConfirmationPasswordSendOperationOptions) (result UserConfirmationPasswordSendOperationResponse, err error) {
	req, err := c.preparerForUserConfirmationPasswordSend(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "userconfirmationpasswordsend.UserConfirmationPasswordSendClient", "UserConfirmationPasswordSend", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "userconfirmationpasswordsend.UserConfirmationPasswordSendClient", "UserConfirmationPasswordSend", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUserConfirmationPasswordSend(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "userconfirmationpasswordsend.UserConfirmationPasswordSendClient", "UserConfirmationPasswordSend", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUserConfirmationPasswordSend prepares the UserConfirmationPasswordSend request.
func (c UserConfirmationPasswordSendClient) preparerForUserConfirmationPasswordSend(ctx context.Context, id UserId, options UserConfirmationPasswordSendOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/confirmations/password/send", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUserConfirmationPasswordSend handles the response to the UserConfirmationPasswordSend request. The method always
// closes the http.Response Body.
func (c UserConfirmationPasswordSendClient) responderForUserConfirmationPasswordSend(resp *http.Response) (result UserConfirmationPasswordSendOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
