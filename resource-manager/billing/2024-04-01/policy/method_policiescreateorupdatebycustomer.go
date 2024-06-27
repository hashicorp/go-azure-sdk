package policy

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

type PoliciesCreateOrUpdateByCustomerOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CustomerPolicy
}

// PoliciesCreateOrUpdateByCustomer ...
func (c PolicyClient) PoliciesCreateOrUpdateByCustomer(ctx context.Context, id BillingProfileCustomerId, input CustomerPolicy) (result PoliciesCreateOrUpdateByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/policies/default", id.ID()),
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

// PoliciesCreateOrUpdateByCustomerThenPoll performs PoliciesCreateOrUpdateByCustomer then polls until it's completed
func (c PolicyClient) PoliciesCreateOrUpdateByCustomerThenPoll(ctx context.Context, id BillingProfileCustomerId, input CustomerPolicy) error {
	result, err := c.PoliciesCreateOrUpdateByCustomer(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PoliciesCreateOrUpdateByCustomer: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PoliciesCreateOrUpdateByCustomer: %+v", err)
	}

	return nil
}
