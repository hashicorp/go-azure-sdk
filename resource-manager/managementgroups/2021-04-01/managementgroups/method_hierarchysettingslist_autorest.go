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

type HierarchySettingsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *HierarchySettingsList
}

// HierarchySettingsList ...
func (c ManagementGroupsClient) HierarchySettingsList(ctx context.Context, id commonids.ManagementGroupId) (result HierarchySettingsListOperationResponse, err error) {
	req, err := c.preparerForHierarchySettingsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHierarchySettingsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHierarchySettingsList prepares the HierarchySettingsList request.
func (c ManagementGroupsClient) preparerForHierarchySettingsList(ctx context.Context, id commonids.ManagementGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/settings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHierarchySettingsList handles the response to the HierarchySettingsList request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForHierarchySettingsList(resp *http.Response) (result HierarchySettingsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
