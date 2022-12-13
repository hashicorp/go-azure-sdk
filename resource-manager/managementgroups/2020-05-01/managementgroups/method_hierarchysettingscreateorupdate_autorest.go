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

type HierarchySettingsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *HierarchySettings
}

// HierarchySettingsCreateOrUpdate ...
func (c ManagementGroupsClient) HierarchySettingsCreateOrUpdate(ctx context.Context, id commonids.ManagementGroupId, input CreateOrUpdateSettingsRequest) (result HierarchySettingsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForHierarchySettingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHierarchySettingsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHierarchySettingsCreateOrUpdate prepares the HierarchySettingsCreateOrUpdate request.
func (c ManagementGroupsClient) preparerForHierarchySettingsCreateOrUpdate(ctx context.Context, id commonids.ManagementGroupId, input CreateOrUpdateSettingsRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/settings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHierarchySettingsCreateOrUpdate handles the response to the HierarchySettingsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForHierarchySettingsCreateOrUpdate(resp *http.Response) (result HierarchySettingsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
