package managedclusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListClusterUserCredentialsOperationResponse struct {
	HttpResponse *http.Response
	Model        *CredentialResults
}

// ListClusterUserCredentials ...
func (c ManagedClustersClient) ListClusterUserCredentials(ctx context.Context, id ManagedClusterId) (result ListClusterUserCredentialsOperationResponse, err error) {
	req, err := c.preparerForListClusterUserCredentials(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedclusters.ManagedClustersClient", "ListClusterUserCredentials", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedclusters.ManagedClustersClient", "ListClusterUserCredentials", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListClusterUserCredentials(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedclusters.ManagedClustersClient", "ListClusterUserCredentials", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListClusterUserCredentials prepares the ListClusterUserCredentials request.
func (c ManagedClustersClient) preparerForListClusterUserCredentials(ctx context.Context, id ManagedClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listClusterUserCredential", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListClusterUserCredentials handles the response to the ListClusterUserCredentials request. The method always
// closes the http.Response Body.
func (c ManagedClustersClient) responderForListClusterUserCredentials(resp *http.Response) (result ListClusterUserCredentialsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
