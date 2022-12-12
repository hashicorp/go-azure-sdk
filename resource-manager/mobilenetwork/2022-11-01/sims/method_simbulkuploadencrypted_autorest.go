package sims

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

type SimBulkUploadEncryptedOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SimBulkUploadEncrypted ...
func (c SIMsClient) SimBulkUploadEncrypted(ctx context.Context, id SimGroupId, input EncryptedSimUploadList) (result SimBulkUploadEncryptedOperationResponse, err error) {
	req, err := c.preparerForSimBulkUploadEncrypted(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "SimBulkUploadEncrypted", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSimBulkUploadEncrypted(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "SimBulkUploadEncrypted", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SimBulkUploadEncryptedThenPoll performs SimBulkUploadEncrypted then polls until it's completed
func (c SIMsClient) SimBulkUploadEncryptedThenPoll(ctx context.Context, id SimGroupId, input EncryptedSimUploadList) error {
	result, err := c.SimBulkUploadEncrypted(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SimBulkUploadEncrypted: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SimBulkUploadEncrypted: %+v", err)
	}

	return nil
}

// preparerForSimBulkUploadEncrypted prepares the SimBulkUploadEncrypted request.
func (c SIMsClient) preparerForSimBulkUploadEncrypted(ctx context.Context, id SimGroupId, input EncryptedSimUploadList) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/uploadEncryptedSims", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSimBulkUploadEncrypted sends the SimBulkUploadEncrypted request. The method will close the
// http.Response Body if it receives an error.
func (c SIMsClient) senderForSimBulkUploadEncrypted(ctx context.Context, req *http.Request) (future SimBulkUploadEncryptedOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
