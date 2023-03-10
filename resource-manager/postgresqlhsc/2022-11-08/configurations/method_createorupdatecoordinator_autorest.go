package configurations

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

type CreateOrUpdateCoordinatorOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateCoordinator ...
func (c ConfigurationsClient) CreateOrUpdateCoordinator(ctx context.Context, id CoordinatorConfigurationId, input ServerConfiguration) (result CreateOrUpdateCoordinatorOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateCoordinator(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurations.ConfigurationsClient", "CreateOrUpdateCoordinator", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateCoordinator(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurations.ConfigurationsClient", "CreateOrUpdateCoordinator", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateCoordinatorThenPoll performs CreateOrUpdateCoordinator then polls until it's completed
func (c ConfigurationsClient) CreateOrUpdateCoordinatorThenPoll(ctx context.Context, id CoordinatorConfigurationId, input ServerConfiguration) error {
	result, err := c.CreateOrUpdateCoordinator(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateCoordinator: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateCoordinator: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateCoordinator prepares the CreateOrUpdateCoordinator request.
func (c ConfigurationsClient) preparerForCreateOrUpdateCoordinator(ctx context.Context, id CoordinatorConfigurationId, input ServerConfiguration) (*http.Request, error) {
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

// senderForCreateOrUpdateCoordinator sends the CreateOrUpdateCoordinator request. The method will close the
// http.Response Body if it receives an error.
func (c ConfigurationsClient) senderForCreateOrUpdateCoordinator(ctx context.Context, req *http.Request) (future CreateOrUpdateCoordinatorOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
