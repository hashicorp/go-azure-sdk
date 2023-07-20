package applyupdates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateParentOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplyUpdate
}

// CreateOrUpdateParent ...
func (c ApplyUpdatesClient) CreateOrUpdateParent(ctx context.Context, id commonids.ScopeId) (result CreateOrUpdateParentOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateParent(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "CreateOrUpdateParent", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "CreateOrUpdateParent", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateParent(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applyupdates.ApplyUpdatesClient", "CreateOrUpdateParent", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateParent prepares the CreateOrUpdateParent request.
func (c ApplyUpdatesClient) preparerForCreateOrUpdateParent(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Maintenance/applyUpdates/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateParent handles the response to the CreateOrUpdateParent request. The method always
// closes the http.Response Body.
func (c ApplyUpdatesClient) responderForCreateOrUpdateParent(resp *http.Response) (result CreateOrUpdateParentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
