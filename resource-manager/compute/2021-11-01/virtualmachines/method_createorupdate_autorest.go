package virtualmachines

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/client/base"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type CreateOrUpdateOperationResponse struct {
	Poller       *resourcemanager.LongRunningPoller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateOrUpdate ...
func (c VirtualMachinesClient) CreateOrUpdate(ctx context.Context, id VirtualMachineId, input VirtualMachine) (result CreateOrUpdateOperationResponse, err error) {
	req, err := c.Client2.NewPutRequest(ctx, id.ID(), defaultApiVersion, odata.Query{}, input)
	if err != nil {
		return
	}

	var resp *base.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.NewPollerFromResponse(ctx, resp, c.Client2)
	if err != nil {
		return
	}

	return
}

// CreateOrUpdateThenPoll performs CreateOrUpdate then polls until it's completed
func (c VirtualMachinesClient) CreateOrUpdateThenPoll(ctx context.Context, id VirtualMachineId, input VirtualMachine) error {
	result, err := c.CreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdate: %+v", err)
	}

	return nil
}
