package resourceproviders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifyHostingEnvironmentVnetOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetValidationFailureDetails
}

// VerifyHostingEnvironmentVnet ...
func (c ResourceProvidersClient) VerifyHostingEnvironmentVnet(ctx context.Context, id commonids.SubscriptionId, input VnetParameters) (result VerifyHostingEnvironmentVnetOperationResponse, err error) {
	req, err := c.preparerForVerifyHostingEnvironmentVnet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "VerifyHostingEnvironmentVnet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "VerifyHostingEnvironmentVnet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForVerifyHostingEnvironmentVnet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "VerifyHostingEnvironmentVnet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForVerifyHostingEnvironmentVnet prepares the VerifyHostingEnvironmentVnet request.
func (c ResourceProvidersClient) preparerForVerifyHostingEnvironmentVnet(ctx context.Context, id commonids.SubscriptionId, input VnetParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/verifyHostingEnvironmentVnet", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForVerifyHostingEnvironmentVnet handles the response to the VerifyHostingEnvironmentVnet request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForVerifyHostingEnvironmentVnet(resp *http.Response) (result VerifyHostingEnvironmentVnetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
