package global

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeletedWebAppSnapshotsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Snapshot
}

// GetDeletedWebAppSnapshots ...
func (c GlobalClient) GetDeletedWebAppSnapshots(ctx context.Context, id DeletedSiteId) (result GetDeletedWebAppSnapshotsOperationResponse, err error) {
	req, err := c.preparerForGetDeletedWebAppSnapshots(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebAppSnapshots", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebAppSnapshots", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeletedWebAppSnapshots(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebAppSnapshots", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeletedWebAppSnapshots prepares the GetDeletedWebAppSnapshots request.
func (c GlobalClient) preparerForGetDeletedWebAppSnapshots(ctx context.Context, id DeletedSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/snapshots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDeletedWebAppSnapshots handles the response to the GetDeletedWebAppSnapshots request. The method always
// closes the http.Response Body.
func (c GlobalClient) responderForGetDeletedWebAppSnapshots(resp *http.Response) (result GetDeletedWebAppSnapshotsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
