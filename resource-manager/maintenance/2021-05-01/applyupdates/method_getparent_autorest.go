package applyupdates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetParentOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplyUpdate
}

// GetParent ...
func (c ApplyUpdatesClient) GetParent(ctx context.Context, id ScopedApplyUpdateId) (result GetParentOperationResponse, err error) {
	req, err := c.preparerForGetParent(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "GetParent", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "GetParent", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetParent(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "GetParent", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetParent prepares the GetParent request.
func (c ApplyUpdatesClient) preparerForGetParent(ctx context.Context, id ScopedApplyUpdateId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetParent handles the response to the GetParent request. The method always
// closes the http.Response Body.
func (c ApplyUpdatesClient) responderForGetParent(resp *http.Response) (result GetParentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
