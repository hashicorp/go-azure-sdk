package key

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteRenderingAccountsRegenerateKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccountKeys
}

// RemoteRenderingAccountsRegenerateKeys ...
func (c KeyClient) RemoteRenderingAccountsRegenerateKeys(ctx context.Context, id RemoteRenderingAccountId, input AccountKeyRegenerateRequest) (result RemoteRenderingAccountsRegenerateKeysOperationResponse, err error) {
	req, err := c.preparerForRemoteRenderingAccountsRegenerateKeys(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsRegenerateKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsRegenerateKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRemoteRenderingAccountsRegenerateKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsRegenerateKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRemoteRenderingAccountsRegenerateKeys prepares the RemoteRenderingAccountsRegenerateKeys request.
func (c KeyClient) preparerForRemoteRenderingAccountsRegenerateKeys(ctx context.Context, id RemoteRenderingAccountId, input AccountKeyRegenerateRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateKeys", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRemoteRenderingAccountsRegenerateKeys handles the response to the RemoteRenderingAccountsRegenerateKeys request. The method always
// closes the http.Response Body.
func (c KeyClient) responderForRemoteRenderingAccountsRegenerateKeys(resp *http.Response) (result RemoteRenderingAccountsRegenerateKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
