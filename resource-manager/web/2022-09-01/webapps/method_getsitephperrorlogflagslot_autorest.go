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

type GetSitePhpErrorLogFlagSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SitePhpErrorLogFlag
}

// GetSitePhpErrorLogFlagSlot ...
func (c WebAppsClient) GetSitePhpErrorLogFlagSlot(ctx context.Context, id SlotId) (result GetSitePhpErrorLogFlagSlotOperationResponse, err error) {
	req, err := c.preparerForGetSitePhpErrorLogFlagSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSitePhpErrorLogFlagSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSitePhpErrorLogFlagSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSitePhpErrorLogFlagSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSitePhpErrorLogFlagSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSitePhpErrorLogFlagSlot prepares the GetSitePhpErrorLogFlagSlot request.
func (c WebAppsClient) preparerForGetSitePhpErrorLogFlagSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/phplogging", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSitePhpErrorLogFlagSlot handles the response to the GetSitePhpErrorLogFlagSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSitePhpErrorLogFlagSlot(resp *http.Response) (result GetSitePhpErrorLogFlagSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
