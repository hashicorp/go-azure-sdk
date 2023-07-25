package domains

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
	Model        *DomainOwnershipIdentifier
}

// GetOwnershipIdentifier ...
func (c DomainsClient) GetOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (result GetOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForGetOwnershipIdentifier(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetOwnershipIdentifier prepares the GetOwnershipIdentifier request.
func (c DomainsClient) preparerForGetOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId) (*http.Request, error) {
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

// responderForGetOwnershipIdentifier handles the response to the GetOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForGetOwnershipIdentifier(resp *http.Response) (result GetOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
