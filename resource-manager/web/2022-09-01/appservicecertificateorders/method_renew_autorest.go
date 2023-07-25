package appservicecertificateorders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RenewOperationResponse struct {
	HttpResponse *http.Response
}

// Renew ...
func (c AppServiceCertificateOrdersClient) Renew(ctx context.Context, id CertificateOrderId, input RenewCertificateOrderRequest) (result RenewOperationResponse, err error) {
	req, err := c.preparerForRenew(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Renew", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Renew", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRenew(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Renew", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRenew prepares the Renew request.
func (c AppServiceCertificateOrdersClient) preparerForRenew(ctx context.Context, id CertificateOrderId, input RenewCertificateOrderRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/renew", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRenew handles the response to the Renew request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForRenew(resp *http.Response) (result RenewOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
