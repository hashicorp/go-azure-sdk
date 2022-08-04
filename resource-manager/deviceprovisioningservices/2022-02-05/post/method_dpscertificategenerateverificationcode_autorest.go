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

type DpsCertificateGenerateVerificationCodeOperationResponse struct {
	HttpResponse *http.Response
	Model        *VerificationCodeResponse
}

type DpsCertificateGenerateVerificationCodeOperationOptions struct {
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

func DefaultDpsCertificateGenerateVerificationCodeOperationOptions() DpsCertificateGenerateVerificationCodeOperationOptions {
	return DpsCertificateGenerateVerificationCodeOperationOptions{}
}

func (o DpsCertificateGenerateVerificationCodeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o DpsCertificateGenerateVerificationCodeOperationOptions) toQueryString() map[string]interface{} {
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

// DpsCertificateGenerateVerificationCode ...
func (c POSTClient) DpsCertificateGenerateVerificationCode(ctx context.Context, id CertificateId, options DpsCertificateGenerateVerificationCodeOperationOptions) (result DpsCertificateGenerateVerificationCodeOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateGenerateVerificationCode(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateGenerateVerificationCode", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateGenerateVerificationCode", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateGenerateVerificationCode(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "DpsCertificateGenerateVerificationCode", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateGenerateVerificationCode prepares the DpsCertificateGenerateVerificationCode request.
func (c POSTClient) preparerForDpsCertificateGenerateVerificationCode(ctx context.Context, id CertificateId, options DpsCertificateGenerateVerificationCodeOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/generateVerificationCode", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateGenerateVerificationCode handles the response to the DpsCertificateGenerateVerificationCode request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForDpsCertificateGenerateVerificationCode(resp *http.Response) (result DpsCertificateGenerateVerificationCodeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
