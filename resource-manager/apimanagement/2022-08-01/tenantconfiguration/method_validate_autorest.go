package tenantconfiguration

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

type ValidateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Validate ...
func (c TenantConfigurationClient) Validate(ctx context.Context, id ServiceId, input DeployConfigurationParameters) (result ValidateOperationResponse, err error) {
	req, err := c.preparerForValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Validate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Validate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateThenPoll performs Validate then polls until it's completed
func (c TenantConfigurationClient) ValidateThenPoll(ctx context.Context, id ServiceId, input DeployConfigurationParameters) error {
	result, err := c.Validate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Validate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Validate: %+v", err)
	}

	return nil
}

// preparerForValidate prepares the Validate request.
func (c TenantConfigurationClient) preparerForValidate(ctx context.Context, id ServiceId, input DeployConfigurationParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/tenant/configuration/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForValidate sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (c TenantConfigurationClient) senderForValidate(ctx context.Context, req *http.Request) (future ValidateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
