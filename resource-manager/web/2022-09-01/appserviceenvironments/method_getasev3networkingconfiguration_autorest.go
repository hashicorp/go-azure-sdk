package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAseV3NetworkingConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *AseV3NetworkingConfiguration
}

// GetAseV3NetworkingConfiguration ...
func (c AppServiceEnvironmentsClient) GetAseV3NetworkingConfiguration(ctx context.Context, id HostingEnvironmentId) (result GetAseV3NetworkingConfigurationOperationResponse, err error) {
	req, err := c.preparerForGetAseV3NetworkingConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseV3NetworkingConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseV3NetworkingConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAseV3NetworkingConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseV3NetworkingConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAseV3NetworkingConfiguration prepares the GetAseV3NetworkingConfiguration request.
func (c AppServiceEnvironmentsClient) preparerForGetAseV3NetworkingConfiguration(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configurations/networking", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAseV3NetworkingConfiguration handles the response to the GetAseV3NetworkingConfiguration request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetAseV3NetworkingConfiguration(resp *http.Response) (result GetAseV3NetworkingConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
