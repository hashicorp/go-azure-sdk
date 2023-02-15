package usertoken

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserGetSharedAccessTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *UserTokenResult
}

// UserGetSharedAccessToken ...
func (c UserTokenClient) UserGetSharedAccessToken(ctx context.Context, id UserId, input UserTokenParameters) (result UserGetSharedAccessTokenOperationResponse, err error) {
	req, err := c.preparerForUserGetSharedAccessToken(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usertoken.UserTokenClient", "UserGetSharedAccessToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "usertoken.UserTokenClient", "UserGetSharedAccessToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUserGetSharedAccessToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usertoken.UserTokenClient", "UserGetSharedAccessToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUserGetSharedAccessToken prepares the UserGetSharedAccessToken request.
func (c UserTokenClient) preparerForUserGetSharedAccessToken(ctx context.Context, id UserId, input UserTokenParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/token", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUserGetSharedAccessToken handles the response to the UserGetSharedAccessToken request. The method always
// closes the http.Response Body.
func (c UserTokenClient) responderForUserGetSharedAccessToken(resp *http.Response) (result UserGetSharedAccessTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
