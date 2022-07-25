package namespacesnetworksecurityperimeterconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// NetworkSecurityPerimeterConfigurationsCreateOrUpdate ...
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) NetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx context.Context, id NetworkSecurityPerimeterConfigurationId) (result NetworkSecurityPerimeterConfigurationsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForNetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient", "NetworkSecurityPerimeterConfigurationsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForNetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient", "NetworkSecurityPerimeterConfigurationsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// NetworkSecurityPerimeterConfigurationsCreateOrUpdateThenPoll performs NetworkSecurityPerimeterConfigurationsCreateOrUpdate then polls until it's completed
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) NetworkSecurityPerimeterConfigurationsCreateOrUpdateThenPoll(ctx context.Context, id NetworkSecurityPerimeterConfigurationId) error {
	result, err := c.NetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing NetworkSecurityPerimeterConfigurationsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after NetworkSecurityPerimeterConfigurationsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForNetworkSecurityPerimeterConfigurationsCreateOrUpdate prepares the NetworkSecurityPerimeterConfigurationsCreateOrUpdate request.
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) preparerForNetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx context.Context, id NetworkSecurityPerimeterConfigurationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reconcile", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForNetworkSecurityPerimeterConfigurationsCreateOrUpdate sends the NetworkSecurityPerimeterConfigurationsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c NamespacesNetworkSecurityPerimeterConfigurationsClient) senderForNetworkSecurityPerimeterConfigurationsCreateOrUpdate(ctx context.Context, req *http.Request) (future NetworkSecurityPerimeterConfigurationsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
