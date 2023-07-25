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

type UpdateFtpAllowedSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// UpdateFtpAllowedSlot ...
func (c WebAppsClient) UpdateFtpAllowedSlot(ctx context.Context, id SlotId, input CsmPublishingCredentialsPoliciesEntity) (result UpdateFtpAllowedSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateFtpAllowedSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowedSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowedSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateFtpAllowedSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowedSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateFtpAllowedSlot prepares the UpdateFtpAllowedSlot request.
func (c WebAppsClient) preparerForUpdateFtpAllowedSlot(ctx context.Context, id SlotId, input CsmPublishingCredentialsPoliciesEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/ftp", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateFtpAllowedSlot handles the response to the UpdateFtpAllowedSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateFtpAllowedSlot(resp *http.Response) (result UpdateFtpAllowedSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
