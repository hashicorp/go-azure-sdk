package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDomainOwnershipIdentifierSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Identifier
}

// GetDomainOwnershipIdentifierSlot ...
func (c WebAppsClient) GetDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId) (result GetDomainOwnershipIdentifierSlotOperationResponse, err error) {
	req, err := c.preparerForGetDomainOwnershipIdentifierSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDomainOwnershipIdentifierSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDomainOwnershipIdentifierSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDomainOwnershipIdentifierSlot prepares the GetDomainOwnershipIdentifierSlot request.
func (c WebAppsClient) preparerForGetDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId) (*http.Request, error) {
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

// responderForGetDomainOwnershipIdentifierSlot handles the response to the GetDomainOwnershipIdentifierSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetDomainOwnershipIdentifierSlot(resp *http.Response) (result GetDomainOwnershipIdentifierSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
