package zone

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersListZonesOperationResponse struct {
	HttpResponse *http.Response
	Model        *ClusterZoneList
}

// ClustersListZones ...
func (c ZoneClient) ClustersListZones(ctx context.Context, id ClusterId) (result ClustersListZonesOperationResponse, err error) {
	req, err := c.preparerForClustersListZones(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "zone.ZoneClient", "ClustersListZones", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "zone.ZoneClient", "ClustersListZones", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForClustersListZones(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "zone.ZoneClient", "ClustersListZones", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForClustersListZones prepares the ClustersListZones request.
func (c ZoneClient) preparerForClustersListZones(ctx context.Context, id ClusterId) (*http.Request, error) {
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

// responderForClustersListZones handles the response to the ClustersListZones request. The method always
// closes the http.Response Body.
func (c ZoneClient) responderForClustersListZones(resp *http.Response) (result ClustersListZonesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
