package backups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupStatus
}

// GetStatus ...
func (c BackupsClient) GetStatus(ctx context.Context, id VolumeId) (result GetStatusOperationResponse, err error) {
	req, err := c.preparerForGetStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "GetStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "GetStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "GetStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetStatus prepares the GetStatus request.
func (c BackupsClient) preparerForGetStatus(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupStatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetStatus handles the response to the GetStatus request. The method always
// closes the http.Response Body.
func (c BackupsClient) responderForGetStatus(resp *http.Response) (result GetStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
