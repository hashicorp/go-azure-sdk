package resourceproviders

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

type GetSubscriptionDeploymentLocationsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentLocations
}

// GetSubscriptionDeploymentLocations ...
func (c ResourceProvidersClient) GetSubscriptionDeploymentLocations(ctx context.Context, id commonids.SubscriptionId) (result GetSubscriptionDeploymentLocationsOperationResponse, err error) {
	req, err := c.preparerForGetSubscriptionDeploymentLocations(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSubscriptionDeploymentLocations", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSubscriptionDeploymentLocations", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSubscriptionDeploymentLocations(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetSubscriptionDeploymentLocations", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSubscriptionDeploymentLocations prepares the GetSubscriptionDeploymentLocations request.
func (c ResourceProvidersClient) preparerForGetSubscriptionDeploymentLocations(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/deploymentLocations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSubscriptionDeploymentLocations handles the response to the GetSubscriptionDeploymentLocations request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForGetSubscriptionDeploymentLocations(resp *http.Response) (result GetSubscriptionDeploymentLocationsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
