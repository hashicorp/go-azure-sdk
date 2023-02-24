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

type GetFeatureValueOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExposureControlResponse
}

// GetFeatureValue ...
func (c ExposureControlClient) GetFeatureValue(ctx context.Context, id LocationId, input ExposureControlRequest) (result GetFeatureValueOperationResponse, err error) {
	req, err := c.preparerForGetFeatureValue(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValue", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValue", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFeatureValue(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "GetFeatureValue", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFeatureValue prepares the GetFeatureValue request.
func (c ExposureControlClient) preparerForGetFeatureValue(ctx context.Context, id LocationId, input ExposureControlRequest) (*http.Request, error) {
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

// responderForGetFeatureValue handles the response to the GetFeatureValue request. The method always
// closes the http.Response Body.
func (c ExposureControlClient) responderForGetFeatureValue(resp *http.Response) (result GetFeatureValueOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
