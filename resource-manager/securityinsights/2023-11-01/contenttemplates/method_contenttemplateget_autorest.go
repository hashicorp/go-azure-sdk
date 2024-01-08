package contenttemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTemplateGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *TemplateModel
}

// ContentTemplateGet ...
func (c ContentTemplatesClient) ContentTemplateGet(ctx context.Context, id ContentTemplateId) (result ContentTemplateGetOperationResponse, err error) {
	req, err := c.preparerForContentTemplateGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentTemplateGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentTemplateGet prepares the ContentTemplateGet request.
func (c ContentTemplatesClient) preparerForContentTemplateGet(ctx context.Context, id ContentTemplateId) (*http.Request, error) {
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

// responderForContentTemplateGet handles the response to the ContentTemplateGet request. The method always
// closes the http.Response Body.
func (c ContentTemplatesClient) responderForContentTemplateGet(resp *http.Response) (result ContentTemplateGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
