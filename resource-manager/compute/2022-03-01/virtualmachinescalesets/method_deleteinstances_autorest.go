package virtualmachinescalesets

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

type DeleteInstancesOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type DeleteInstancesOperationOptions struct {
	ForceDeletion *bool
}

func DefaultDeleteInstancesOperationOptions() DeleteInstancesOperationOptions {
	return DeleteInstancesOperationOptions{}
}

func (o DeleteInstancesOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeleteInstancesOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ForceDeletion != nil {
		out["forceDeletion"] = *o.ForceDeletion
	}

	return out
}

// DeleteInstances ...
func (c VirtualMachineScaleSetsClient) DeleteInstances(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceRequiredIDs, options DeleteInstancesOperationOptions) (result DeleteInstancesOperationResponse, err error) {
	req, err := c.preparerForDeleteInstances(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "DeleteInstances", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteInstances(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "DeleteInstances", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteInstancesThenPoll performs DeleteInstances then polls until it's completed
func (c VirtualMachineScaleSetsClient) DeleteInstancesThenPoll(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceRequiredIDs, options DeleteInstancesOperationOptions) error {
	result, err := c.DeleteInstances(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing DeleteInstances: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteInstances: %+v", err)
	}

	return nil
}

// preparerForDeleteInstances prepares the DeleteInstances request.
func (c VirtualMachineScaleSetsClient) preparerForDeleteInstances(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceRequiredIDs, options DeleteInstancesOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/delete", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteInstances sends the DeleteInstances request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetsClient) senderForDeleteInstances(ctx context.Context, req *http.Request) (future DeleteInstancesOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
