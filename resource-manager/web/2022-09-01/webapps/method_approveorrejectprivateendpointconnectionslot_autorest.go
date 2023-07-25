package webapps

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

type ApproveOrRejectPrivateEndpointConnectionSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApproveOrRejectPrivateEndpointConnectionSlot ...
func (c WebAppsClient) ApproveOrRejectPrivateEndpointConnectionSlot(ctx context.Context, id SlotPrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) (result ApproveOrRejectPrivateEndpointConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForApproveOrRejectPrivateEndpointConnectionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApproveOrRejectPrivateEndpointConnectionSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApproveOrRejectPrivateEndpointConnectionSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApproveOrRejectPrivateEndpointConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApproveOrRejectPrivateEndpointConnectionSlotThenPoll performs ApproveOrRejectPrivateEndpointConnectionSlot then polls until it's completed
func (c WebAppsClient) ApproveOrRejectPrivateEndpointConnectionSlotThenPoll(ctx context.Context, id SlotPrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) error {
	result, err := c.ApproveOrRejectPrivateEndpointConnectionSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ApproveOrRejectPrivateEndpointConnectionSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApproveOrRejectPrivateEndpointConnectionSlot: %+v", err)
	}

	return nil
}

// preparerForApproveOrRejectPrivateEndpointConnectionSlot prepares the ApproveOrRejectPrivateEndpointConnectionSlot request.
func (c WebAppsClient) preparerForApproveOrRejectPrivateEndpointConnectionSlot(ctx context.Context, id SlotPrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForApproveOrRejectPrivateEndpointConnectionSlot sends the ApproveOrRejectPrivateEndpointConnectionSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForApproveOrRejectPrivateEndpointConnectionSlot(ctx context.Context, req *http.Request) (future ApproveOrRejectPrivateEndpointConnectionSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
