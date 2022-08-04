package put

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpsCertificateCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *CertificateResponse
}

type DpsCertificateCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultDpsCertificateCreateOrUpdateOperationOptions() DpsCertificateCreateOrUpdateOperationOptions {
	return DpsCertificateCreateOrUpdateOperationOptions{}
}

func (o DpsCertificateCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o DpsCertificateCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// DpsCertificateCreateOrUpdate ...
func (c PUTClient) DpsCertificateCreateOrUpdate(ctx context.Context, id CertificateId, input CertificateResponse, options DpsCertificateCreateOrUpdateOperationOptions) (result DpsCertificateCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "DpsCertificateCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "DpsCertificateCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "DpsCertificateCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateCreateOrUpdate prepares the DpsCertificateCreateOrUpdate request.
func (c PUTClient) preparerForDpsCertificateCreateOrUpdate(ctx context.Context, id CertificateId, input CertificateResponse, options DpsCertificateCreateOrUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateCreateOrUpdate handles the response to the DpsCertificateCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c PUTClient) responderForDpsCertificateCreateOrUpdate(resp *http.Response) (result DpsCertificateCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
