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

type SaveOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Save ...
func (c TenantConfigurationClient) Save(ctx context.Context, id ServiceId, input SaveConfigurationParameter) (result SaveOperationResponse, err error) {
	req, err := c.preparerForSave(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Save", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSave(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfiguration.TenantConfigurationClient", "Save", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SaveThenPoll performs Save then polls until it's completed
func (c TenantConfigurationClient) SaveThenPoll(ctx context.Context, id ServiceId, input SaveConfigurationParameter) error {
	result, err := c.Save(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Save: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Save: %+v", err)
	}

	return nil
}

// preparerForSave prepares the Save request.
func (c TenantConfigurationClient) preparerForSave(ctx context.Context, id ServiceId, input SaveConfigurationParameter) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/tenant/configuration/save", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSave sends the Save request. The method will close the
// http.Response Body if it receives an error.
func (c TenantConfigurationClient) senderForSave(ctx context.Context, req *http.Request) (future SaveOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
