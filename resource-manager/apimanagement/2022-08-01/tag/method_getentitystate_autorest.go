package tag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntityStateOperationResponse struct {
	HttpResponse *http.Response
}

// GetEntityState ...
func (c TagClient) GetEntityState(ctx context.Context, id TagId) (result GetEntityStateOperationResponse, err error) {
	req, err := c.preparerForGetEntityState(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tag.TagClient", "GetEntityState", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tag.TagClient", "GetEntityState", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetEntityState(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tag.TagClient", "GetEntityState", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetEntityState prepares the GetEntityState request.
func (c TagClient) preparerForGetEntityState(ctx context.Context, id TagId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetEntityState handles the response to the GetEntityState request. The method always
// closes the http.Response Body.
func (c TagClient) responderForGetEntityState(resp *http.Response) (result GetEntityStateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
