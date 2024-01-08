package contenttemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTemplateDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// ContentTemplateDelete ...
func (c ContentTemplatesClient) ContentTemplateDelete(ctx context.Context, id ContentTemplateId) (result ContentTemplateDeleteOperationResponse, err error) {
	req, err := c.preparerForContentTemplateDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentTemplateDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentTemplateDelete prepares the ContentTemplateDelete request.
func (c ContentTemplatesClient) preparerForContentTemplateDelete(ctx context.Context, id ContentTemplateId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentTemplateDelete handles the response to the ContentTemplateDelete request. The method always
// closes the http.Response Body.
func (c ContentTemplatesClient) responderForContentTemplateDelete(resp *http.Response) (result ContentTemplateDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
