package emailtemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *EmailTemplateContract
}

type EmailTemplateCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultEmailTemplateCreateOrUpdateOperationOptions() EmailTemplateCreateOrUpdateOperationOptions {
	return EmailTemplateCreateOrUpdateOperationOptions{}
}

func (o EmailTemplateCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o EmailTemplateCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// EmailTemplateCreateOrUpdate ...
func (c EmailTemplatesClient) EmailTemplateCreateOrUpdate(ctx context.Context, id TemplateId, input EmailTemplateUpdateParameters, options EmailTemplateCreateOrUpdateOperationOptions) (result EmailTemplateCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForEmailTemplateCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEmailTemplateCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEmailTemplateCreateOrUpdate prepares the EmailTemplateCreateOrUpdate request.
func (c EmailTemplatesClient) preparerForEmailTemplateCreateOrUpdate(ctx context.Context, id TemplateId, input EmailTemplateUpdateParameters, options EmailTemplateCreateOrUpdateOperationOptions) (*http.Request, error) {
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

// responderForEmailTemplateCreateOrUpdate handles the response to the EmailTemplateCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c EmailTemplatesClient) responderForEmailTemplateCreateOrUpdate(resp *http.Response) (result EmailTemplateCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
