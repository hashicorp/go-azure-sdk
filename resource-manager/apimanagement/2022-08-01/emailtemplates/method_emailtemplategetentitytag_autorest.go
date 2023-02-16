package emailtemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateGetEntityTagOperationResponse struct {
	HttpResponse *http.Response
}

// EmailTemplateGetEntityTag ...
func (c EmailTemplatesClient) EmailTemplateGetEntityTag(ctx context.Context, id TemplateId) (result EmailTemplateGetEntityTagOperationResponse, err error) {
	req, err := c.preparerForEmailTemplateGetEntityTag(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGetEntityTag", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGetEntityTag", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEmailTemplateGetEntityTag(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGetEntityTag", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEmailTemplateGetEntityTag prepares the EmailTemplateGetEntityTag request.
func (c EmailTemplatesClient) preparerForEmailTemplateGetEntityTag(ctx context.Context, id TemplateId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEmailTemplateGetEntityTag handles the response to the EmailTemplateGetEntityTag request. The method always
// closes the http.Response Body.
func (c EmailTemplatesClient) responderForEmailTemplateGetEntityTag(resp *http.Response) (result EmailTemplateGetEntityTagOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
