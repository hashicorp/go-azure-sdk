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

type UpdateAseCustomDnsSuffixConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomDnsSuffixConfiguration
}

// UpdateAseCustomDnsSuffixConfiguration ...
func (c AppServiceEnvironmentsClient) UpdateAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId, input CustomDnsSuffixConfiguration) (result UpdateAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	req, err := c.preparerForUpdateAseCustomDnsSuffixConfiguration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseCustomDnsSuffixConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAseCustomDnsSuffixConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAseCustomDnsSuffixConfiguration prepares the UpdateAseCustomDnsSuffixConfiguration request.
func (c AppServiceEnvironmentsClient) preparerForUpdateAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId, input CustomDnsSuffixConfiguration) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configurations/customdnssuffix", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateAseCustomDnsSuffixConfiguration handles the response to the UpdateAseCustomDnsSuffixConfiguration request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForUpdateAseCustomDnsSuffixConfiguration(resp *http.Response) (result UpdateAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
