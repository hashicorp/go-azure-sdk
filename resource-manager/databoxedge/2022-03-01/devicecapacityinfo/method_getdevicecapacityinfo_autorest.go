package devicecapacityinfo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceCapacityInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeviceCapacityInfo
}

// GetDeviceCapacityInfo ...
func (c DeviceCapacityInfoClient) GetDeviceCapacityInfo(ctx context.Context, id DataBoxEdgeDeviceId) (result GetDeviceCapacityInfoOperationResponse, err error) {
	req, err := c.preparerForGetDeviceCapacityInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devicecapacityinfo.DeviceCapacityInfoClient", "GetDeviceCapacityInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "devicecapacityinfo.DeviceCapacityInfoClient", "GetDeviceCapacityInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeviceCapacityInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devicecapacityinfo.DeviceCapacityInfoClient", "GetDeviceCapacityInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeviceCapacityInfo prepares the GetDeviceCapacityInfo request.
func (c DeviceCapacityInfoClient) preparerForGetDeviceCapacityInfo(ctx context.Context, id DataBoxEdgeDeviceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deviceCapacityInfo/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDeviceCapacityInfo handles the response to the GetDeviceCapacityInfo request. The method always
// closes the http.Response Body.
func (c DeviceCapacityInfoClient) responderForGetDeviceCapacityInfo(resp *http.Response) (result GetDeviceCapacityInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
