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

type ResendEmailOperationResponse struct {
	HttpResponse *http.Response
}

// ResendEmail ...
func (c AppServiceCertificateOrdersClient) ResendEmail(ctx context.Context, id CertificateOrderId) (result ResendEmailOperationResponse, err error) {
	req, err := c.preparerForResendEmail(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendEmail", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendEmail", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResendEmail(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendEmail", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResendEmail prepares the ResendEmail request.
func (c AppServiceCertificateOrdersClient) preparerForResendEmail(ctx context.Context, id CertificateOrderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resendEmail", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResendEmail handles the response to the ResendEmail request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForResendEmail(resp *http.Response) (result ResendEmailOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
