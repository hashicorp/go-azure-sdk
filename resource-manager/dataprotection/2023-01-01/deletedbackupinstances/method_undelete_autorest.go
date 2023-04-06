package deletedbackupinstances

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

type UndeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Undelete ...
func (c DeletedBackupInstancesClient) Undelete(ctx context.Context, id DeletedBackupInstanceId) (result UndeleteOperationResponse, err error) {
	req, err := c.preparerForUndelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedbackupinstances.DeletedBackupInstancesClient", "Undelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUndelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedbackupinstances.DeletedBackupInstancesClient", "Undelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UndeleteThenPoll performs Undelete then polls until it's completed
func (c DeletedBackupInstancesClient) UndeleteThenPoll(ctx context.Context, id DeletedBackupInstanceId) error {
	result, err := c.Undelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Undelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Undelete: %+v", err)
	}

	return nil
}

// preparerForUndelete prepares the Undelete request.
func (c DeletedBackupInstancesClient) preparerForUndelete(ctx context.Context, id DeletedBackupInstanceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/undelete", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUndelete sends the Undelete request. The method will close the
// http.Response Body if it receives an error.
func (c DeletedBackupInstancesClient) senderForUndelete(ctx context.Context, req *http.Request) (future UndeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
