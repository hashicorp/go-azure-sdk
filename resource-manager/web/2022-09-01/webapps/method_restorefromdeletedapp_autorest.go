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

type RestoreFromDeletedAppOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreFromDeletedApp ...
func (c WebAppsClient) RestoreFromDeletedApp(ctx context.Context, id SiteId, input DeletedAppRestoreRequest) (result RestoreFromDeletedAppOperationResponse, err error) {
	req, err := c.preparerForRestoreFromDeletedApp(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromDeletedApp", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreFromDeletedApp(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromDeletedApp", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreFromDeletedAppThenPoll performs RestoreFromDeletedApp then polls until it's completed
func (c WebAppsClient) RestoreFromDeletedAppThenPoll(ctx context.Context, id SiteId, input DeletedAppRestoreRequest) error {
	result, err := c.RestoreFromDeletedApp(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFromDeletedApp: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreFromDeletedApp: %+v", err)
	}

	return nil
}

// preparerForRestoreFromDeletedApp prepares the RestoreFromDeletedApp request.
func (c WebAppsClient) preparerForRestoreFromDeletedApp(ctx context.Context, id SiteId, input DeletedAppRestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreFromDeletedApp", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreFromDeletedApp sends the RestoreFromDeletedApp request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreFromDeletedApp(ctx context.Context, req *http.Request) (future RestoreFromDeletedAppOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
