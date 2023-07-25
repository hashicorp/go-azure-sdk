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

type UpdateSourceControlSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteSourceControl
}

// UpdateSourceControlSlot ...
func (c WebAppsClient) UpdateSourceControlSlot(ctx context.Context, id SlotId, input SiteSourceControl) (result UpdateSourceControlSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateSourceControlSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSourceControlSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSourceControlSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSourceControlSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSourceControlSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSourceControlSlot prepares the UpdateSourceControlSlot request.
func (c WebAppsClient) preparerForUpdateSourceControlSlot(ctx context.Context, id SlotId, input SiteSourceControl) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSourceControlSlot handles the response to the UpdateSourceControlSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateSourceControlSlot(resp *http.Response) (result UpdateSourceControlSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
