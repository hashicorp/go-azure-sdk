package namespacesnetworksecurityperimeterconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationListOperationResponse struct {
	HttpResponse *http.Response
	Model        *NetworkSecurityPerimeterConfigurationList
}

// NetworkSecurityPerimeterConfigurationList ...
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) NetworkSecurityPerimeterConfigurationList(ctx context.Context, id NamespaceId) (result NetworkSecurityPerimeterConfigurationListOperationResponse, err error) {
	req, err := c.preparerForNetworkSecurityPerimeterConfigurationList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient", "NetworkSecurityPerimeterConfigurationList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient", "NetworkSecurityPerimeterConfigurationList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetworkSecurityPerimeterConfigurationList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient", "NetworkSecurityPerimeterConfigurationList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetworkSecurityPerimeterConfigurationList prepares the NetworkSecurityPerimeterConfigurationList request.
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) preparerForNetworkSecurityPerimeterConfigurationList(ctx context.Context, id NamespaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkSecurityPerimeterConfigurations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetworkSecurityPerimeterConfigurationList handles the response to the NetworkSecurityPerimeterConfigurationList request. The method always
// closes the http.Response Body.
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) responderForNetworkSecurityPerimeterConfigurationList(resp *http.Response) (result NetworkSecurityPerimeterConfigurationListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
