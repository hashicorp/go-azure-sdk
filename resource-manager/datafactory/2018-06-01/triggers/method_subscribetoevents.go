package triggers

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

type SubscribeToEventsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *TriggerSubscriptionOperationStatus
}

// SubscribeToEvents ...
func (c TriggersClient) SubscribeToEvents(ctx context.Context, id TriggerId) (result SubscribeToEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,

		Path: fmt.Sprintf("%s/subscribeToEvents", id.ID()),
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

// SubscribeToEventsThenPoll performs SubscribeToEvents then polls until it's completed
func (c TriggersClient) SubscribeToEventsThenPoll(ctx context.Context, id TriggerId) error {
	result, err := c.SubscribeToEvents(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SubscribeToEvents: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SubscribeToEvents: %+v", err)
	}

	return nil
}
