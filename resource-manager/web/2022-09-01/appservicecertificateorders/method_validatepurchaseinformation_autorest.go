package appservicecertificateorders

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

type ValidatePurchaseInformationOperationResponse struct {
	HttpResponse *http.Response
}

// ValidatePurchaseInformation ...
func (c AppServiceCertificateOrdersClient) ValidatePurchaseInformation(ctx context.Context, id commonids.SubscriptionId, input AppServiceCertificateOrder) (result ValidatePurchaseInformationOperationResponse, err error) {
	req, err := c.preparerForValidatePurchaseInformation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ValidatePurchaseInformation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ValidatePurchaseInformation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForValidatePurchaseInformation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ValidatePurchaseInformation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForValidatePurchaseInformation prepares the ValidatePurchaseInformation request.
func (c AppServiceCertificateOrdersClient) preparerForValidatePurchaseInformation(ctx context.Context, id commonids.SubscriptionId, input AppServiceCertificateOrder) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CertificateRegistration/validateCertificateRegistrationInformation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForValidatePurchaseInformation handles the response to the ValidatePurchaseInformation request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForValidatePurchaseInformation(resp *http.Response) (result ValidatePurchaseInformationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
