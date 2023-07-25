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

type ReissueOperationResponse struct {
	HttpResponse *http.Response
}

// Reissue ...
func (c AppServiceCertificateOrdersClient) Reissue(ctx context.Context, id CertificateOrderId, input ReissueCertificateOrderRequest) (result ReissueOperationResponse, err error) {
	req, err := c.preparerForReissue(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Reissue", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Reissue", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForReissue(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "Reissue", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForReissue prepares the Reissue request.
func (c AppServiceCertificateOrdersClient) preparerForReissue(ctx context.Context, id CertificateOrderId, input ReissueCertificateOrderRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reissue", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForReissue handles the response to the Reissue request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForReissue(resp *http.Response) (result ReissueOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
