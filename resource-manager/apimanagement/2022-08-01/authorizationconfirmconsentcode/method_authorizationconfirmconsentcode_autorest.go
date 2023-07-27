package authorizationconfirmconsentcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationConfirmConsentCodeOperationResponse struct {
	HttpResponse *http.Response
}

// AuthorizationConfirmConsentCode ...
func (c AuthorizationConfirmConsentCodeClient) AuthorizationConfirmConsentCode(ctx context.Context, id AuthorizationId, input AuthorizationConfirmConsentCodeRequestContract) (result AuthorizationConfirmConsentCodeOperationResponse, err error) {
	req, err := c.preparerForAuthorizationConfirmConsentCode(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationconfirmconsentcode.AuthorizationConfirmConsentCodeClient", "AuthorizationConfirmConsentCode", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationconfirmconsentcode.AuthorizationConfirmConsentCodeClient", "AuthorizationConfirmConsentCode", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAuthorizationConfirmConsentCode(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationconfirmconsentcode.AuthorizationConfirmConsentCodeClient", "AuthorizationConfirmConsentCode", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAuthorizationConfirmConsentCode prepares the AuthorizationConfirmConsentCode request.
func (c AuthorizationConfirmConsentCodeClient) preparerForAuthorizationConfirmConsentCode(ctx context.Context, id AuthorizationId, input AuthorizationConfirmConsentCodeRequestContract) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/confirmConsentCode", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAuthorizationConfirmConsentCode handles the response to the AuthorizationConfirmConsentCode request. The method always
// closes the http.Response Body.
func (c AuthorizationConfirmConsentCodeClient) responderForAuthorizationConfirmConsentCode(resp *http.Response) (result AuthorizationConfirmConsentCodeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
