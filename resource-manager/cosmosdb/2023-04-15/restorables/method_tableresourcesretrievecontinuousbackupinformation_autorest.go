package restorables

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

type TableResourcesRetrieveContinuousBackupInformationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *BackupInformation
}

// TableResourcesRetrieveContinuousBackupInformation ...
func (c RestorablesClient) TableResourcesRetrieveContinuousBackupInformation(ctx context.Context, id TableId, input ContinuousBackupRestoreLocation) (result TableResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	req, err := c.preparerForTableResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "TableResourcesRetrieveContinuousBackupInformation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTableResourcesRetrieveContinuousBackupInformation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "TableResourcesRetrieveContinuousBackupInformation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TableResourcesRetrieveContinuousBackupInformationThenPoll performs TableResourcesRetrieveContinuousBackupInformation then polls until it's completed
func (c RestorablesClient) TableResourcesRetrieveContinuousBackupInformationThenPoll(ctx context.Context, id TableId, input ContinuousBackupRestoreLocation) error {
	result, err := c.TableResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing TableResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after TableResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	return nil
}

// preparerForTableResourcesRetrieveContinuousBackupInformation prepares the TableResourcesRetrieveContinuousBackupInformation request.
func (c RestorablesClient) preparerForTableResourcesRetrieveContinuousBackupInformation(ctx context.Context, id TableId, input ContinuousBackupRestoreLocation) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/retrieveContinuousBackupInformation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTableResourcesRetrieveContinuousBackupInformation sends the TableResourcesRetrieveContinuousBackupInformation request. The method will close the
// http.Response Body if it receives an error.
func (c RestorablesClient) senderForTableResourcesRetrieveContinuousBackupInformation(ctx context.Context, req *http.Request) (future TableResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
