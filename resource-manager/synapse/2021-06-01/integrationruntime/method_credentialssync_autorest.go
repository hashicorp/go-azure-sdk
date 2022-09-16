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

type CredentialsSyncOperationResponse struct {
	HttpResponse *http.Response
}

// CredentialsSync ...
func (c IntegrationRuntimeClient) CredentialsSync(ctx context.Context, id IntegrationRuntimeId) (result CredentialsSyncOperationResponse, err error) {
	req, err := c.preparerForCredentialsSync(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "CredentialsSync", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "CredentialsSync", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCredentialsSync(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "CredentialsSync", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCredentialsSync prepares the CredentialsSync request.
func (c IntegrationRuntimeClient) preparerForCredentialsSync(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/syncCredentials", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCredentialsSync handles the response to the CredentialsSync request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForCredentialsSync(resp *http.Response) (result CredentialsSyncOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
