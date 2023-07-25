package appserviceenvironments

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

type CreateOrUpdateMultiRolePoolOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateMultiRolePool ...
func (c AppServiceEnvironmentsClient) CreateOrUpdateMultiRolePool(ctx context.Context, id HostingEnvironmentId, input WorkerPoolResource) (result CreateOrUpdateMultiRolePoolOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateMultiRolePool(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "CreateOrUpdateMultiRolePool", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateMultiRolePool(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "CreateOrUpdateMultiRolePool", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateMultiRolePoolThenPoll performs CreateOrUpdateMultiRolePool then polls until it's completed
func (c AppServiceEnvironmentsClient) CreateOrUpdateMultiRolePoolThenPoll(ctx context.Context, id HostingEnvironmentId, input WorkerPoolResource) error {
	result, err := c.CreateOrUpdateMultiRolePool(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateMultiRolePool: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateMultiRolePool: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateMultiRolePool prepares the CreateOrUpdateMultiRolePool request.
func (c AppServiceEnvironmentsClient) preparerForCreateOrUpdateMultiRolePool(ctx context.Context, id HostingEnvironmentId, input WorkerPoolResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdateMultiRolePool sends the CreateOrUpdateMultiRolePool request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceEnvironmentsClient) senderForCreateOrUpdateMultiRolePool(ctx context.Context, req *http.Request) (future CreateOrUpdateMultiRolePoolOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
