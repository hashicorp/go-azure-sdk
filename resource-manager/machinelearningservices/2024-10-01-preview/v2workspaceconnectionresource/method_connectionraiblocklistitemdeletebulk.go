package v2workspaceconnectionresource

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

type ConnectionRaiBlocklistItemDeleteBulkOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ConnectionRaiBlocklistItemDeleteBulk ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemDeleteBulk(ctx context.Context, id RaiBlocklistId, input interface{}) (result ConnectionRaiBlocklistItemDeleteBulkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/deleteRaiBlocklistItems", id.ID()),
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

// ConnectionRaiBlocklistItemDeleteBulkThenPoll performs ConnectionRaiBlocklistItemDeleteBulk then polls until it's completed
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemDeleteBulkThenPoll(ctx context.Context, id RaiBlocklistId, input interface{}) error {
	result, err := c.ConnectionRaiBlocklistItemDeleteBulk(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConnectionRaiBlocklistItemDeleteBulk: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConnectionRaiBlocklistItemDeleteBulk: %+v", err)
	}

	return nil
}
