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

type RetrieveCertificateEmailHistoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CertificateEmail
}

// RetrieveCertificateEmailHistory ...
func (c AppServiceCertificateOrdersClient) RetrieveCertificateEmailHistory(ctx context.Context, id CertificateOrderId) (result RetrieveCertificateEmailHistoryOperationResponse, err error) {
	req, err := c.preparerForRetrieveCertificateEmailHistory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateEmailHistory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateEmailHistory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRetrieveCertificateEmailHistory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateEmailHistory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRetrieveCertificateEmailHistory prepares the RetrieveCertificateEmailHistory request.
func (c AppServiceCertificateOrdersClient) preparerForRetrieveCertificateEmailHistory(ctx context.Context, id CertificateOrderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/retrieveEmailHistory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRetrieveCertificateEmailHistory handles the response to the RetrieveCertificateEmailHistory request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForRetrieveCertificateEmailHistory(resp *http.Response) (result RetrieveCertificateEmailHistoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
