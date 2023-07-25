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

type ListPublishingCredentialsSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ListPublishingCredentialsSlot ...
func (c WebAppsClient) ListPublishingCredentialsSlot(ctx context.Context, id SlotId) (result ListPublishingCredentialsSlotOperationResponse, err error) {
	req, err := c.preparerForListPublishingCredentialsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingCredentialsSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForListPublishingCredentialsSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingCredentialsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ListPublishingCredentialsSlotThenPoll performs ListPublishingCredentialsSlot then polls until it's completed
func (c WebAppsClient) ListPublishingCredentialsSlotThenPoll(ctx context.Context, id SlotId) error {
	result, err := c.ListPublishingCredentialsSlot(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ListPublishingCredentialsSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ListPublishingCredentialsSlot: %+v", err)
	}

	return nil
}

// preparerForListPublishingCredentialsSlot prepares the ListPublishingCredentialsSlot request.
func (c WebAppsClient) preparerForListPublishingCredentialsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/publishingcredentials/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForListPublishingCredentialsSlot sends the ListPublishingCredentialsSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForListPublishingCredentialsSlot(ctx context.Context, req *http.Request) (future ListPublishingCredentialsSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
