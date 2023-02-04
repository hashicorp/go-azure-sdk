package hypervsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteUsageOperationResponse struct {
	HttpResponse *http.Response
	Model        *HyperVSiteUsage
}

// GetSiteUsage ...
func (c HyperVSitesClient) GetSiteUsage(ctx context.Context, id HyperVSiteId) (result GetSiteUsageOperationResponse, err error) {
	req, err := c.preparerForGetSiteUsage(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSiteUsage", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSiteUsage", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteUsage(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSiteUsage", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteUsage prepares the GetSiteUsage request.
func (c HyperVSitesClient) preparerForGetSiteUsage(ctx context.Context, id HyperVSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/summary", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSiteUsage handles the response to the GetSiteUsage request. The method always
// closes the http.Response Body.
func (c HyperVSitesClient) responderForGetSiteUsage(resp *http.Response) (result GetSiteUsageOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
