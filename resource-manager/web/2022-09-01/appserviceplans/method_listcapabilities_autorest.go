package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCapabilitiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Capability
}

// ListCapabilities ...
func (c AppServicePlansClient) ListCapabilities(ctx context.Context, id ServerFarmId) (result ListCapabilitiesOperationResponse, err error) {
	req, err := c.preparerForListCapabilities(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListCapabilities", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListCapabilities", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListCapabilities(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListCapabilities", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListCapabilities prepares the ListCapabilities request.
func (c AppServicePlansClient) preparerForListCapabilities(ctx context.Context, id ServerFarmId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capabilities", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListCapabilities handles the response to the ListCapabilities request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListCapabilities(resp *http.Response) (result ListCapabilitiesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
