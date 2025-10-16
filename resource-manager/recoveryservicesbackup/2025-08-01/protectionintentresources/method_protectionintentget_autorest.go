package protectionintentresources

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntentGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProtectionIntentResource
}

// ProtectionIntentGet ...
func (c ProtectionIntentResourcesClient) ProtectionIntentGet(ctx context.Context, id BackupProtectionIntentId) (result ProtectionIntentGetOperationResponse, err error) {
	req, err := c.preparerForProtectionIntentGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProtectionIntentGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProtectionIntentGet prepares the ProtectionIntentGet request.
func (c ProtectionIntentResourcesClient) preparerForProtectionIntentGet(ctx context.Context, id BackupProtectionIntentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProtectionIntentGet handles the response to the ProtectionIntentGet request. The method always
// closes the http.Response Body.
func (c ProtectionIntentResourcesClient) responderForProtectionIntentGet(resp *http.Response) (result ProtectionIntentGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
