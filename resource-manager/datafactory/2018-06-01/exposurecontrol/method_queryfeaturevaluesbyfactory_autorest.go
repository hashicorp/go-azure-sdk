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

type QueryFeatureValuesByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExposureControlBatchResponse
}

// QueryFeatureValuesByFactory ...
func (c ExposureControlClient) QueryFeatureValuesByFactory(ctx context.Context, id FactoryId, input ExposureControlBatchRequest) (result QueryFeatureValuesByFactoryOperationResponse, err error) {
	req, err := c.preparerForQueryFeatureValuesByFactory(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "QueryFeatureValuesByFactory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "QueryFeatureValuesByFactory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForQueryFeatureValuesByFactory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "exposurecontrol.ExposureControlClient", "QueryFeatureValuesByFactory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForQueryFeatureValuesByFactory prepares the QueryFeatureValuesByFactory request.
func (c ExposureControlClient) preparerForQueryFeatureValuesByFactory(ctx context.Context, id FactoryId, input ExposureControlBatchRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/queryFeaturesValue", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForQueryFeatureValuesByFactory handles the response to the QueryFeatureValuesByFactory request. The method always
// closes the http.Response Body.
func (c ExposureControlClient) responderForQueryFeatureValuesByFactory(resp *http.Response) (result QueryFeatureValuesByFactoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
