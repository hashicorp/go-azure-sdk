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

type GetScmAllowedSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// GetScmAllowedSlot ...
func (c WebAppsClient) GetScmAllowedSlot(ctx context.Context, id SlotId) (result GetScmAllowedSlotOperationResponse, err error) {
	req, err := c.preparerForGetScmAllowedSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetScmAllowedSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetScmAllowedSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetScmAllowedSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetScmAllowedSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetScmAllowedSlot prepares the GetScmAllowedSlot request.
func (c WebAppsClient) preparerForGetScmAllowedSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/scm", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetScmAllowedSlot handles the response to the GetScmAllowedSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetScmAllowedSlot(resp *http.Response) (result GetScmAllowedSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
