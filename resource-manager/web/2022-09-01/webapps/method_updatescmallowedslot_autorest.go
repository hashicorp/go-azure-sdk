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

type UpdateScmAllowedSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// UpdateScmAllowedSlot ...
func (c WebAppsClient) UpdateScmAllowedSlot(ctx context.Context, id SlotId, input CsmPublishingCredentialsPoliciesEntity) (result UpdateScmAllowedSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateScmAllowedSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowedSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowedSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateScmAllowedSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowedSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateScmAllowedSlot prepares the UpdateScmAllowedSlot request.
func (c WebAppsClient) preparerForUpdateScmAllowedSlot(ctx context.Context, id SlotId, input CsmPublishingCredentialsPoliciesEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/scm", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateScmAllowedSlot handles the response to the UpdateScmAllowedSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateScmAllowedSlot(resp *http.Response) (result UpdateScmAllowedSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
