package emailtemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *EmailTemplateContract
}

type EmailTemplateUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultEmailTemplateUpdateOperationOptions() EmailTemplateUpdateOperationOptions {
	return EmailTemplateUpdateOperationOptions{}
}

func (o EmailTemplateUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o EmailTemplateUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// EmailTemplateUpdate ...
func (c EmailTemplatesClient) EmailTemplateUpdate(ctx context.Context, id TemplateId, input EmailTemplateUpdateParameters, options EmailTemplateUpdateOperationOptions) (result EmailTemplateUpdateOperationResponse, err error) {
	req, err := c.preparerForEmailTemplateUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEmailTemplateUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEmailTemplateUpdate prepares the EmailTemplateUpdate request.
func (c EmailTemplatesClient) preparerForEmailTemplateUpdate(ctx context.Context, id TemplateId, input EmailTemplateUpdateParameters, options EmailTemplateUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEmailTemplateUpdate handles the response to the EmailTemplateUpdate request. The method always
// closes the http.Response Body.
func (c EmailTemplatesClient) responderForEmailTemplateUpdate(resp *http.Response) (result EmailTemplateUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
