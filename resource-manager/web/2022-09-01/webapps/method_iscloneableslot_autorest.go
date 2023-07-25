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

type IsCloneableSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteCloneability
}

// IsCloneableSlot ...
func (c WebAppsClient) IsCloneableSlot(ctx context.Context, id SlotId) (result IsCloneableSlotOperationResponse, err error) {
	req, err := c.preparerForIsCloneableSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneableSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneableSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIsCloneableSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneableSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIsCloneableSlot prepares the IsCloneableSlot request.
func (c WebAppsClient) preparerForIsCloneableSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/iscloneable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIsCloneableSlot handles the response to the IsCloneableSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForIsCloneableSlot(resp *http.Response) (result IsCloneableSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
