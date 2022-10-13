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

type ConfigurationServicesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigurationServicesDelete ...
func (c AppPlatformClient) ConfigurationServicesDelete(ctx context.Context, id ConfigurationServiceId) (result ConfigurationServicesDeleteOperationResponse, err error) {
	req, err := c.preparerForConfigurationServicesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigurationServicesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigurationServicesDeleteThenPoll performs ConfigurationServicesDelete then polls until it's completed
func (c AppPlatformClient) ConfigurationServicesDeleteThenPoll(ctx context.Context, id ConfigurationServiceId) error {
	result, err := c.ConfigurationServicesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ConfigurationServicesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigurationServicesDelete: %+v", err)
	}

	return nil
}

// preparerForConfigurationServicesDelete prepares the ConfigurationServicesDelete request.
func (c AppPlatformClient) preparerForConfigurationServicesDelete(ctx context.Context, id ConfigurationServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForConfigurationServicesDelete sends the ConfigurationServicesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigurationServicesDelete(ctx context.Context, req *http.Request) (future ConfigurationServicesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
