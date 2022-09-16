package workspacemanagedsqlserverserverencryptionprotector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerEncryptionProtectorGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EncryptionProtector
}

// WorkspaceManagedSqlServerEncryptionProtectorGet ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorGet(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerEncryptionProtectorGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerEncryptionProtectorGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerEncryptionProtectorGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerEncryptionProtectorGet prepares the WorkspaceManagedSqlServerEncryptionProtectorGet request.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) preparerForWorkspaceManagedSqlServerEncryptionProtectorGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/encryptionProtector/current", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceManagedSqlServerEncryptionProtectorGet handles the response to the WorkspaceManagedSqlServerEncryptionProtectorGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) responderForWorkspaceManagedSqlServerEncryptionProtectorGet(resp *http.Response) (result WorkspaceManagedSqlServerEncryptionProtectorGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
