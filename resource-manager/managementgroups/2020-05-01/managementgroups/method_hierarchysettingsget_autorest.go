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

type HierarchySettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *HierarchySettings
}

// HierarchySettingsGet ...
func (c ManagementGroupsClient) HierarchySettingsGet(ctx context.Context, id commonids.ManagementGroupId) (result HierarchySettingsGetOperationResponse, err error) {
	req, err := c.preparerForHierarchySettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHierarchySettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHierarchySettingsGet prepares the HierarchySettingsGet request.
func (c ManagementGroupsClient) preparerForHierarchySettingsGet(ctx context.Context, id commonids.ManagementGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/settings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHierarchySettingsGet handles the response to the HierarchySettingsGet request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForHierarchySettingsGet(resp *http.Response) (result HierarchySettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
