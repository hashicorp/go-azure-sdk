package netappresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceQueryRegionInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegionInfo
}

// NetAppResourceQueryRegionInfo ...
func (c NetAppResourceClient) NetAppResourceQueryRegionInfo(ctx context.Context, id LocationId) (result NetAppResourceQueryRegionInfoOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceQueryRegionInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQueryRegionInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQueryRegionInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceQueryRegionInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQueryRegionInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceQueryRegionInfo prepares the NetAppResourceQueryRegionInfo request.
func (c NetAppResourceClient) preparerForNetAppResourceQueryRegionInfo(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regionInfo", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceQueryRegionInfo handles the response to the NetAppResourceQueryRegionInfo request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceQueryRegionInfo(resp *http.Response) (result NetAppResourceQueryRegionInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
