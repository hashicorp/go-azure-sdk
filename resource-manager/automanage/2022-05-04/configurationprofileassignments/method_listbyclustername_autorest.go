package configurationprofileassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByClusterNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationProfileAssignmentList
}

// ListByClusterName ...
func (c ConfigurationProfileAssignmentsClient) ListByClusterName(ctx context.Context, id ClusterId) (result ListByClusterNameOperationResponse, err error) {
	req, err := c.preparerForListByClusterName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByClusterName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByClusterName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByClusterName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByClusterName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByClusterName prepares the ListByClusterName request.
func (c ConfigurationProfileAssignmentsClient) preparerForListByClusterName(ctx context.Context, id ClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.AutoManage/configurationProfileAssignments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByClusterName handles the response to the ListByClusterName request. The method always
// closes the http.Response Body.
func (c ConfigurationProfileAssignmentsClient) responderForListByClusterName(resp *http.Response) (result ListByClusterNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
