package webapps

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

type CreateInstanceMSDeployOperationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateInstanceMSDeployOperation ...
func (c WebAppsClient) CreateInstanceMSDeployOperation(ctx context.Context, id InstanceId, input MSDeploy) (result CreateInstanceMSDeployOperationOperationResponse, err error) {
	req, err := c.preparerForCreateInstanceMSDeployOperation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceMSDeployOperation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateInstanceMSDeployOperation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateInstanceMSDeployOperation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateInstanceMSDeployOperationThenPoll performs CreateInstanceMSDeployOperation then polls until it's completed
func (c WebAppsClient) CreateInstanceMSDeployOperationThenPoll(ctx context.Context, id InstanceId, input MSDeploy) error {
	result, err := c.CreateInstanceMSDeployOperation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateInstanceMSDeployOperation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateInstanceMSDeployOperation: %+v", err)
	}

	return nil
}

// preparerForCreateInstanceMSDeployOperation prepares the CreateInstanceMSDeployOperation request.
func (c WebAppsClient) preparerForCreateInstanceMSDeployOperation(ctx context.Context, id InstanceId, input MSDeploy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/mSDeploy", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateInstanceMSDeployOperation sends the CreateInstanceMSDeployOperation request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateInstanceMSDeployOperation(ctx context.Context, req *http.Request) (future CreateInstanceMSDeployOperationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
