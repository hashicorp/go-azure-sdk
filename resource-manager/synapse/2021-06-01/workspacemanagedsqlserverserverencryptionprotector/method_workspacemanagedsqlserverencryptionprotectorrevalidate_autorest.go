package workspacemanagedsqlserverserverencryptionprotector

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

type WorkspaceManagedSqlServerEncryptionProtectorRevalidateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerEncryptionProtectorRevalidate ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerEncryptionProtectorRevalidateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorRevalidate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorRevalidate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerEncryptionProtectorRevalidateThenPoll performs WorkspaceManagedSqlServerEncryptionProtectorRevalidate then polls until it's completed
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorRevalidateThenPoll(ctx context.Context, id WorkspaceId) error {
	result, err := c.WorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerEncryptionProtectorRevalidate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerEncryptionProtectorRevalidate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerEncryptionProtectorRevalidate prepares the WorkspaceManagedSqlServerEncryptionProtectorRevalidate request.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) preparerForWorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/encryptionProtector/current/revalidate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedSqlServerEncryptionProtectorRevalidate sends the WorkspaceManagedSqlServerEncryptionProtectorRevalidate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) senderForWorkspaceManagedSqlServerEncryptionProtectorRevalidate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerEncryptionProtectorRevalidateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
