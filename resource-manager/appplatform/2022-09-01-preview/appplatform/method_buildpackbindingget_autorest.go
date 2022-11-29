package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildpackBindingGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *BuildpackBindingResource
}

// BuildpackBindingGet ...
func (c AppPlatformClient) BuildpackBindingGet(ctx context.Context, id BuildPackBindingId) (result BuildpackBindingGetOperationResponse, err error) {
	req, err := c.preparerForBuildpackBindingGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBuildpackBindingGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBuildpackBindingGet prepares the BuildpackBindingGet request.
func (c AppPlatformClient) preparerForBuildpackBindingGet(ctx context.Context, id BuildPackBindingId) (*http.Request, error) {
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

// responderForBuildpackBindingGet handles the response to the BuildpackBindingGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildpackBindingGet(resp *http.Response) (result BuildpackBindingGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
