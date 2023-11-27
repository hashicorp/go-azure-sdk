package itemlevelrecoveryconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeOperationResponse struct {
	HttpResponse *http.Response
}

// Revoke ...
func (c ItemLevelRecoveryConnectionsClient) Revoke(ctx context.Context, id RecoveryPointId) (result RevokeOperationResponse, err error) {
	req, err := c.preparerForRevoke(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Revoke", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Revoke", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRevoke(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Revoke", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRevoke prepares the Revoke request.
func (c ItemLevelRecoveryConnectionsClient) preparerForRevoke(ctx context.Context, id RecoveryPointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/revokeInstantItemRecovery", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRevoke handles the response to the Revoke request. The method always
// closes the http.Response Body.
func (c ItemLevelRecoveryConnectionsClient) responderForRevoke(resp *http.Response) (result RevokeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
