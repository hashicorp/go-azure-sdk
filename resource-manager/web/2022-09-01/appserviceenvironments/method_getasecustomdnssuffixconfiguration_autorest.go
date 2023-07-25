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

type GetAseCustomDnsSuffixConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomDnsSuffixConfiguration
}

// GetAseCustomDnsSuffixConfiguration ...
func (c AppServiceEnvironmentsClient) GetAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId) (result GetAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	req, err := c.preparerForGetAseCustomDnsSuffixConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseCustomDnsSuffixConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAseCustomDnsSuffixConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAseCustomDnsSuffixConfiguration prepares the GetAseCustomDnsSuffixConfiguration request.
func (c AppServiceEnvironmentsClient) preparerForGetAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configurations/customdnssuffix", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAseCustomDnsSuffixConfiguration handles the response to the GetAseCustomDnsSuffixConfiguration request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetAseCustomDnsSuffixConfiguration(resp *http.Response) (result GetAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
