package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceRegistriesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServiceRegistryResource
}

// ServiceRegistriesGet ...
func (c AppPlatformClient) ServiceRegistriesGet(ctx context.Context, id ServiceRegistryId) (result ServiceRegistriesGetOperationResponse, err error) {
	req, err := c.preparerForServiceRegistriesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServiceRegistriesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServiceRegistriesGet prepares the ServiceRegistriesGet request.
func (c AppPlatformClient) preparerForServiceRegistriesGet(ctx context.Context, id ServiceRegistryId) (*http.Request, error) {
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

// responderForServiceRegistriesGet handles the response to the ServiceRegistriesGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForServiceRegistriesGet(resp *http.Response) (result ServiceRegistriesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
