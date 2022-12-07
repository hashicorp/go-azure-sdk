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

type ConfigurationServicesValidateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigurationServicesValidate ...
func (c AppPlatformClient) ConfigurationServicesValidate(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceSettings) (result ConfigurationServicesValidateOperationResponse, err error) {
	req, err := c.preparerForConfigurationServicesValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesValidate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigurationServicesValidate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigurationServicesValidate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigurationServicesValidateThenPoll performs ConfigurationServicesValidate then polls until it's completed
func (c AppPlatformClient) ConfigurationServicesValidateThenPoll(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceSettings) error {
	result, err := c.ConfigurationServicesValidate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConfigurationServicesValidate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigurationServicesValidate: %+v", err)
	}

	return nil
}

// preparerForConfigurationServicesValidate prepares the ConfigurationServicesValidate request.
func (c AppPlatformClient) preparerForConfigurationServicesValidate(ctx context.Context, id ConfigurationServiceId, input ConfigurationServiceSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForConfigurationServicesValidate sends the ConfigurationServicesValidate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigurationServicesValidate(ctx context.Context, req *http.Request) (future ConfigurationServicesValidateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
