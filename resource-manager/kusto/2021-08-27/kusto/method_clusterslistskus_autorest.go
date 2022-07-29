package kusto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersListSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *SkuDescriptionList
}

// ClustersListSkus ...
func (c KustoClient) ClustersListSkus(ctx context.Context, id commonids.SubscriptionId) (result ClustersListSkusOperationResponse, err error) {
	req, err := c.preparerForClustersListSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "ClustersListSkus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "ClustersListSkus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForClustersListSkus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "ClustersListSkus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForClustersListSkus prepares the ClustersListSkus request.
func (c KustoClient) preparerForClustersListSkus(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Kusto/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForClustersListSkus handles the response to the ClustersListSkus request. The method always
// closes the http.Response Body.
func (c KustoClient) responderForClustersListSkus(resp *http.Response) (result ClustersListSkusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
