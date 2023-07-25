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

type CreateZipDeploymentForStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateZipDeploymentForStaticSite ...
func (c StaticSitesClient) CreateZipDeploymentForStaticSite(ctx context.Context, id StaticSiteId, input StaticSiteZipDeploymentARMResource) (result CreateZipDeploymentForStaticSiteOperationResponse, err error) {
	req, err := c.preparerForCreateZipDeploymentForStaticSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateZipDeploymentForStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateZipDeploymentForStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateZipDeploymentForStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateZipDeploymentForStaticSiteThenPoll performs CreateZipDeploymentForStaticSite then polls until it's completed
func (c StaticSitesClient) CreateZipDeploymentForStaticSiteThenPoll(ctx context.Context, id StaticSiteId, input StaticSiteZipDeploymentARMResource) error {
	result, err := c.CreateZipDeploymentForStaticSite(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateZipDeploymentForStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateZipDeploymentForStaticSite: %+v", err)
	}

	return nil
}

// preparerForCreateZipDeploymentForStaticSite prepares the CreateZipDeploymentForStaticSite request.
func (c StaticSitesClient) preparerForCreateZipDeploymentForStaticSite(ctx context.Context, id StaticSiteId, input StaticSiteZipDeploymentARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/zipdeploy", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateZipDeploymentForStaticSite sends the CreateZipDeploymentForStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForCreateZipDeploymentForStaticSite(ctx context.Context, req *http.Request) (future CreateZipDeploymentForStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
