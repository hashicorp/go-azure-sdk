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

type ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ManagedInstanceLongTermRetentionBackup
}

// ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability ...
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability(ctx context.Context, id LongTermRetentionManagedInstanceBackupId) (result ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/setLegalHoldImmutability", id.ID()),
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

// ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityThenPoll performs ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability then polls until it's completed
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityThenPoll(ctx context.Context, id LongTermRetentionManagedInstanceBackupId) error {
	return c.ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityCallbackThenPoll(ctx, id, nil)
}

// ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityCallbackThenPoll performs ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability, runs the optional callback function, then polls until it's completed
func (c LongTermRetentionManagedInstanceBackupsClient) ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutabilityCallbackThenPoll(ctx context.Context, id LongTermRetentionManagedInstanceBackupId, callback func() error) error {
	result, err := c.ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ManagedInstanceLongTermRetentionBackupsSetLegalHoldImmutability: %+v", err)
	}

	return nil
}
