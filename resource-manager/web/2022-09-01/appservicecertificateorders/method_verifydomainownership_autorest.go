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

type VerifyDomainOwnershipOperationResponse struct {
	HttpResponse *http.Response
}

// VerifyDomainOwnership ...
func (c AppServiceCertificateOrdersClient) VerifyDomainOwnership(ctx context.Context, id CertificateOrderId) (result VerifyDomainOwnershipOperationResponse, err error) {
	req, err := c.preparerForVerifyDomainOwnership(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "VerifyDomainOwnership", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "VerifyDomainOwnership", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForVerifyDomainOwnership(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "VerifyDomainOwnership", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForVerifyDomainOwnership prepares the VerifyDomainOwnership request.
func (c AppServiceCertificateOrdersClient) preparerForVerifyDomainOwnership(ctx context.Context, id CertificateOrderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/verifyDomainOwnership", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForVerifyDomainOwnership handles the response to the VerifyDomainOwnership request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForVerifyDomainOwnership(resp *http.Response) (result VerifyDomainOwnershipOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
