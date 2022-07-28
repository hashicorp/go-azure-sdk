package snapshotpolicylistvolumes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotPoliciesListVolumesOperationResponse struct {
	HttpResponse *http.Response
	Model        *SnapshotPolicyVolumeList
}

// SnapshotPoliciesListVolumes ...
func (c SnapshotPolicyListVolumesClient) SnapshotPoliciesListVolumes(ctx context.Context, id SnapshotPoliciesId) (result SnapshotPoliciesListVolumesOperationResponse, err error) {
	req, err := c.preparerForSnapshotPoliciesListVolumes(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient", "SnapshotPoliciesListVolumes", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient", "SnapshotPoliciesListVolumes", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSnapshotPoliciesListVolumes(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient", "SnapshotPoliciesListVolumes", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSnapshotPoliciesListVolumes prepares the SnapshotPoliciesListVolumes request.
func (c SnapshotPolicyListVolumesClient) preparerForSnapshotPoliciesListVolumes(ctx context.Context, id SnapshotPoliciesId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/volumes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSnapshotPoliciesListVolumes handles the response to the SnapshotPoliciesListVolumes request. The method always
// closes the http.Response Body.
func (c SnapshotPolicyListVolumesClient) responderForSnapshotPoliciesListVolumes(resp *http.Response) (result SnapshotPoliciesListVolumesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
