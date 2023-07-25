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

type CreateOrUpdateSourceControlOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateSourceControl ...
func (c WebAppsClient) CreateOrUpdateSourceControl(ctx context.Context, id SiteId, input SiteSourceControl) (result CreateOrUpdateSourceControlOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateSourceControl(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSourceControl", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateSourceControl(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSourceControl", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateSourceControlThenPoll performs CreateOrUpdateSourceControl then polls until it's completed
func (c WebAppsClient) CreateOrUpdateSourceControlThenPoll(ctx context.Context, id SiteId, input SiteSourceControl) error {
	result, err := c.CreateOrUpdateSourceControl(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateSourceControl: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateSourceControl: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateSourceControl prepares the CreateOrUpdateSourceControl request.
func (c WebAppsClient) preparerForCreateOrUpdateSourceControl(ctx context.Context, id SiteId, input SiteSourceControl) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdateSourceControl sends the CreateOrUpdateSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateOrUpdateSourceControl(ctx context.Context, req *http.Request) (future CreateOrUpdateSourceControlOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
