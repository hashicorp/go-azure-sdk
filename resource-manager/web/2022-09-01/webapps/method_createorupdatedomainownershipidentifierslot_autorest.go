package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateDomainOwnershipIdentifierSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Identifier
}

// CreateOrUpdateDomainOwnershipIdentifierSlot ...
func (c WebAppsClient) CreateOrUpdateDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId, input Identifier) (result CreateOrUpdateDomainOwnershipIdentifierSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateDomainOwnershipIdentifierSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateDomainOwnershipIdentifierSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateDomainOwnershipIdentifierSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateDomainOwnershipIdentifierSlot prepares the CreateOrUpdateDomainOwnershipIdentifierSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId, input Identifier) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateDomainOwnershipIdentifierSlot handles the response to the CreateOrUpdateDomainOwnershipIdentifierSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateDomainOwnershipIdentifierSlot(resp *http.Response) (result CreateOrUpdateDomainOwnershipIdentifierSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
