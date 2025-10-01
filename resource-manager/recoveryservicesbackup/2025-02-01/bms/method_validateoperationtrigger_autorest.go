package bms

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

type ValidateOperationTriggerOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateOperationTrigger ...
func (c BmsClient) ValidateOperationTrigger(ctx context.Context, id VaultId, input ValidateOperationRequestResource) (result ValidateOperationTriggerOperationResponse, err error) {
	req, err := c.preparerForValidateOperationTrigger(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ValidateOperationTrigger", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateOperationTrigger(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ValidateOperationTrigger", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateOperationTriggerThenPoll performs ValidateOperationTrigger then polls until it's completed
func (c BmsClient) ValidateOperationTriggerThenPoll(ctx context.Context, id VaultId, input ValidateOperationRequestResource) error {
	result, err := c.ValidateOperationTrigger(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateOperationTrigger: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateOperationTrigger: %+v", err)
	}

	return nil
}

// preparerForValidateOperationTrigger prepares the ValidateOperationTrigger request.
func (c BmsClient) preparerForValidateOperationTrigger(ctx context.Context, id VaultId, input ValidateOperationRequestResource) (*http.Request, error) {
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

// senderForValidateOperationTrigger sends the ValidateOperationTrigger request. The method will close the
// http.Response Body if it receives an error.
func (c BmsClient) senderForValidateOperationTrigger(ctx context.Context, req *http.Request) (future ValidateOperationTriggerOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
