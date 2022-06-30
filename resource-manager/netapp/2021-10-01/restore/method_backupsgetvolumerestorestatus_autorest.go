package restore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupsGetVolumeRestoreStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestoreStatus
}

// BackupsGetVolumeRestoreStatus ...
func (c RestoreClient) BackupsGetVolumeRestoreStatus(ctx context.Context, id VolumeId) (result BackupsGetVolumeRestoreStatusOperationResponse, err error) {
	req, err := c.preparerForBackupsGetVolumeRestoreStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restore.RestoreClient", "BackupsGetVolumeRestoreStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restore.RestoreClient", "BackupsGetVolumeRestoreStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupsGetVolumeRestoreStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restore.RestoreClient", "BackupsGetVolumeRestoreStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupsGetVolumeRestoreStatus prepares the BackupsGetVolumeRestoreStatus request.
func (c RestoreClient) preparerForBackupsGetVolumeRestoreStatus(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreStatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupsGetVolumeRestoreStatus handles the response to the BackupsGetVolumeRestoreStatus request. The method always
// closes the http.Response Body.
func (c RestoreClient) responderForBackupsGetVolumeRestoreStatus(resp *http.Response) (result BackupsGetVolumeRestoreStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
