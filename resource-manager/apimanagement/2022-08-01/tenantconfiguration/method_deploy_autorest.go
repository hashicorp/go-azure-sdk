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

type DeployOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Deploy ...
func (c TenantConfigurationClient) Deploy(ctx context.Context, id ServiceId, input DeployConfigurationParameters) (result DeployOperationResponse, err error) {
	req, err := c.preparerForDeploy(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Deploy", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploy(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Deploy", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeployThenPoll performs Deploy then polls until it's completed
func (c TenantConfigurationClient) DeployThenPoll(ctx context.Context, id ServiceId, input DeployConfigurationParameters) error {
	result, err := c.Deploy(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Deploy: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Deploy: %+v", err)
	}

	return nil
}

// preparerForDeploy prepares the Deploy request.
func (c TenantConfigurationClient) preparerForDeploy(ctx context.Context, id ServiceId, input DeployConfigurationParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/tenant/configuration/deploy", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploy sends the Deploy request. The method will close the
// http.Response Body if it receives an error.
func (c TenantConfigurationClient) senderForDeploy(ctx context.Context, req *http.Request) (future DeployOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
