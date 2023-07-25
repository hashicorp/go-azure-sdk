package domains

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteOwnershipIdentifier ...
func (c DomainsClient) DeleteOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (result DeleteOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForDeleteOwnershipIdentifier(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "DeleteOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "DeleteOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "DeleteOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteOwnershipIdentifier prepares the DeleteOwnershipIdentifier request.
func (c DomainsClient) preparerForDeleteOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (*http.Request, error) {
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

// responderForDeleteOwnershipIdentifier handles the response to the DeleteOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForDeleteOwnershipIdentifier(resp *http.Response) (result DeleteOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
