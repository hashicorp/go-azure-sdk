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

type RetrieveSiteSealOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteSeal
}

// RetrieveSiteSeal ...
func (c AppServiceCertificateOrdersClient) RetrieveSiteSeal(ctx context.Context, id CertificateOrderId, input SiteSealRequest) (result RetrieveSiteSealOperationResponse, err error) {
	req, err := c.preparerForRetrieveSiteSeal(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveSiteSeal", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveSiteSeal", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRetrieveSiteSeal(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "RetrieveSiteSeal", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRetrieveSiteSeal prepares the RetrieveSiteSeal request.
func (c AppServiceCertificateOrdersClient) preparerForRetrieveSiteSeal(ctx context.Context, id CertificateOrderId, input SiteSealRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/retrieveSiteSeal", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRetrieveSiteSeal handles the response to the RetrieveSiteSeal request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForRetrieveSiteSeal(resp *http.Response) (result RetrieveSiteSealOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
