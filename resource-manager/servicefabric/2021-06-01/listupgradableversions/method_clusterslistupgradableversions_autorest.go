package listupgradableversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersListUpgradableVersionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *UpgradableVersionPathResult
}

// ClustersListUpgradableVersions ...
func (c ListUpgradableVersionsClient) ClustersListUpgradableVersions(ctx context.Context, id ClusterId, input UpgradableVersionsDescription) (result ClustersListUpgradableVersionsOperationResponse, err error) {
	req, err := c.preparerForClustersListUpgradableVersions(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listupgradableversions.ListUpgradableVersionsClient", "ClustersListUpgradableVersions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "listupgradableversions.ListUpgradableVersionsClient", "ClustersListUpgradableVersions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForClustersListUpgradableVersions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listupgradableversions.ListUpgradableVersionsClient", "ClustersListUpgradableVersions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForClustersListUpgradableVersions prepares the ClustersListUpgradableVersions request.
func (c ListUpgradableVersionsClient) preparerForClustersListUpgradableVersions(ctx context.Context, id ClusterId, input UpgradableVersionsDescription) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listUpgradableVersions", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForClustersListUpgradableVersions handles the response to the ClustersListUpgradableVersions request. The method always
// closes the http.Response Body.
func (c ListUpgradableVersionsClient) responderForClustersListUpgradableVersions(resp *http.Response) (result ClustersListUpgradableVersionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
