package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDomainOwnershipIdentifierSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteDomainOwnershipIdentifierSlot ...
func (c WebAppsClient) DeleteDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId) (result DeleteDomainOwnershipIdentifierSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteDomainOwnershipIdentifierSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifierSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteDomainOwnershipIdentifierSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifierSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteDomainOwnershipIdentifierSlot prepares the DeleteDomainOwnershipIdentifierSlot request.
func (c WebAppsClient) preparerForDeleteDomainOwnershipIdentifierSlot(ctx context.Context, id SlotDomainOwnershipIdentifierId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteDomainOwnershipIdentifierSlot handles the response to the DeleteDomainOwnershipIdentifierSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteDomainOwnershipIdentifierSlot(resp *http.Response) (result DeleteDomainOwnershipIdentifierSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
