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

type DeleteSourceControlOperationResponse struct {
	HttpResponse *http.Response
}

type DeleteSourceControlOperationOptions struct {
	AdditionalFlags *string
}

func DefaultDeleteSourceControlOperationOptions() DeleteSourceControlOperationOptions {
	return DeleteSourceControlOperationOptions{}
}

func (o DeleteSourceControlOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeleteSourceControlOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.AdditionalFlags != nil {
		out["additionalFlags"] = *o.AdditionalFlags
	}

	return out
}

// DeleteSourceControl ...
func (c WebAppsClient) DeleteSourceControl(ctx context.Context, id SiteId, options DeleteSourceControlOperationOptions) (result DeleteSourceControlOperationResponse, err error) {
	req, err := c.preparerForDeleteSourceControl(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSourceControl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSourceControl prepares the DeleteSourceControl request.
func (c WebAppsClient) preparerForDeleteSourceControl(ctx context.Context, id SiteId, options DeleteSourceControlOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteSourceControl handles the response to the DeleteSourceControl request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSourceControl(resp *http.Response) (result DeleteSourceControlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
