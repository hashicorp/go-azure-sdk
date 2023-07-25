package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetProductionSlotConfigOperationResponse struct {
	HttpResponse *http.Response
}

// ResetProductionSlotConfig ...
func (c WebAppsClient) ResetProductionSlotConfig(ctx context.Context, id SiteId) (result ResetProductionSlotConfigOperationResponse, err error) {
	req, err := c.preparerForResetProductionSlotConfig(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetProductionSlotConfig", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetProductionSlotConfig", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetProductionSlotConfig(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetProductionSlotConfig", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetProductionSlotConfig prepares the ResetProductionSlotConfig request.
func (c WebAppsClient) preparerForResetProductionSlotConfig(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resetSlotConfig", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetProductionSlotConfig handles the response to the ResetProductionSlotConfig request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForResetProductionSlotConfig(resp *http.Response) (result ResetProductionSlotConfigOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
