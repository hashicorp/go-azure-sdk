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

type DeleteStaticSiteBuildOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteStaticSiteBuild ...
func (c StaticSitesClient) DeleteStaticSiteBuild(ctx context.Context, id BuildId) (result DeleteStaticSiteBuildOperationResponse, err error) {
	req, err := c.preparerForDeleteStaticSiteBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteBuild", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteStaticSiteBuild(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteBuild", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteStaticSiteBuildThenPoll performs DeleteStaticSiteBuild then polls until it's completed
func (c StaticSitesClient) DeleteStaticSiteBuildThenPoll(ctx context.Context, id BuildId) error {
	result, err := c.DeleteStaticSiteBuild(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteStaticSiteBuild: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteStaticSiteBuild: %+v", err)
	}

	return nil
}

// preparerForDeleteStaticSiteBuild prepares the DeleteStaticSiteBuild request.
func (c StaticSitesClient) preparerForDeleteStaticSiteBuild(ctx context.Context, id BuildId) (*http.Request, error) {
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

// senderForDeleteStaticSiteBuild sends the DeleteStaticSiteBuild request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForDeleteStaticSiteBuild(ctx context.Context, req *http.Request) (future DeleteStaticSiteBuildOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
