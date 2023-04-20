package configurations

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

type CreateInResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigData
}

// CreateInResourceGroup ...
func (c ConfigurationsClient) CreateInResourceGroup(ctx context.Context, id commonids.ResourceGroupId, input ConfigData) (result CreateInResourceGroupOperationResponse, err error) {
	req, err := c.preparerForCreateInResourceGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurations.ConfigurationsClient", "CreateInResourceGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurations.ConfigurationsClient", "CreateInResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateInResourceGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurations.ConfigurationsClient", "CreateInResourceGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateInResourceGroup prepares the CreateInResourceGroup request.
func (c ConfigurationsClient) preparerForCreateInResourceGroup(ctx context.Context, id commonids.ResourceGroupId, input ConfigData) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Advisor/configurations/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateInResourceGroup handles the response to the CreateInResourceGroup request. The method always
// closes the http.Response Body.
func (c ConfigurationsClient) responderForCreateInResourceGroup(resp *http.Response) (result CreateInResourceGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
