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

type CreateOrUpdateSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateSlot ...
func (c WebAppsClient) CreateOrUpdateSlot(ctx context.Context, id SlotId, input Site) (result CreateOrUpdateSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateSlotThenPoll performs CreateOrUpdateSlot then polls until it's completed
func (c WebAppsClient) CreateOrUpdateSlotThenPoll(ctx context.Context, id SlotId, input Site) error {
	result, err := c.CreateOrUpdateSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateSlot: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateSlot prepares the CreateOrUpdateSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateSlot(ctx context.Context, id SlotId, input Site) (*http.Request, error) {
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

// senderForCreateOrUpdateSlot sends the CreateOrUpdateSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateOrUpdateSlot(ctx context.Context, req *http.Request) (future CreateOrUpdateSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
