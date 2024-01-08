package contenttemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTemplateInstallOperationResponse struct {
	HttpResponse *http.Response
	Model        *TemplateModel
}

// ContentTemplateInstall ...
func (c ContentTemplatesClient) ContentTemplateInstall(ctx context.Context, id ContentTemplateId, input TemplateModel) (result ContentTemplateInstallOperationResponse, err error) {
	req, err := c.preparerForContentTemplateInstall(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateInstall", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateInstall", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentTemplateInstall(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttemplates.ContentTemplatesClient", "ContentTemplateInstall", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentTemplateInstall prepares the ContentTemplateInstall request.
func (c ContentTemplatesClient) preparerForContentTemplateInstall(ctx context.Context, id ContentTemplateId, input TemplateModel) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentTemplateInstall handles the response to the ContentTemplateInstall request. The method always
// closes the http.Response Body.
func (c ContentTemplatesClient) responderForContentTemplateInstall(resp *http.Response) (result ContentTemplateInstallOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
