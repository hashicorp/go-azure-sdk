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

type CreateInstanceFunctionSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateInstanceFunctionSlot ...
func (c WebAppsClient) CreateInstanceFunctionSlot(ctx context.Context, id SlotFunctionId, input FunctionEnvelope) (result CreateInstanceFunctionSlotOperationResponse, err error) {
	req, err := c.preparerForCreateInstanceFunctionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceFunctionSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateInstanceFunctionSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceFunctionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateInstanceFunctionSlotThenPoll performs CreateInstanceFunctionSlot then polls until it's completed
func (c WebAppsClient) CreateInstanceFunctionSlotThenPoll(ctx context.Context, id SlotFunctionId, input FunctionEnvelope) error {
	result, err := c.CreateInstanceFunctionSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateInstanceFunctionSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateInstanceFunctionSlot: %+v", err)
	}

	return nil
}

// preparerForCreateInstanceFunctionSlot prepares the CreateInstanceFunctionSlot request.
func (c WebAppsClient) preparerForCreateInstanceFunctionSlot(ctx context.Context, id SlotFunctionId, input FunctionEnvelope) (*http.Request, error) {
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

// senderForCreateInstanceFunctionSlot sends the CreateInstanceFunctionSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateInstanceFunctionSlot(ctx context.Context, req *http.Request) (future CreateInstanceFunctionSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
