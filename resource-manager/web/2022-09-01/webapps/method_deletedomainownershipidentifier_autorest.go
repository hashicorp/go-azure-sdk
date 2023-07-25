package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDomainOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteDomainOwnershipIdentifier ...
func (c WebAppsClient) DeleteDomainOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (result DeleteDomainOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForDeleteDomainOwnershipIdentifier(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteDomainOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDomainOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteDomainOwnershipIdentifier prepares the DeleteDomainOwnershipIdentifier request.
func (c WebAppsClient) preparerForDeleteDomainOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (*http.Request, error) {
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

// responderForDeleteDomainOwnershipIdentifier handles the response to the DeleteDomainOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteDomainOwnershipIdentifier(resp *http.Response) (result DeleteDomainOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
