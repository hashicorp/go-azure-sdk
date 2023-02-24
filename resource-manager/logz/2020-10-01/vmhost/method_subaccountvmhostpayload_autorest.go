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

type SubAccountVMHostPayloadOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMExtensionPayload
}

// SubAccountVMHostPayload ...
func (c VMHostClient) SubAccountVMHostPayload(ctx context.Context, id AccountId) (result SubAccountVMHostPayloadOperationResponse, err error) {
	req, err := c.preparerForSubAccountVMHostPayload(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountVMHostPayload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountVMHostPayload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSubAccountVMHostPayload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountVMHostPayload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSubAccountVMHostPayload prepares the SubAccountVMHostPayload request.
func (c VMHostClient) preparerForSubAccountVMHostPayload(ctx context.Context, id AccountId) (*http.Request, error) {
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

// responderForSubAccountVMHostPayload handles the response to the SubAccountVMHostPayload request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForSubAccountVMHostPayload(resp *http.Response) (result SubAccountVMHostPayloadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
