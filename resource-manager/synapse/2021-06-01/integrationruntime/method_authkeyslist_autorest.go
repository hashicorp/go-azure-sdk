package integrationruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthKeysListOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeAuthKeys
}

// AuthKeysList ...
func (c IntegrationRuntimeClient) AuthKeysList(ctx context.Context, id IntegrationRuntimeId) (result AuthKeysListOperationResponse, err error) {
	req, err := c.preparerForAuthKeysList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAuthKeysList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAuthKeysList prepares the AuthKeysList request.
func (c IntegrationRuntimeClient) preparerForAuthKeysList(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listAuthKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAuthKeysList handles the response to the AuthKeysList request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForAuthKeysList(resp *http.Response) (result AuthKeysListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
