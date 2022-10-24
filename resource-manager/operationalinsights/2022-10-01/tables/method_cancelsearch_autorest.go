package tables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelSearchOperationResponse struct {
	HttpResponse *http.Response
}

// CancelSearch ...
func (c TablesClient) CancelSearch(ctx context.Context, id TableId) (result CancelSearchOperationResponse, err error) {
	req, err := c.preparerForCancelSearch(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "CancelSearch", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "CancelSearch", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCancelSearch(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "CancelSearch", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCancelSearch prepares the CancelSearch request.
func (c TablesClient) preparerForCancelSearch(ctx context.Context, id TableId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancelSearch", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCancelSearch handles the response to the CancelSearch request. The method always
// closes the http.Response Body.
func (c TablesClient) responderForCancelSearch(resp *http.Response) (result CancelSearchOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
