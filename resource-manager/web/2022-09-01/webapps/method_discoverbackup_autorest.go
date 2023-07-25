package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoverBackupOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestoreRequest
}

// DiscoverBackup ...
func (c WebAppsClient) DiscoverBackup(ctx context.Context, id SiteId, input RestoreRequest) (result DiscoverBackupOperationResponse, err error) {
	req, err := c.preparerForDiscoverBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDiscoverBackup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDiscoverBackup prepares the DiscoverBackup request.
func (c WebAppsClient) preparerForDiscoverBackup(ctx context.Context, id SiteId, input RestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/discoverbackup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDiscoverBackup handles the response to the DiscoverBackup request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDiscoverBackup(resp *http.Response) (result DiscoverBackupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
