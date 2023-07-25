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

type RetrieveCertificateActionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CertificateOrderAction
}

// RetrieveCertificateActions ...
func (c AppServiceCertificateOrdersClient) RetrieveCertificateActions(ctx context.Context, id CertificateOrderId) (result RetrieveCertificateActionsOperationResponse, err error) {
	req, err := c.preparerForRetrieveCertificateActions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateActions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateActions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRetrieveCertificateActions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveCertificateActions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRetrieveCertificateActions prepares the RetrieveCertificateActions request.
func (c AppServiceCertificateOrdersClient) preparerForRetrieveCertificateActions(ctx context.Context, id CertificateOrderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/retrieveCertificateActions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRetrieveCertificateActions handles the response to the RetrieveCertificateActions request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForRetrieveCertificateActions(resp *http.Response) (result RetrieveCertificateActionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
