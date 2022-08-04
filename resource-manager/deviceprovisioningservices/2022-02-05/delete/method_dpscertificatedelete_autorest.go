package delete

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpsCertificateDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type DpsCertificateDeleteOperationOptions struct {
	CertificateCreated       *string
	CertificateHasPrivateKey *bool
	CertificateIsVerified    *bool
	CertificateLastUpdated   *string
	CertificateName          *string
	CertificateNonce         *string
	CertificatePurpose       *CertificatePurpose
	CertificateRawBytes      *string
	IfMatch                  *string
}

func DefaultDpsCertificateDeleteOperationOptions() DpsCertificateDeleteOperationOptions {
	return DpsCertificateDeleteOperationOptions{}
}

func (o DpsCertificateDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o DpsCertificateDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.CertificateCreated != nil {
		out["certificate.created"] = *o.CertificateCreated
	}

	if o.CertificateHasPrivateKey != nil {
		out["certificate.hasPrivateKey"] = *o.CertificateHasPrivateKey
	}

	if o.CertificateIsVerified != nil {
		out["certificate.isVerified"] = *o.CertificateIsVerified
	}

	if o.CertificateLastUpdated != nil {
		out["certificate.lastUpdated"] = *o.CertificateLastUpdated
	}

	if o.CertificateName != nil {
		out["certificate.name"] = *o.CertificateName
	}

	if o.CertificateNonce != nil {
		out["certificate.nonce"] = *o.CertificateNonce
	}

	if o.CertificatePurpose != nil {
		out["certificate.purpose"] = *o.CertificatePurpose
	}

	if o.CertificateRawBytes != nil {
		out["certificate.rawBytes"] = *o.CertificateRawBytes
	}

	return out
}

// DpsCertificateDelete ...
func (c DELETEClient) DpsCertificateDelete(ctx context.Context, id CertificateId, options DpsCertificateDeleteOperationOptions) (result DpsCertificateDeleteOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "DpsCertificateDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "DpsCertificateDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "DpsCertificateDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateDelete prepares the DpsCertificateDelete request.
func (c DELETEClient) preparerForDpsCertificateDelete(ctx context.Context, id CertificateId, options DpsCertificateDeleteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateDelete handles the response to the DpsCertificateDelete request. The method always
// closes the http.Response Body.
func (c DELETEClient) responderForDpsCertificateDelete(resp *http.Response) (result DpsCertificateDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
