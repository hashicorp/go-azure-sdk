package workspacemanagedsqlserverblobauditing

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

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ExtendedServerBlobAuditingPolicy
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, id WorkspaceId, input ExtendedServerBlobAuditingPolicy) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/extendedAuditingSettings/default", id.ID()),
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

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input ExtendedServerBlobAuditingPolicy) error {
	result, err := c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	return nil
}
