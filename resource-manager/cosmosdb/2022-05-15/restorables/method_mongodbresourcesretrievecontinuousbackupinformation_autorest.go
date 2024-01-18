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

type MongoDBResourcesRetrieveContinuousBackupInformationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *BackupInformation
}

// MongoDBResourcesRetrieveContinuousBackupInformation ...
func (c RestorablesClient) MongoDBResourcesRetrieveContinuousBackupInformation(ctx context.Context, id MongodbDatabaseCollectionId, input ContinuousBackupRestoreLocation) (result MongoDBResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	req, err := c.preparerForMongoDBResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "MongoDBResourcesRetrieveContinuousBackupInformation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMongoDBResourcesRetrieveContinuousBackupInformation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "MongoDBResourcesRetrieveContinuousBackupInformation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MongoDBResourcesRetrieveContinuousBackupInformationThenPoll performs MongoDBResourcesRetrieveContinuousBackupInformation then polls until it's completed
func (c RestorablesClient) MongoDBResourcesRetrieveContinuousBackupInformationThenPoll(ctx context.Context, id MongodbDatabaseCollectionId, input ContinuousBackupRestoreLocation) error {
	result, err := c.MongoDBResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MongoDBResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MongoDBResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	return nil
}

// preparerForMongoDBResourcesRetrieveContinuousBackupInformation prepares the MongoDBResourcesRetrieveContinuousBackupInformation request.
func (c RestorablesClient) preparerForMongoDBResourcesRetrieveContinuousBackupInformation(ctx context.Context, id MongodbDatabaseCollectionId, input ContinuousBackupRestoreLocation) (*http.Request, error) {
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

// senderForMongoDBResourcesRetrieveContinuousBackupInformation sends the MongoDBResourcesRetrieveContinuousBackupInformation request. The method will close the
// http.Response Body if it receives an error.
func (c RestorablesClient) senderForMongoDBResourcesRetrieveContinuousBackupInformation(ctx context.Context, req *http.Request) (future MongoDBResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
