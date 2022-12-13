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

type HierarchySettingsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *HierarchySettings
}

// HierarchySettingsUpdate ...
func (c ManagementGroupsClient) HierarchySettingsUpdate(ctx context.Context, id commonids.ManagementGroupId, input CreateOrUpdateSettingsRequest) (result HierarchySettingsUpdateOperationResponse, err error) {
	req, err := c.preparerForHierarchySettingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHierarchySettingsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "HierarchySettingsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHierarchySettingsUpdate prepares the HierarchySettingsUpdate request.
func (c ManagementGroupsClient) preparerForHierarchySettingsUpdate(ctx context.Context, id commonids.ManagementGroupId, input CreateOrUpdateSettingsRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/settings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHierarchySettingsUpdate handles the response to the HierarchySettingsUpdate request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForHierarchySettingsUpdate(resp *http.Response) (result HierarchySettingsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
