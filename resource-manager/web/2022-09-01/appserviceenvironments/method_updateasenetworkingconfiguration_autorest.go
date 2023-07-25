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

type UpdateAseNetworkingConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *AseV3NetworkingConfiguration
}

// UpdateAseNetworkingConfiguration ...
func (c AppServiceEnvironmentsClient) UpdateAseNetworkingConfiguration(ctx context.Context, id HostingEnvironmentId, input AseV3NetworkingConfiguration) (result UpdateAseNetworkingConfigurationOperationResponse, err error) {
	req, err := c.preparerForUpdateAseNetworkingConfiguration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseNetworkingConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseNetworkingConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAseNetworkingConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseNetworkingConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAseNetworkingConfiguration prepares the UpdateAseNetworkingConfiguration request.
func (c AppServiceEnvironmentsClient) preparerForUpdateAseNetworkingConfiguration(ctx context.Context, id HostingEnvironmentId, input AseV3NetworkingConfiguration) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configurations/networking", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateAseNetworkingConfiguration handles the response to the UpdateAseNetworkingConfiguration request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForUpdateAseNetworkingConfiguration(resp *http.Response) (result UpdateAseNetworkingConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
