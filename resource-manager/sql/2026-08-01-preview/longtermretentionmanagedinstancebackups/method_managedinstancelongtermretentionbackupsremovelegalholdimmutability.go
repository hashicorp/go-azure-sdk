package longtermretentionmanagedinstancebackups

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

type ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ManagedInstanceLongTermRetentionBackup
}

// ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability ...
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability(ctx context.Context, id LongTermRetentionManagedInstanceBackupId) (result ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/removeLegalHoldImmutability", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

// ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityThenPoll performs ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability then polls until it's completed
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityThenPoll(ctx context.Context, id LongTermRetentionManagedInstanceBackupId) error {
	return c.ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityCallbackThenPoll(ctx, id, nil)
}

// ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityCallbackThenPoll performs ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability, runs the optional callback function, then polls until it's completed
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutabilityCallbackThenPoll(ctx context.Context, id LongTermRetentionManagedInstanceBackupId, callback func() error) error {
	result, err := c.ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ManagedInstanceLongTermRetentionBackupsRemoveLegalHoldImmutability: %+v", err)
	}

	return nil
}
