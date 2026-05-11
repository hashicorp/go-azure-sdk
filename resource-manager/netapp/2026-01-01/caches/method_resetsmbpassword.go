package caches

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

type ResetSmbPasswordOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Cache
}

// ResetSmbPassword ...
func (c CachesClient) ResetSmbPassword(ctx context.Context, id CacheId) (result ResetSmbPasswordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/resetSmbPassword", id.ID()),
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

// ResetSmbPasswordThenPoll performs ResetSmbPassword then polls until it's completed
func (c CachesClient) ResetSmbPasswordThenPoll(ctx context.Context, id CacheId) error {
	return c.ResetSmbPasswordCallbackThenPoll(ctx, id, nil)
}

// ResetSmbPasswordCallbackThenPoll performs ResetSmbPassword, runs the optional callback function, then polls until it's completed
func (c CachesClient) ResetSmbPasswordCallbackThenPoll(ctx context.Context, id CacheId, callback func() error) error {
	result, err := c.ResetSmbPassword(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ResetSmbPassword: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResetSmbPassword: %+v", err)
	}

	return nil
}
