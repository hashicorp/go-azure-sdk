package serviceprincipals

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

type ServicePrincipalsListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServicePrincipalListResult
}

// ServicePrincipalsListBySubscription ...
func (c ServicePrincipalsClient) ServicePrincipalsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result ServicePrincipalsListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForServicePrincipalsListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceprincipals.ServicePrincipalsClient", "ServicePrincipalsListBySubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceprincipals.ServicePrincipalsClient", "ServicePrincipalsListBySubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicePrincipalsListBySubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceprincipals.ServicePrincipalsClient", "ServicePrincipalsListBySubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicePrincipalsListBySubscription prepares the ServicePrincipalsListBySubscription request.
func (c ServicePrincipalsClient) preparerForServicePrincipalsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.AutoManage/servicePrincipals", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServicePrincipalsListBySubscription handles the response to the ServicePrincipalsListBySubscription request. The method always
// closes the http.Response Body.
func (c ServicePrincipalsClient) responderForServicePrincipalsListBySubscription(resp *http.Response) (result ServicePrincipalsListBySubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
