package dataprotections

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

type BackupInstancesValidateCrossRegionRestoreOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *OperationJobExtendedInfo
}

// BackupInstancesValidateCrossRegionRestore ...
func (c DataprotectionsClient) BackupInstancesValidateCrossRegionRestore(ctx context.Context, id ProviderLocationId, input ValidateCrossRegionRestoreRequestObject) (result BackupInstancesValidateCrossRegionRestoreOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/validateCrossRegionRestore", id.ID()),
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

// BackupInstancesValidateCrossRegionRestoreThenPoll performs BackupInstancesValidateCrossRegionRestore then polls until it's completed
func (c DataprotectionsClient) BackupInstancesValidateCrossRegionRestoreThenPoll(ctx context.Context, id ProviderLocationId, input ValidateCrossRegionRestoreRequestObject) error {
	result, err := c.BackupInstancesValidateCrossRegionRestore(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BackupInstancesValidateCrossRegionRestore: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BackupInstancesValidateCrossRegionRestore: %+v", err)
	}

	return nil
}
