package triggers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerUnsubscribeTriggerFromEventsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *TriggerSubscriptionOperationStatus
}

// TriggerUnsubscribeTriggerFromEvents ...
func (c TriggersClient) TriggerUnsubscribeTriggerFromEvents(ctx context.Context, id TriggerId) (result TriggerUnsubscribeTriggerFromEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/unsubscribeFromEvents", id.Path()),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// TriggerUnsubscribeTriggerFromEventsThenPoll performs TriggerUnsubscribeTriggerFromEvents then polls until it's completed
func (c TriggersClient) TriggerUnsubscribeTriggerFromEventsThenPoll(ctx context.Context, id TriggerId) error {
	result, err := c.TriggerUnsubscribeTriggerFromEvents(ctx, id)
	if err != nil {
		return fmt.Errorf("performing TriggerUnsubscribeTriggerFromEvents: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after TriggerUnsubscribeTriggerFromEvents: %+v", err)
	}

	return nil
}
