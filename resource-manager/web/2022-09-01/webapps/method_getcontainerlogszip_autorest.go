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

type GetContainerLogsZipOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetContainerLogsZip ...
func (c WebAppsClient) GetContainerLogsZip(ctx context.Context, id SiteId) (result GetContainerLogsZipOperationResponse, err error) {
	req, err := c.preparerForGetContainerLogsZip(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZip", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZip", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetContainerLogsZip(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZip", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetContainerLogsZip prepares the GetContainerLogsZip request.
func (c WebAppsClient) preparerForGetContainerLogsZip(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/containerlogs/zip/download", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetContainerLogsZip handles the response to the GetContainerLogsZip request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetContainerLogsZip(resp *http.Response) (result GetContainerLogsZipOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
