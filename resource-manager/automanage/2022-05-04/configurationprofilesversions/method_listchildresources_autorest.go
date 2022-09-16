package configurationprofilesversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListChildResourcesOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationProfileList
}

// ListChildResources ...
func (c ConfigurationProfilesVersionsClient) ListChildResources(ctx context.Context, id ConfigurationProfileId) (result ListChildResourcesOperationResponse, err error) {
	req, err := c.preparerForListChildResources(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofilesversions.ConfigurationProfilesVersionsClient", "ListChildResources", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofilesversions.ConfigurationProfilesVersionsClient", "ListChildResources", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListChildResources(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofilesversions.ConfigurationProfilesVersionsClient", "ListChildResources", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListChildResources prepares the ListChildResources request.
func (c ConfigurationProfilesVersionsClient) preparerForListChildResources(ctx context.Context, id ConfigurationProfileId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListChildResources handles the response to the ListChildResources request. The method always
// closes the http.Response Body.
func (c ConfigurationProfilesVersionsClient) responderForListChildResources(resp *http.Response) (result ListChildResourcesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
