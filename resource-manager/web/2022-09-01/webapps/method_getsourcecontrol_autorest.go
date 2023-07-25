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

type GetSourceControlOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteSourceControl
}

// GetSourceControl ...
func (c WebAppsClient) GetSourceControl(ctx context.Context, id SiteId) (result GetSourceControlOperationResponse, err error) {
	req, err := c.preparerForGetSourceControl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSourceControl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSourceControl prepares the GetSourceControl request.
func (c WebAppsClient) preparerForGetSourceControl(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSourceControl handles the response to the GetSourceControl request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSourceControl(resp *http.Response) (result GetSourceControlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
