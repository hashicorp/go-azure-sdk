package vmhost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorVMHostPayloadOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMExtensionPayload
}

// MonitorVMHostPayload ...
func (c VMHostClient) MonitorVMHostPayload(ctx context.Context, id MonitorId) (result MonitorVMHostPayloadOperationResponse, err error) {
	req, err := c.preparerForMonitorVMHostPayload(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorVMHostPayload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorVMHostPayload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMonitorVMHostPayload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorVMHostPayload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMonitorVMHostPayload prepares the MonitorVMHostPayload request.
func (c VMHostClient) preparerForMonitorVMHostPayload(ctx context.Context, id MonitorId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/vmHostPayload", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForMonitorVMHostPayload handles the response to the MonitorVMHostPayload request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForMonitorVMHostPayload(resp *http.Response) (result MonitorVMHostPayloadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
