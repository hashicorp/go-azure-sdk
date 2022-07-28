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

type SpatialAnchorsAccountsRegenerateKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccountKeys
}

// SpatialAnchorsAccountsRegenerateKeys ...
func (c KeyClient) SpatialAnchorsAccountsRegenerateKeys(ctx context.Context, id SpatialAnchorsAccountId, input AccountKeyRegenerateRequest) (result SpatialAnchorsAccountsRegenerateKeysOperationResponse, err error) {
	req, err := c.preparerForSpatialAnchorsAccountsRegenerateKeys(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsRegenerateKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsRegenerateKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSpatialAnchorsAccountsRegenerateKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsRegenerateKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSpatialAnchorsAccountsRegenerateKeys prepares the SpatialAnchorsAccountsRegenerateKeys request.
func (c KeyClient) preparerForSpatialAnchorsAccountsRegenerateKeys(ctx context.Context, id SpatialAnchorsAccountId, input AccountKeyRegenerateRequest) (*http.Request, error) {
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

// responderForSpatialAnchorsAccountsRegenerateKeys handles the response to the SpatialAnchorsAccountsRegenerateKeys request. The method always
// closes the http.Response Body.
func (c KeyClient) responderForSpatialAnchorsAccountsRegenerateKeys(resp *http.Response) (result SpatialAnchorsAccountsRegenerateKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
