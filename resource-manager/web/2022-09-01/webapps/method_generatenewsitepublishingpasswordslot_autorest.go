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

type GenerateNewSitePublishingPasswordSlotOperationResponse struct {
	HttpResponse *http.Response
}

// GenerateNewSitePublishingPasswordSlot ...
func (c WebAppsClient) GenerateNewSitePublishingPasswordSlot(ctx context.Context, id SlotId) (result GenerateNewSitePublishingPasswordSlotOperationResponse, err error) {
	req, err := c.preparerForGenerateNewSitePublishingPasswordSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPasswordSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPasswordSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGenerateNewSitePublishingPasswordSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPasswordSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGenerateNewSitePublishingPasswordSlot prepares the GenerateNewSitePublishingPasswordSlot request.
func (c WebAppsClient) preparerForGenerateNewSitePublishingPasswordSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/newpassword", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGenerateNewSitePublishingPasswordSlot handles the response to the GenerateNewSitePublishingPasswordSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGenerateNewSitePublishingPasswordSlot(resp *http.Response) (result GenerateNewSitePublishingPasswordSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
