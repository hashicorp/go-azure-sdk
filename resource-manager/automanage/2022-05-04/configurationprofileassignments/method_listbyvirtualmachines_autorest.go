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

type ListByVirtualMachinesOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationProfileAssignmentList
}

// ListByVirtualMachines ...
func (c ConfigurationProfileAssignmentsClient) ListByVirtualMachines(ctx context.Context, id VirtualMachineId) (result ListByVirtualMachinesOperationResponse, err error) {
	req, err := c.preparerForListByVirtualMachines(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByVirtualMachines", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByVirtualMachines", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByVirtualMachines(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationprofileassignments.ConfigurationProfileAssignmentsClient", "ListByVirtualMachines", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByVirtualMachines prepares the ListByVirtualMachines request.
func (c ConfigurationProfileAssignmentsClient) preparerForListByVirtualMachines(ctx context.Context, id VirtualMachineId) (*http.Request, error) {
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

// responderForListByVirtualMachines handles the response to the ListByVirtualMachines request. The method always
// closes the http.Response Body.
func (c ConfigurationProfileAssignmentsClient) responderForListByVirtualMachines(resp *http.Response) (result ListByVirtualMachinesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
