package domains

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

type GetControlCenterSsoRequestOperationResponse struct {
	HttpResponse *http.Response
	Model        *DomainControlCenterSsoRequest
}

// GetControlCenterSsoRequest ...
func (c DomainsClient) GetControlCenterSsoRequest(ctx context.Context, id commonids.SubscriptionId) (result GetControlCenterSsoRequestOperationResponse, err error) {
	req, err := c.preparerForGetControlCenterSsoRequest(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetControlCenterSsoRequest", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetControlCenterSsoRequest", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetControlCenterSsoRequest(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "GetControlCenterSsoRequest", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetControlCenterSsoRequest prepares the GetControlCenterSsoRequest request.
func (c DomainsClient) preparerForGetControlCenterSsoRequest(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.DomainRegistration/generateSsoRequest", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetControlCenterSsoRequest handles the response to the GetControlCenterSsoRequest request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForGetControlCenterSsoRequest(resp *http.Response) (result GetControlCenterSsoRequestOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
