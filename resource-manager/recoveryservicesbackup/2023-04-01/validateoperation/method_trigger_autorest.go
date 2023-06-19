package validateoperation

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

type TriggerOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Trigger ...
func (c ValidateOperationClient) Trigger(ctx context.Context, id VaultId, input ValidateOperationRequest) (result TriggerOperationResponse, err error) {
	req, err := c.preparerForTrigger(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validateoperation.ValidateOperationClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTrigger(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validateoperation.ValidateOperationClient", "Trigger", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TriggerThenPoll performs Trigger then polls until it's completed
func (c ValidateOperationClient) TriggerThenPoll(ctx context.Context, id VaultId, input ValidateOperationRequest) error {
	result, err := c.Trigger(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Trigger: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Trigger: %+v", err)
	}

	return nil
}

// preparerForTrigger prepares the Trigger request.
func (c ValidateOperationClient) preparerForTrigger(ctx context.Context, id VaultId, input ValidateOperationRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupTriggerValidateOperation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTrigger sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (c ValidateOperationClient) senderForTrigger(ctx context.Context, req *http.Request) (future TriggerOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
