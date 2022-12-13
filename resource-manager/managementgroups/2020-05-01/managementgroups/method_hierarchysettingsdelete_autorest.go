package managementgroups

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

type HierarchySettingsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// HierarchySettingsDelete ...
func (c ManagementGroupsClient) HierarchySettingsDelete(ctx context.Context, id commonids.ManagementGroupId) (result HierarchySettingsDeleteOperationResponse, err error) {
	req, err := c.preparerForHierarchySettingsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHierarchySettingsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHierarchySettingsDelete prepares the HierarchySettingsDelete request.
func (c ManagementGroupsClient) preparerForHierarchySettingsDelete(ctx context.Context, id commonids.ManagementGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/settings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHierarchySettingsDelete handles the response to the HierarchySettingsDelete request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForHierarchySettingsDelete(resp *http.Response) (result HierarchySettingsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
