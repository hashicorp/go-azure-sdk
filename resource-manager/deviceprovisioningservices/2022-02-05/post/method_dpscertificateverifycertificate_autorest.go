package post

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpsCertificateVerifyCertificateOperationResponse struct {
	HttpResponse *http.Response
	Model        *CertificateResponse
}

type DpsCertificateVerifyCertificateOperationOptions struct {
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

func DefaultDpsCertificateVerifyCertificateOperationOptions() DpsCertificateVerifyCertificateOperationOptions {
	return DpsCertificateVerifyCertificateOperationOptions{}
}

func (o DpsCertificateVerifyCertificateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o DpsCertificateVerifyCertificateOperationOptions) toQueryString() map[string]interface{} {
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

// DpsCertificateVerifyCertificate ...
func (c POSTClient) DpsCertificateVerifyCertificate(ctx context.Context, id CertificateId, input VerificationCodeRequest, options DpsCertificateVerifyCertificateOperationOptions) (result DpsCertificateVerifyCertificateOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateVerifyCertificate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateVerifyCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateVerifyCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateVerifyCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateVerifyCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateVerifyCertificate prepares the DpsCertificateVerifyCertificate request.
func (c POSTClient) preparerForDpsCertificateVerifyCertificate(ctx context.Context, id CertificateId, input VerificationCodeRequest, options DpsCertificateVerifyCertificateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/verify", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateVerifyCertificate handles the response to the DpsCertificateVerifyCertificate request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForDpsCertificateVerifyCertificate(resp *http.Response) (result DpsCertificateVerifyCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
