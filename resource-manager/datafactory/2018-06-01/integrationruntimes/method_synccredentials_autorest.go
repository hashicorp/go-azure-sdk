package integrationruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncCredentialsOperationResponse struct {
	HttpResponse *http.Response
}

// SyncCredentials ...
func (c IntegrationRuntimesClient) SyncCredentials(ctx context.Context, id IntegrationRuntimeId) (result SyncCredentialsOperationResponse, err error) {
	req, err := c.preparerForSyncCredentials(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "SyncCredentials", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "SyncCredentials", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncCredentials(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "SyncCredentials", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncCredentials prepares the SyncCredentials request.
func (c IntegrationRuntimesClient) preparerForSyncCredentials(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
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

// responderForSyncCredentials handles the response to the SyncCredentials request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForSyncCredentials(resp *http.Response) (result SyncCredentialsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
