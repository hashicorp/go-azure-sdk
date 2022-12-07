package appplatform

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

type ConfigurationServicesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigurationServicesCreateOrUpdate ...
func (c AppPlatformClient) ConfigurationServicesCreateOrUpdate(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceResource) (result ConfigurationServicesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForConfigurationServicesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigurationServicesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigurationServicesCreateOrUpdateThenPoll performs ConfigurationServicesCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) ConfigurationServicesCreateOrUpdateThenPoll(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceResource) error {
	result, err := c.ConfigurationServicesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConfigurationServicesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigurationServicesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForConfigurationServicesCreateOrUpdate prepares the ConfigurationServicesCreateOrUpdate request.
func (c AppPlatformClient) preparerForConfigurationServicesCreateOrUpdate(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceResource) (*http.Request, error) {
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

// senderForConfigurationServicesCreateOrUpdate sends the ConfigurationServicesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigurationServicesCreateOrUpdate(ctx context.Context, req *http.Request) (future ConfigurationServicesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
