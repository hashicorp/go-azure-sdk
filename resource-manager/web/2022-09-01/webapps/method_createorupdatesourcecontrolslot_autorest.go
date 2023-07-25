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

type CreateOrUpdateSourceControlSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateSourceControlSlot ...
func (c WebAppsClient) CreateOrUpdateSourceControlSlot(ctx context.Context, id SlotId, input SiteSourceControl) (result CreateOrUpdateSourceControlSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateSourceControlSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSourceControlSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateSourceControlSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSourceControlSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateSourceControlSlotThenPoll performs CreateOrUpdateSourceControlSlot then polls until it's completed
func (c WebAppsClient) CreateOrUpdateSourceControlSlotThenPoll(ctx context.Context, id SlotId, input SiteSourceControl) error {
	result, err := c.CreateOrUpdateSourceControlSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateSourceControlSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateSourceControlSlot: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateSourceControlSlot prepares the CreateOrUpdateSourceControlSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateSourceControlSlot(ctx context.Context, id SlotId, input SiteSourceControl) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdateSourceControlSlot sends the CreateOrUpdateSourceControlSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateOrUpdateSourceControlSlot(ctx context.Context, req *http.Request) (future CreateOrUpdateSourceControlSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
