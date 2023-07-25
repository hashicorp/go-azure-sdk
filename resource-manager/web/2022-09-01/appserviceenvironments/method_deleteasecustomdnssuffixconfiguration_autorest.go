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

type DeleteAseCustomDnsSuffixConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// DeleteAseCustomDnsSuffixConfiguration ...
func (c AppServiceEnvironmentsClient) DeleteAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId) (result DeleteAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	req, err := c.preparerForDeleteAseCustomDnsSuffixConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "DeleteAseCustomDnsSuffixConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "DeleteAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteAseCustomDnsSuffixConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "DeleteAseCustomDnsSuffixConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteAseCustomDnsSuffixConfiguration prepares the DeleteAseCustomDnsSuffixConfiguration request.
func (c AppServiceEnvironmentsClient) preparerForDeleteAseCustomDnsSuffixConfiguration(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configurations/customdnssuffix", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteAseCustomDnsSuffixConfiguration handles the response to the DeleteAseCustomDnsSuffixConfiguration request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForDeleteAseCustomDnsSuffixConfiguration(resp *http.Response) (result DeleteAseCustomDnsSuffixConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
