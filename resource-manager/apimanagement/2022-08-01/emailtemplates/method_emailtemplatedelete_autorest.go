package emailtemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type EmailTemplateDeleteOperationOptions struct {
	IfMatch *string
}

func DefaultEmailTemplateDeleteOperationOptions() EmailTemplateDeleteOperationOptions {
	return EmailTemplateDeleteOperationOptions{}
}

func (o EmailTemplateDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o EmailTemplateDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// EmailTemplateDelete ...
func (c EmailTemplatesClient) EmailTemplateDelete(ctx context.Context, id TemplateId, options EmailTemplateDeleteOperationOptions) (result EmailTemplateDeleteOperationResponse, err error) {
	req, err := c.preparerForEmailTemplateDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEmailTemplateDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEmailTemplateDelete prepares the EmailTemplateDelete request.
func (c EmailTemplatesClient) preparerForEmailTemplateDelete(ctx context.Context, id TemplateId, options EmailTemplateDeleteOperationOptions) (*http.Request, error) {
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

// responderForEmailTemplateDelete handles the response to the EmailTemplateDelete request. The method always
// closes the http.Response Body.
func (c EmailTemplatesClient) responderForEmailTemplateDelete(resp *http.Response) (result EmailTemplateDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
