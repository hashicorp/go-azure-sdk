package exposurecontrol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetFeatureValueByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExposureControlResponse
}

// GetFeatureValueByFactory ...
func (c ExposureControlClient) GetFeatureValueByFactory(ctx context.Context, id FactoryId, input ExposureControlRequest) (result GetFeatureValueByFactoryOperationResponse, err error) {
	req, err := c.preparerForGetFeatureValueByFactory(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValueByFactory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValueByFactory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFeatureValueByFactory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValueByFactory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFeatureValueByFactory prepares the GetFeatureValueByFactory request.
func (c ExposureControlClient) preparerForGetFeatureValueByFactory(ctx context.Context, id FactoryId, input ExposureControlRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getFeatureValue", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFeatureValueByFactory handles the response to the GetFeatureValueByFactory request. The method always
// closes the http.Response Body.
func (c ExposureControlClient) responderForGetFeatureValueByFactory(resp *http.Response) (result GetFeatureValueByFactoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
