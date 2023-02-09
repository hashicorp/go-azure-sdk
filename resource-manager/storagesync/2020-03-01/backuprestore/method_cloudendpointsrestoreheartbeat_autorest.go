package backuprestore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointsrestoreheartbeatOperationResponse struct {
	HttpResponse *http.Response
}

// CloudEndpointsrestoreheartbeat ...
func (c BackupRestoreClient) CloudEndpointsrestoreheartbeat(ctx context.Context, id CloudEndpointId) (result CloudEndpointsrestoreheartbeatOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsrestoreheartbeat(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuprestore.BackupRestoreClient", "CloudEndpointsrestoreheartbeat", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuprestore.BackupRestoreClient", "CloudEndpointsrestoreheartbeat", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCloudEndpointsrestoreheartbeat(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuprestore.BackupRestoreClient", "CloudEndpointsrestoreheartbeat", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCloudEndpointsrestoreheartbeat prepares the CloudEndpointsrestoreheartbeat request.
func (c BackupRestoreClient) preparerForCloudEndpointsrestoreheartbeat(ctx context.Context, id CloudEndpointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreheartbeat", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCloudEndpointsrestoreheartbeat handles the response to the CloudEndpointsrestoreheartbeat request. The method always
// closes the http.Response Body.
func (c BackupRestoreClient) responderForCloudEndpointsrestoreheartbeat(resp *http.Response) (result CloudEndpointsrestoreheartbeatOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
