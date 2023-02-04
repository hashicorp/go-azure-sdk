package hypervcluster

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetClusterOperationResponse struct {
	HttpResponse *http.Response
	Model        *HyperVCluster
}

// GetCluster ...
func (c HyperVClusterClient) GetCluster(ctx context.Context, id ClusterId) (result GetClusterOperationResponse, err error) {
	req, err := c.preparerForGetCluster(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetCluster", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetCluster", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetCluster(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetCluster", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetCluster prepares the GetCluster request.
func (c HyperVClusterClient) preparerForGetCluster(ctx context.Context, id ClusterId) (*http.Request, error) {
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

// responderForGetCluster handles the response to the GetCluster request. The method always
// closes the http.Response Body.
func (c HyperVClusterClient) responderForGetCluster(resp *http.Response) (result GetClusterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
