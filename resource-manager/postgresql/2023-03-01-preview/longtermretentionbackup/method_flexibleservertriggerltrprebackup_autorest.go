package longtermretentionbackup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlexibleServerTriggerLtrPreBackupOperationResponse struct {
	HttpResponse *http.Response
	Model        *LtrPreBackupResponse
}

// FlexibleServerTriggerLtrPreBackup ...
func (c LongTermRetentionBackupClient) FlexibleServerTriggerLtrPreBackup(ctx context.Context, id FlexibleServerId, input BackupRequestBase) (result FlexibleServerTriggerLtrPreBackupOperationResponse, err error) {
	req, err := c.preparerForFlexibleServerTriggerLtrPreBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "FlexibleServerTriggerLtrPreBackup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "FlexibleServerTriggerLtrPreBackup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFlexibleServerTriggerLtrPreBackup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "FlexibleServerTriggerLtrPreBackup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFlexibleServerTriggerLtrPreBackup prepares the FlexibleServerTriggerLtrPreBackup request.
func (c LongTermRetentionBackupClient) preparerForFlexibleServerTriggerLtrPreBackup(ctx context.Context, id FlexibleServerId, input BackupRequestBase) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/ltrPreBackup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFlexibleServerTriggerLtrPreBackup handles the response to the FlexibleServerTriggerLtrPreBackup request. The method always
// closes the http.Response Body.
func (c LongTermRetentionBackupClient) responderForFlexibleServerTriggerLtrPreBackup(resp *http.Response) (result FlexibleServerTriggerLtrPreBackupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
