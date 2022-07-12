package virtualmachinescalesetvms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceViewOperationResponse struct {
	HttpResponse *http.Response
	Model        *VirtualMachineScaleSetVMInstanceView
}

// GetInstanceView ...
func (c VirtualMachineScaleSetVMsClient) GetInstanceView(ctx context.Context, id VirtualMachineScaleSetVirtualMachineId) (result GetInstanceViewOperationResponse, err error) {
	req, err := c.preparerForGetInstanceView(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "GetInstanceView", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "GetInstanceView", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceView(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "GetInstanceView", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceView prepares the GetInstanceView request.
func (c VirtualMachineScaleSetVMsClient) preparerForGetInstanceView(ctx context.Context, id VirtualMachineScaleSetVirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/instanceView", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInstanceView handles the response to the GetInstanceView request. The method always
// closes the http.Response Body.
func (c VirtualMachineScaleSetVMsClient) responderForGetInstanceView(resp *http.Response) (result GetInstanceViewOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
