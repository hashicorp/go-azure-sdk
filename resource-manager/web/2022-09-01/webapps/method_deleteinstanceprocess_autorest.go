package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteInstanceProcessOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteInstanceProcess ...
func (c WebAppsClient) DeleteInstanceProcess(ctx context.Context, id InstanceProcessId) (result DeleteInstanceProcessOperationResponse, err error) {
	req, err := c.preparerForDeleteInstanceProcess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteInstanceProcess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteInstanceProcess prepares the DeleteInstanceProcess request.
func (c WebAppsClient) preparerForDeleteInstanceProcess(ctx context.Context, id InstanceProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteInstanceProcess handles the response to the DeleteInstanceProcess request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteInstanceProcess(resp *http.Response) (result DeleteInstanceProcessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
