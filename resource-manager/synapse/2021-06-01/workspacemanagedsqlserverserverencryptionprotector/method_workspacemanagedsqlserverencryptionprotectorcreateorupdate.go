package workspacemanagedsqlserverserverencryptionprotector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *EncryptionProtector
}

// WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx context.Context, id WorkspaceId, input EncryptionProtector) (result WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/encryptionProtector/current", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input EncryptionProtector) error {
	return c.WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateCallbackThenPoll(ctx, id, input, nil)
}

// WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateCallbackThenPoll performs WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate, runs the optional callback function, then polls until it's completed
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateCallbackThenPoll(ctx context.Context, id WorkspaceId, input EncryptionProtector, callback func() error) error {
	result, err := c.WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate: %+v", err)
	}

	return nil
}
