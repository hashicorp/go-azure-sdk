package emailservices

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

type ListVerifiedExchangeOnlineDomainsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]string
}

// ListVerifiedExchangeOnlineDomains ...
func (c EmailServicesClient) ListVerifiedExchangeOnlineDomains(ctx context.Context, id commonids.SubscriptionId) (result ListVerifiedExchangeOnlineDomainsOperationResponse, err error) {
	req, err := c.preparerForListVerifiedExchangeOnlineDomains(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailservices.EmailServicesClient", "ListVerifiedExchangeOnlineDomains", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailservices.EmailServicesClient", "ListVerifiedExchangeOnlineDomains", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListVerifiedExchangeOnlineDomains(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailservices.EmailServicesClient", "ListVerifiedExchangeOnlineDomains", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListVerifiedExchangeOnlineDomains prepares the ListVerifiedExchangeOnlineDomains request.
func (c EmailServicesClient) preparerForListVerifiedExchangeOnlineDomains(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Communication/listVerifiedExchangeOnlineDomains", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListVerifiedExchangeOnlineDomains handles the response to the ListVerifiedExchangeOnlineDomains request. The method always
// closes the http.Response Body.
func (c EmailServicesClient) responderForListVerifiedExchangeOnlineDomains(resp *http.Response) (result ListVerifiedExchangeOnlineDomainsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
