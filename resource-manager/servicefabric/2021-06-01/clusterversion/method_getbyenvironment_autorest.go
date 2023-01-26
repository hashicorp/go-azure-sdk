package clusterversion

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	Model        *ClusterCodeVersionsListResult
}

// GetByEnvironment ...
func (c ClusterVersionClient) GetByEnvironment(ctx context.Context, id EnvironmentClusterVersionId) (result GetByEnvironmentOperationResponse, err error) {
	req, err := c.preparerForGetByEnvironment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusterversion.ClusterVersionClient", "GetByEnvironment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusterversion.ClusterVersionClient", "GetByEnvironment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByEnvironment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusterversion.ClusterVersionClient", "GetByEnvironment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByEnvironment prepares the GetByEnvironment request.
func (c ClusterVersionClient) preparerForGetByEnvironment(ctx context.Context, id EnvironmentClusterVersionId) (*http.Request, error) {
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

// responderForGetByEnvironment handles the response to the GetByEnvironment request. The method always
// closes the http.Response Body.
func (c ClusterVersionClient) responderForGetByEnvironment(resp *http.Response) (result GetByEnvironmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
