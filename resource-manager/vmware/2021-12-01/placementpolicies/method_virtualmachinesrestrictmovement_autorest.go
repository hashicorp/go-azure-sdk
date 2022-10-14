package placementpolicies

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

type VirtualMachinesRestrictMovementOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VirtualMachinesRestrictMovement ...
func (c PlacementPoliciesClient) VirtualMachinesRestrictMovement(ctx context.Context, id VirtualMachineId, input VirtualMachineRestrictMovement) (result VirtualMachinesRestrictMovementOperationResponse, err error) {
	req, err := c.preparerForVirtualMachinesRestrictMovement(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "placementpolicies.PlacementPoliciesClient", "VirtualMachinesRestrictMovement", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVirtualMachinesRestrictMovement(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "placementpolicies.PlacementPoliciesClient", "VirtualMachinesRestrictMovement", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VirtualMachinesRestrictMovementThenPoll performs VirtualMachinesRestrictMovement then polls until it's completed
func (c PlacementPoliciesClient) VirtualMachinesRestrictMovementThenPoll(ctx context.Context, id VirtualMachineId, input VirtualMachineRestrictMovement) error {
	result, err := c.VirtualMachinesRestrictMovement(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing VirtualMachinesRestrictMovement: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VirtualMachinesRestrictMovement: %+v", err)
	}

	return nil
}

// preparerForVirtualMachinesRestrictMovement prepares the VirtualMachinesRestrictMovement request.
func (c PlacementPoliciesClient) preparerForVirtualMachinesRestrictMovement(ctx context.Context, id VirtualMachineId, input VirtualMachineRestrictMovement) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restrictMovement", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVirtualMachinesRestrictMovement sends the VirtualMachinesRestrictMovement request. The method will close the
// http.Response Body if it receives an error.
func (c PlacementPoliciesClient) senderForVirtualMachinesRestrictMovement(ctx context.Context, req *http.Request) (future VirtualMachinesRestrictMovementOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
