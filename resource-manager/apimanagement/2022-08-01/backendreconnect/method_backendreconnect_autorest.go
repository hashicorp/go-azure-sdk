package backendreconnect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackendReconnectOperationResponse struct {
	HttpResponse *http.Response
}

// BackendReconnect ...
func (c BackendReconnectClient) BackendReconnect(ctx context.Context, id BackendId, input BackendReconnectContract) (result BackendReconnectOperationResponse, err error) {
	req, err := c.preparerForBackendReconnect(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backendreconnect.BackendReconnectClient", "BackendReconnect", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backendreconnect.BackendReconnectClient", "BackendReconnect", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackendReconnect(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backendreconnect.BackendReconnectClient", "BackendReconnect", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackendReconnect prepares the BackendReconnect request.
func (c BackendReconnectClient) preparerForBackendReconnect(ctx context.Context, id BackendId, input BackendReconnectContract) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reconnect", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackendReconnect handles the response to the BackendReconnect request. The method always
// closes the http.Response Body.
func (c BackendReconnectClient) responderForBackendReconnect(resp *http.Response) (result BackendReconnectOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
