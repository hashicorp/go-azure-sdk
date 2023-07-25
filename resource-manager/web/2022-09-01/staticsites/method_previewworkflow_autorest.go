package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreviewWorkflowOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSitesWorkflowPreview
}

// PreviewWorkflow ...
func (c StaticSitesClient) PreviewWorkflow(ctx context.Context, id ProviderLocationId, input StaticSitesWorkflowPreviewRequest) (result PreviewWorkflowOperationResponse, err error) {
	req, err := c.preparerForPreviewWorkflow(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "PreviewWorkflow", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "PreviewWorkflow", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPreviewWorkflow(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "PreviewWorkflow", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPreviewWorkflow prepares the PreviewWorkflow request.
func (c StaticSitesClient) preparerForPreviewWorkflow(ctx context.Context, id ProviderLocationId, input StaticSitesWorkflowPreviewRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/previewStaticSiteWorkflowFile", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPreviewWorkflow handles the response to the PreviewWorkflow request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForPreviewWorkflow(resp *http.Response) (result PreviewWorkflowOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
