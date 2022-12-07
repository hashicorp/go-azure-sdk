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

type ServiceRegistriesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServiceRegistriesCreateOrUpdate ...
func (c AppPlatformClient) ServiceRegistriesCreateOrUpdate(ctx context.Context, id ServiceRegistryId) (result ServiceRegistriesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForServiceRegistriesCreateOrUpdate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServiceRegistriesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServiceRegistriesCreateOrUpdateThenPoll performs ServiceRegistriesCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) ServiceRegistriesCreateOrUpdateThenPoll(ctx context.Context, id ServiceRegistryId) error {
	result, err := c.ServiceRegistriesCreateOrUpdate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServiceRegistriesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServiceRegistriesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForServiceRegistriesCreateOrUpdate prepares the ServiceRegistriesCreateOrUpdate request.
func (c AppPlatformClient) preparerForServiceRegistriesCreateOrUpdate(ctx context.Context, id ServiceRegistryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServiceRegistriesCreateOrUpdate sends the ServiceRegistriesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForServiceRegistriesCreateOrUpdate(ctx context.Context, req *http.Request) (future ServiceRegistriesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
