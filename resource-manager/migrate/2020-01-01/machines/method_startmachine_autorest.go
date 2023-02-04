package machines

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

type StartMachineOperationResponse struct {
	HttpResponse *http.Response
}

// StartMachine ...
func (c MachinesClient) StartMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (result StartMachineOperationResponse, err error) {
	req, err := c.preparerForStartMachine(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StartMachine", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StartMachine", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartMachine(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StartMachine", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartMachine prepares the StartMachine request.
func (c MachinesClient) preparerForStartMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStartMachine handles the response to the StartMachine request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForStartMachine(resp *http.Response) (result StartMachineOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
