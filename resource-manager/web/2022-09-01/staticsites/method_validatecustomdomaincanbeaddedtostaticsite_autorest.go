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

type ValidateCustomDomainCanBeAddedToStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateCustomDomainCanBeAddedToStaticSite ...
func (c StaticSitesClient) ValidateCustomDomainCanBeAddedToStaticSite(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) (result ValidateCustomDomainCanBeAddedToStaticSiteOperationResponse, err error) {
	req, err := c.preparerForValidateCustomDomainCanBeAddedToStaticSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateCustomDomainCanBeAddedToStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateCustomDomainCanBeAddedToStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateCustomDomainCanBeAddedToStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateCustomDomainCanBeAddedToStaticSiteThenPoll performs ValidateCustomDomainCanBeAddedToStaticSite then polls until it's completed
func (c StaticSitesClient) ValidateCustomDomainCanBeAddedToStaticSiteThenPoll(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) error {
	result, err := c.ValidateCustomDomainCanBeAddedToStaticSite(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateCustomDomainCanBeAddedToStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateCustomDomainCanBeAddedToStaticSite: %+v", err)
	}

	return nil
}

// preparerForValidateCustomDomainCanBeAddedToStaticSite prepares the ValidateCustomDomainCanBeAddedToStaticSite request.
func (c StaticSitesClient) preparerForValidateCustomDomainCanBeAddedToStaticSite(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForValidateCustomDomainCanBeAddedToStaticSite sends the ValidateCustomDomainCanBeAddedToStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForValidateCustomDomainCanBeAddedToStaticSite(ctx context.Context, req *http.Request) (future ValidateCustomDomainCanBeAddedToStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
