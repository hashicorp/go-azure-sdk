package incrementalrestorepoints

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskRestorePointGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiskRestorePoint
}

// DiskRestorePointGet ...
func (c IncrementalRestorePointsClient) DiskRestorePointGet(ctx context.Context, id DiskRestorePointId) (result DiskRestorePointGetOperationResponse, err error) {
	req, err := c.preparerForDiskRestorePointGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDiskRestorePointGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDiskRestorePointGet prepares the DiskRestorePointGet request.
func (c IncrementalRestorePointsClient) preparerForDiskRestorePointGet(ctx context.Context, id DiskRestorePointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDiskRestorePointGet handles the response to the DiskRestorePointGet request. The method always
// closes the http.Response Body.
func (c IncrementalRestorePointsClient) responderForDiskRestorePointGet(resp *http.Response) (result DiskRestorePointGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
