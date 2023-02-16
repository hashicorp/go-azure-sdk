package emailtemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EmailTemplateContract
}

// EmailTemplateGet ...
func (c EmailTemplatesClient) EmailTemplateGet(ctx context.Context, id TemplateId) (result EmailTemplateGetOperationResponse, err error) {
	req, err := c.preparerForEmailTemplateGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEmailTemplateGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailtemplates.EmailTemplatesClient", "EmailTemplateGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEmailTemplateGet prepares the EmailTemplateGet request.
func (c EmailTemplatesClient) preparerForEmailTemplateGet(ctx context.Context, id TemplateId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEmailTemplateGet handles the response to the EmailTemplateGet request. The method always
// closes the http.Response Body.
func (c EmailTemplatesClient) responderForEmailTemplateGet(resp *http.Response) (result EmailTemplateGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
