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

type ListByMachineNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationProfileAssignmentList
}

// ListByMachineName ...
func (c ConfigurationProfileAssignmentsClient) ListByMachineName(ctx context.Context, id MachineId) (result ListByMachineNameOperationResponse, err error) {
	req, err := c.preparerForListByMachineName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByMachineName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByMachineName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByMachineName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByMachineName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByMachineName prepares the ListByMachineName request.
func (c ConfigurationProfileAssignmentsClient) preparerForListByMachineName(ctx context.Context, id MachineId) (*http.Request, error) {
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

// responderForListByMachineName handles the response to the ListByMachineName request. The method always
// closes the http.Response Body.
func (c ConfigurationProfileAssignmentsClient) responderForListByMachineName(resp *http.Response) (result ListByMachineNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
