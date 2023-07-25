package resourceproviders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSourceControlOperationResponse struct {
	HttpResponse *http.Response
	Model        *SourceControl
}

// UpdateSourceControl ...
func (c ResourceProvidersClient) UpdateSourceControl(ctx context.Context, id SourceControlId, input SourceControl) (result UpdateSourceControlOperationResponse, err error) {
	req, err := c.preparerForUpdateSourceControl(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdateSourceControl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdateSourceControl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSourceControl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdateSourceControl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSourceControl prepares the UpdateSourceControl request.
func (c ResourceProvidersClient) preparerForUpdateSourceControl(ctx context.Context, id SourceControlId, input SourceControl) (*http.Request, error) {
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

// responderForUpdateSourceControl handles the response to the UpdateSourceControl request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForUpdateSourceControl(resp *http.Response) (result UpdateSourceControlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
