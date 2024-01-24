package billingsubscriptions

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

type BillingSubscriptionsMoveOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *BillingSubscription
}

// BillingSubscriptionsMove ...
func (c BillingSubscriptionsClient) BillingSubscriptionsMove(ctx context.Context, id BillingSubscriptionId, input MoveBillingSubscriptionRequest) (result BillingSubscriptionsMoveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/move", id.ID()),
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

// BillingSubscriptionsMoveThenPoll performs BillingSubscriptionsMove then polls until it's completed
func (c BillingSubscriptionsClient) BillingSubscriptionsMoveThenPoll(ctx context.Context, id BillingSubscriptionId, input MoveBillingSubscriptionRequest) error {
	result, err := c.BillingSubscriptionsMove(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BillingSubscriptionsMove: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BillingSubscriptionsMove: %+v", err)
	}

	return nil
}
