package cluster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListZonesOperationResponse struct {
	HttpResponse *http.Response
	Model        *ClusterZoneList
}

// ListZones ...
func (c ClusterClient) ListZones(ctx context.Context, id ClusterId) (result ListZonesOperationResponse, err error) {
	req, err := c.preparerForListZones(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cluster.ClusterClient", "ListZones", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cluster.ClusterClient", "ListZones", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListZones(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cluster.ClusterClient", "ListZones", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListZones prepares the ListZones request.
func (c ClusterClient) preparerForListZones(ctx context.Context, id ClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listZones", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListZones handles the response to the ListZones request. The method always
// closes the http.Response Body.
func (c ClusterClient) responderForListZones(resp *http.Response) (result ListZonesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
