package resources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveResourcesOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MoveResources ...
func (c ResourcesClient) MoveResources(ctx context.Context, id commonids.ResourceGroupId, input ResourcesMoveInfo) (result MoveResourcesOperationResponse, err error) {
	req, err := c.preparerForMoveResources(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMoveResources(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MoveResourcesThenPoll performs MoveResources then polls until it's completed
func (c ResourcesClient) MoveResourcesThenPoll(ctx context.Context, id commonids.ResourceGroupId, input ResourcesMoveInfo) error {
	result, err := c.MoveResources(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MoveResources: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MoveResources: %+v", err)
	}

	return nil
}

// preparerForMoveResources prepares the MoveResources request.
func (c ResourcesClient) preparerForMoveResources(ctx context.Context, id commonids.ResourceGroupId, input ResourcesMoveInfo) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/moveResources", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMoveResources sends the MoveResources request. The method will close the
// http.Response Body if it receives an error.
func (c ResourcesClient) senderForMoveResources(ctx context.Context, req *http.Request) (future MoveResourcesOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
