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

type ServiceRegistriesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServiceRegistriesDelete ...
func (c AppPlatformClient) ServiceRegistriesDelete(ctx context.Context, id ServiceRegistryId) (result ServiceRegistriesDeleteOperationResponse, err error) {
	req, err := c.preparerForServiceRegistriesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServiceRegistriesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServiceRegistriesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServiceRegistriesDeleteThenPoll performs ServiceRegistriesDelete then polls until it's completed
func (c AppPlatformClient) ServiceRegistriesDeleteThenPoll(ctx context.Context, id ServiceRegistryId) error {
	result, err := c.ServiceRegistriesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServiceRegistriesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServiceRegistriesDelete: %+v", err)
	}

	return nil
}

// preparerForServiceRegistriesDelete prepares the ServiceRegistriesDelete request.
func (c AppPlatformClient) preparerForServiceRegistriesDelete(ctx context.Context, id ServiceRegistryId) (*http.Request, error) {
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

// senderForServiceRegistriesDelete sends the ServiceRegistriesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForServiceRegistriesDelete(ctx context.Context, req *http.Request) (future ServiceRegistriesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
