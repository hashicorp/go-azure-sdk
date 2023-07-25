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

type GetFtpAllowedSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// GetFtpAllowedSlot ...
func (c WebAppsClient) GetFtpAllowedSlot(ctx context.Context, id SlotId) (result GetFtpAllowedSlotOperationResponse, err error) {
	req, err := c.preparerForGetFtpAllowedSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowedSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowedSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFtpAllowedSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowedSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFtpAllowedSlot prepares the GetFtpAllowedSlot request.
func (c WebAppsClient) preparerForGetFtpAllowedSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/ftp", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFtpAllowedSlot handles the response to the GetFtpAllowedSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetFtpAllowedSlot(resp *http.Response) (result GetFtpAllowedSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
