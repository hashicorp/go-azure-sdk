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

type ResendRequestEmailsOperationResponse struct {
	HttpResponse *http.Response
}

// ResendRequestEmails ...
func (c AppServiceCertificateOrdersClient) ResendRequestEmails(ctx context.Context, id CertificateOrderId, input NameIdentifier) (result ResendRequestEmailsOperationResponse, err error) {
	req, err := c.preparerForResendRequestEmails(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendRequestEmails", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendRequestEmails", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResendRequestEmails(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ResendRequestEmails", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResendRequestEmails prepares the ResendRequestEmails request.
func (c AppServiceCertificateOrdersClient) preparerForResendRequestEmails(ctx context.Context, id CertificateOrderId, input NameIdentifier) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resendRequestEmails", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResendRequestEmails handles the response to the ResendRequestEmails request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForResendRequestEmails(resp *http.Response) (result ResendRequestEmailsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
