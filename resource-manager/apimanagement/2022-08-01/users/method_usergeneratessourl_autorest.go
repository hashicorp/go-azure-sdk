package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserGenerateSsoUrlOperationResponse struct {
	HttpResponse *http.Response
	Model        *GenerateSsoUrlResult
}

// UserGenerateSsoUrl ...
func (c UsersClient) UserGenerateSsoUrl(ctx context.Context, id UserId) (result UserGenerateSsoUrlOperationResponse, err error) {
	req, err := c.preparerForUserGenerateSsoUrl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "users.UsersClient", "UserGenerateSsoUrl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "users.UsersClient", "UserGenerateSsoUrl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUserGenerateSsoUrl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "users.UsersClient", "UserGenerateSsoUrl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUserGenerateSsoUrl prepares the UserGenerateSsoUrl request.
func (c UsersClient) preparerForUserGenerateSsoUrl(ctx context.Context, id UserId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/generateSsoUrl", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUserGenerateSsoUrl handles the response to the UserGenerateSsoUrl request. The method always
// closes the http.Response Body.
func (c UsersClient) responderForUserGenerateSsoUrl(resp *http.Response) (result UserGenerateSsoUrlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
