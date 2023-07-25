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

type DeleteBackupConfigurationOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteBackupConfiguration ...
func (c WebAppsClient) DeleteBackupConfiguration(ctx context.Context, id SiteId) (result DeleteBackupConfigurationOperationResponse, err error) {
	req, err := c.preparerForDeleteBackupConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteBackupConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteBackupConfiguration prepares the DeleteBackupConfiguration request.
func (c WebAppsClient) preparerForDeleteBackupConfiguration(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/backup", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteBackupConfiguration handles the response to the DeleteBackupConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteBackupConfiguration(resp *http.Response) (result DeleteBackupConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
