package apimanagementservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDomainOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiManagementServiceGetDomainOwnershipIdentifierResult
}

// GetDomainOwnershipIdentifier ...
func (c ApiManagementServiceClient) GetDomainOwnershipIdentifier(ctx context.Context, id commonids.SubscriptionId) (result GetDomainOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForGetDomainOwnershipIdentifier(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetDomainOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetDomainOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDomainOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetDomainOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDomainOwnershipIdentifier prepares the GetDomainOwnershipIdentifier request.
func (c ApiManagementServiceClient) preparerForGetDomainOwnershipIdentifier(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.ApiManagement/getDomainOwnershipIdentifier", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDomainOwnershipIdentifier handles the response to the GetDomainOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c ApiManagementServiceClient) responderForGetDomainOwnershipIdentifier(resp *http.Response) (result GetDomainOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
