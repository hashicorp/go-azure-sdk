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

type CreateInstanceMSDeployOperationSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateInstanceMSDeployOperationSlot ...
func (c WebAppsClient) CreateInstanceMSDeployOperationSlot(ctx context.Context, id SlotInstanceId, input MSDeploy) (result CreateInstanceMSDeployOperationSlotOperationResponse, err error) {
	req, err := c.preparerForCreateInstanceMSDeployOperationSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceMSDeployOperationSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateInstanceMSDeployOperationSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceMSDeployOperationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateInstanceMSDeployOperationSlotThenPoll performs CreateInstanceMSDeployOperationSlot then polls until it's completed
func (c WebAppsClient) CreateInstanceMSDeployOperationSlotThenPoll(ctx context.Context, id SlotInstanceId, input MSDeploy) error {
	result, err := c.CreateInstanceMSDeployOperationSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateInstanceMSDeployOperationSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateInstanceMSDeployOperationSlot: %+v", err)
	}

	return nil
}

// preparerForCreateInstanceMSDeployOperationSlot prepares the CreateInstanceMSDeployOperationSlot request.
func (c WebAppsClient) preparerForCreateInstanceMSDeployOperationSlot(ctx context.Context, id SlotInstanceId, input MSDeploy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/mSDeploy", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateInstanceMSDeployOperationSlot sends the CreateInstanceMSDeployOperationSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateInstanceMSDeployOperationSlot(ctx context.Context, req *http.Request) (future CreateInstanceMSDeployOperationSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
