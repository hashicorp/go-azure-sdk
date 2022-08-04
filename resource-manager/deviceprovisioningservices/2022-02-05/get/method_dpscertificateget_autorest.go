package get

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpsCertificateGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *CertificateResponse
}

type DpsCertificateGetOperationOptions struct {
	IfMatch *string
}

func DefaultDpsCertificateGetOperationOptions() DpsCertificateGetOperationOptions {
	return DpsCertificateGetOperationOptions{}
}

func (o DpsCertificateGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o DpsCertificateGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// DpsCertificateGet ...
func (c GETClient) DpsCertificateGet(ctx context.Context, id CertificateId, options DpsCertificateGetOperationOptions) (result DpsCertificateGetOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateGet prepares the DpsCertificateGet request.
func (c GETClient) preparerForDpsCertificateGet(ctx context.Context, id CertificateId, options DpsCertificateGetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateGet handles the response to the DpsCertificateGet request. The method always
// closes the http.Response Body.
func (c GETClient) responderForDpsCertificateGet(resp *http.Response) (result DpsCertificateGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
