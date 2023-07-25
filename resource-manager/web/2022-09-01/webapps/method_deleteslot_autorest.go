package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSlotOperationResponse struct {
	HttpResponse *http.Response
}

type DeleteSlotOperationOptions struct {
	DeleteEmptyServerFarm *bool
	DeleteMetrics         *bool
}

func DefaultDeleteSlotOperationOptions() DeleteSlotOperationOptions {
	return DeleteSlotOperationOptions{}
}

func (o DeleteSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeleteSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DeleteEmptyServerFarm != nil {
		out["deleteEmptyServerFarm"] = *o.DeleteEmptyServerFarm
	}

	if o.DeleteMetrics != nil {
		out["deleteMetrics"] = *o.DeleteMetrics
	}

	return out
}

// DeleteSlot ...
func (c WebAppsClient) DeleteSlot(ctx context.Context, id SlotId, options DeleteSlotOperationOptions) (result DeleteSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSlot prepares the DeleteSlot request.
func (c WebAppsClient) preparerForDeleteSlot(ctx context.Context, id SlotId, options DeleteSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteSlot handles the response to the DeleteSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSlot(resp *http.Response) (result DeleteSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
