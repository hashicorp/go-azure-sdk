package protectionintentresources

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntentDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// ProtectionIntentDelete ...
func (c ProtectionIntentResourcesClient) ProtectionIntentDelete(ctx context.Context, id BackupProtectionIntentId) (result ProtectionIntentDeleteOperationResponse, err error) {
	req, err := c.preparerForProtectionIntentDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProtectionIntentDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "protectionintentresources.ProtectionIntentResourcesClient", "ProtectionIntentDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProtectionIntentDelete prepares the ProtectionIntentDelete request.
func (c ProtectionIntentResourcesClient) preparerForProtectionIntentDelete(ctx context.Context, id BackupProtectionIntentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProtectionIntentDelete handles the response to the ProtectionIntentDelete request. The method always
// closes the http.Response Body.
func (c ProtectionIntentResourcesClient) responderForProtectionIntentDelete(resp *http.Response) (result ProtectionIntentDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
