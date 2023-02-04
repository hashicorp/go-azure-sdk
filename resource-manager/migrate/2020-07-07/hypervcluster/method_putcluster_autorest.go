package hypervcluster

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PutClusterOperationResponse struct {
	HttpResponse *http.Response
}

// PutCluster ...
func (c HyperVClusterClient) PutCluster(ctx context.Context, id ClusterId, input HyperVCluster) (result PutClusterOperationResponse, err error) {
	req, err := c.preparerForPutCluster(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "PutCluster", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "PutCluster", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutCluster(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "PutCluster", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutCluster prepares the PutCluster request.
func (c HyperVClusterClient) preparerForPutCluster(ctx context.Context, id ClusterId, input HyperVCluster) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPutCluster handles the response to the PutCluster request. The method always
// closes the http.Response Body.
func (c HyperVClusterClient) responderForPutCluster(resp *http.Response) (result PutClusterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
