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

type GetWebSiteContainerLogsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetWebSiteContainerLogs ...
func (c WebAppsClient) GetWebSiteContainerLogs(ctx context.Context, id SiteId) (result GetWebSiteContainerLogsOperationResponse, err error) {
	req, err := c.preparerForGetWebSiteContainerLogs(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogs", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogs", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetWebSiteContainerLogs(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogs", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetWebSiteContainerLogs prepares the GetWebSiteContainerLogs request.
func (c WebAppsClient) preparerForGetWebSiteContainerLogs(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/containerlogs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetWebSiteContainerLogs handles the response to the GetWebSiteContainerLogs request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetWebSiteContainerLogs(resp *http.Response) (result GetWebSiteContainerLogsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
