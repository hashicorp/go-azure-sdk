package contentproducttemplates

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductTemplateGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProductTemplateModel
}

// ProductTemplateGet ...
func (c ContentProductTemplatesClient) ProductTemplateGet(ctx context.Context, id ContentproducttemplateId) (result ProductTemplateGetOperationResponse, err error) {
	req, err := c.preparerForProductTemplateGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplateGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplateGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductTemplateGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplateGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductTemplateGet prepares the ProductTemplateGet request.
func (c ContentProductTemplatesClient) preparerForProductTemplateGet(ctx context.Context, id ContentproducttemplateId) (*http.Request, error) {
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

// responderForProductTemplateGet handles the response to the ProductTemplateGet request. The method always
// closes the http.Response Body.
func (c ContentProductTemplatesClient) responderForProductTemplateGet(resp *http.Response) (result ProductTemplateGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
