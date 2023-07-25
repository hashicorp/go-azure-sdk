package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreFromBackupBlobOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreFromBackupBlob ...
func (c WebAppsClient) RestoreFromBackupBlob(ctx context.Context, id SiteId, input RestoreRequest) (result RestoreFromBackupBlobOperationResponse, err error) {
	req, err := c.preparerForRestoreFromBackupBlob(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromBackupBlob", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreFromBackupBlob(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromBackupBlob", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreFromBackupBlobThenPoll performs RestoreFromBackupBlob then polls until it's completed
func (c WebAppsClient) RestoreFromBackupBlobThenPoll(ctx context.Context, id SiteId, input RestoreRequest) error {
	result, err := c.RestoreFromBackupBlob(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFromBackupBlob: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreFromBackupBlob: %+v", err)
	}

	return nil
}

// preparerForRestoreFromBackupBlob prepares the RestoreFromBackupBlob request.
func (c WebAppsClient) preparerForRestoreFromBackupBlob(ctx context.Context, id SiteId, input RestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreFromBackupBlob", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreFromBackupBlob sends the RestoreFromBackupBlob request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreFromBackupBlob(ctx context.Context, req *http.Request) (future RestoreFromBackupBlobOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
