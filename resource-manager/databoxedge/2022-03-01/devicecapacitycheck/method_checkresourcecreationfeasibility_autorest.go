package devicecapacitycheck

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckResourceCreationFeasibilityOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type CheckResourceCreationFeasibilityOperationOptions struct {
	CapacityName *string
}

func DefaultCheckResourceCreationFeasibilityOperationOptions() CheckResourceCreationFeasibilityOperationOptions {
	return CheckResourceCreationFeasibilityOperationOptions{}
}

func (o CheckResourceCreationFeasibilityOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o CheckResourceCreationFeasibilityOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.CapacityName != nil {
		out["capacityName"] = *o.CapacityName
	}

	return out
}

// CheckResourceCreationFeasibility ...
func (c DeviceCapacityCheckClient) CheckResourceCreationFeasibility(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options CheckResourceCreationFeasibilityOperationOptions) (result CheckResourceCreationFeasibilityOperationResponse, err error) {
	req, err := c.preparerForCheckResourceCreationFeasibility(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devicecapacitycheck.DeviceCapacityCheckClient", "CheckResourceCreationFeasibility", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCheckResourceCreationFeasibility(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devicecapacitycheck.DeviceCapacityCheckClient", "CheckResourceCreationFeasibility", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CheckResourceCreationFeasibilityThenPoll performs CheckResourceCreationFeasibility then polls until it's completed
func (c DeviceCapacityCheckClient) CheckResourceCreationFeasibilityThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options CheckResourceCreationFeasibilityOperationOptions) error {
	result, err := c.CheckResourceCreationFeasibility(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing CheckResourceCreationFeasibility: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CheckResourceCreationFeasibility: %+v", err)
	}

	return nil
}

// preparerForCheckResourceCreationFeasibility prepares the CheckResourceCreationFeasibility request.
func (c DeviceCapacityCheckClient) preparerForCheckResourceCreationFeasibility(ctx context.Context, id DataBoxEdgeDeviceId, input DeviceCapacityRequestInfo, options CheckResourceCreationFeasibilityOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/deviceCapacityCheck", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCheckResourceCreationFeasibility sends the CheckResourceCreationFeasibility request. The method will close the
// http.Response Body if it receives an error.
func (c DeviceCapacityCheckClient) senderForCheckResourceCreationFeasibility(ctx context.Context, req *http.Request) (future CheckResourceCreationFeasibilityOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
