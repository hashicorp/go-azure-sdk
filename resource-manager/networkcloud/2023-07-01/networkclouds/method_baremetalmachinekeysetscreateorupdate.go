package networkclouds

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

type BareMetalMachineKeySetsCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *BareMetalMachineKeySet
}

// BareMetalMachineKeySetsCreateOrUpdate ...
func (c NetworkcloudsClient) BareMetalMachineKeySetsCreateOrUpdate(ctx context.Context, id BareMetalMachineKeySetId, input BareMetalMachineKeySet) (result BareMetalMachineKeySetsCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// BareMetalMachineKeySetsCreateOrUpdateThenPoll performs BareMetalMachineKeySetsCreateOrUpdate then polls until it's completed
func (c NetworkcloudsClient) BareMetalMachineKeySetsCreateOrUpdateThenPoll(ctx context.Context, id BareMetalMachineKeySetId, input BareMetalMachineKeySet) error {
	result, err := c.BareMetalMachineKeySetsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BareMetalMachineKeySetsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BareMetalMachineKeySetsCreateOrUpdate: %+v", err)
	}

	return nil
}
