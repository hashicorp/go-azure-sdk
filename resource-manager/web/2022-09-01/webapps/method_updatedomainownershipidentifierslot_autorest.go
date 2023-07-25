package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDomainOwnershipIdentifierSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Identifier
}

// UpdateDomainOwnershipIdentifierSlot ...
func (c WebAppsClient) UpdateDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId, input Identifier) (result UpdateDomainOwnershipIdentifierSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateDomainOwnershipIdentifierSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifierSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateDomainOwnershipIdentifierSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateDomainOwnershipIdentifierSlot prepares the UpdateDomainOwnershipIdentifierSlot request.
func (c WebAppsClient) preparerForUpdateDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId, input Identifier) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateDomainOwnershipIdentifierSlot handles the response to the UpdateDomainOwnershipIdentifierSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateDomainOwnershipIdentifierSlot(resp *http.Response) (result UpdateDomainOwnershipIdentifierSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
