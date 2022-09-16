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

type WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx context.Context, id WorkspaceId, input EncryptionProtector) (result WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input EncryptionProtector) error {
	result, err := c.WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate prepares the WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate request.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) preparerForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx context.Context, id WorkspaceId, input EncryptionProtector) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/encryptionProtector/current", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate sends the WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) senderForWorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
