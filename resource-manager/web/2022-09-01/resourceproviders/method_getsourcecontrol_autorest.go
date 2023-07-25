package resourceproviders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSourceControlOperationResponse struct {
	HttpResponse *http.Response
	Model        *SourceControl
}

// GetSourceControl ...
func (c ResourceProvidersClient) GetSourceControl(ctx context.Context, id SourceControlId) (result GetSourceControlOperationResponse, err error) {
	req, err := c.preparerForGetSourceControl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSourceControl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSourceControl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSourceControl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSourceControl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSourceControl prepares the GetSourceControl request.
func (c ResourceProvidersClient) preparerForGetSourceControl(ctx context.Context, id SourceControlId) (*http.Request, error) {
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

// responderForGetSourceControl handles the response to the GetSourceControl request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForGetSourceControl(resp *http.Response) (result GetSourceControlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
