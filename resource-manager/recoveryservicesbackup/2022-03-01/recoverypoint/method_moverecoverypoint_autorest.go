package recoverypoint

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

type MoveRecoveryPointOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MoveRecoveryPoint ...
func (c RecoveryPointClient) MoveRecoveryPoint(ctx context.Context, id RecoveryPointId, input MoveRPAcrossTiersRequest) (result MoveRecoveryPointOperationResponse, err error) {
	req, err := c.preparerForMoveRecoveryPoint(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypoint.RecoveryPointClient", "MoveRecoveryPoint", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMoveRecoveryPoint(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypoint.RecoveryPointClient", "MoveRecoveryPoint", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MoveRecoveryPointThenPoll performs MoveRecoveryPoint then polls until it's completed
func (c RecoveryPointClient) MoveRecoveryPointThenPoll(ctx context.Context, id RecoveryPointId, input MoveRPAcrossTiersRequest) error {
	result, err := c.MoveRecoveryPoint(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MoveRecoveryPoint: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MoveRecoveryPoint: %+v", err)
	}

	return nil
}

// preparerForMoveRecoveryPoint prepares the MoveRecoveryPoint request.
func (c RecoveryPointClient) preparerForMoveRecoveryPoint(ctx context.Context, id RecoveryPointId, input MoveRPAcrossTiersRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/move", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMoveRecoveryPoint sends the MoveRecoveryPoint request. The method will close the
// http.Response Body if it receives an error.
func (c RecoveryPointClient) senderForMoveRecoveryPoint(ctx context.Context, req *http.Request) (future MoveRecoveryPointOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
