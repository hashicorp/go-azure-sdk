package domains

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
	Model        *DomainOwnershipIdentifier
}

// CreateOrUpdateOwnershipIdentifier ...
func (c DomainsClient) CreateOrUpdateOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId, input DomainOwnershipIdentifier) (result CreateOrUpdateOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateOwnershipIdentifier(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "CreateOrUpdateOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "CreateOrUpdateOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "CreateOrUpdateOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateOwnershipIdentifier prepares the CreateOrUpdateOwnershipIdentifier request.
func (c DomainsClient) preparerForCreateOrUpdateOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId, input DomainOwnershipIdentifier) (*http.Request, error) {
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

// responderForCreateOrUpdateOwnershipIdentifier handles the response to the CreateOrUpdateOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForCreateOrUpdateOwnershipIdentifier(resp *http.Response) (result CreateOrUpdateOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
