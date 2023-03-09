package replicationvaulthealth

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

type RefreshOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Refresh ...
func (c ReplicationVaultHealthClient) Refresh(ctx context.Context, id VaultId) (result RefreshOperationResponse, err error) {
	req, err := c.preparerForRefresh(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationvaulthealth.ReplicationVaultHealthClient", "Refresh", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRefresh(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationvaulthealth.ReplicationVaultHealthClient", "Refresh", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RefreshThenPoll performs Refresh then polls until it's completed
func (c ReplicationVaultHealthClient) RefreshThenPoll(ctx context.Context, id VaultId) error {
	result, err := c.Refresh(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Refresh: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Refresh: %+v", err)
	}

	return nil
}

// preparerForRefresh prepares the Refresh request.
func (c ReplicationVaultHealthClient) preparerForRefresh(ctx context.Context, id VaultId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/replicationVaultHealth/default/refresh", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRefresh sends the Refresh request. The method will close the
// http.Response Body if it receives an error.
func (c ReplicationVaultHealthClient) senderForRefresh(ctx context.Context, req *http.Request) (future RefreshOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
