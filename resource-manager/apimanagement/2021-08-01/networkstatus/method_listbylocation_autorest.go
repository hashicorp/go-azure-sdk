package networkstatus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByLocationOperationResponse struct {
	HttpResponse *http.Response
	Model        *NetworkStatusContract
}

// ListByLocation ...
func (c NetworkStatusClient) ListByLocation(ctx context.Context, id LocationId) (result ListByLocationOperationResponse, err error) {
	req, err := c.preparerForListByLocation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "networkstatus.NetworkStatusClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "networkstatus.NetworkStatusClient", "ListByLocation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByLocation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "networkstatus.NetworkStatusClient", "ListByLocation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByLocation prepares the ListByLocation request.
func (c NetworkStatusClient) preparerForListByLocation(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkstatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByLocation handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (c NetworkStatusClient) responderForListByLocation(resp *http.Response) (result ListByLocationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
