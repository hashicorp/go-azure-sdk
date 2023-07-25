package staticsites

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

type CreateOrUpdateStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateStaticSite ...
func (c StaticSitesClient) CreateOrUpdateStaticSite(ctx context.Context, id StaticSiteId, input StaticSiteARMResource) (result CreateOrUpdateStaticSiteOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateStaticSiteThenPoll performs CreateOrUpdateStaticSite then polls until it's completed
func (c StaticSitesClient) CreateOrUpdateStaticSiteThenPoll(ctx context.Context, id StaticSiteId, input StaticSiteARMResource) error {
	result, err := c.CreateOrUpdateStaticSite(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateStaticSite: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateStaticSite prepares the CreateOrUpdateStaticSite request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSite(ctx context.Context, id StaticSiteId, input StaticSiteARMResource) (*http.Request, error) {
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

// senderForCreateOrUpdateStaticSite sends the CreateOrUpdateStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForCreateOrUpdateStaticSite(ctx context.Context, req *http.Request) (future CreateOrUpdateStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
