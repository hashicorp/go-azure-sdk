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

type GetFunctionsAdminTokenSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *string
}

// GetFunctionsAdminTokenSlot ...
func (c WebAppsClient) GetFunctionsAdminTokenSlot(ctx context.Context, id SlotId) (result GetFunctionsAdminTokenSlotOperationResponse, err error) {
	req, err := c.preparerForGetFunctionsAdminTokenSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminTokenSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminTokenSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFunctionsAdminTokenSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminTokenSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFunctionsAdminTokenSlot prepares the GetFunctionsAdminTokenSlot request.
func (c WebAppsClient) preparerForGetFunctionsAdminTokenSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/functions/admin/token", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFunctionsAdminTokenSlot handles the response to the GetFunctionsAdminTokenSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetFunctionsAdminTokenSlot(resp *http.Response) (result GetFunctionsAdminTokenSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
