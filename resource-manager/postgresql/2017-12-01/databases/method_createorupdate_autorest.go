package databases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type CreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Poller       *resourcemanager.LongRunningPoller
}

// CreateOrUpdate ...
func (c DatabasesClient) CreateOrUpdate(ctx context.Context, id DatabaseId, input Database) (result CreateOrUpdateOperationResponse, err error) {
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
func (c DatabasesClient) CreateOrUpdateThenPoll(ctx context.Context, id DatabaseId, input Database) error {
	result, err := c.CreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdate: %+v", err)
	}

	return nil
}
