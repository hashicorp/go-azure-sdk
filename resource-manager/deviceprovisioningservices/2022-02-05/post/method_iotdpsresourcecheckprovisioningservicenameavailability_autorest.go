package post

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

type IotDpsResourceCheckProvisioningServiceNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *NameAvailabilityInfo
}

// IotDpsResourceCheckProvisioningServiceNameAvailability ...
func (c POSTClient) IotDpsResourceCheckProvisioningServiceNameAvailability(ctx context.Context, id commonids.SubscriptionId, input OperationInputs) (result IotDpsResourceCheckProvisioningServiceNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceCheckProvisioningServiceNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceCheckProvisioningServiceNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceCheckProvisioningServiceNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotDpsResourceCheckProvisioningServiceNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceCheckProvisioningServiceNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotDpsResourceCheckProvisioningServiceNameAvailability prepares the IotDpsResourceCheckProvisioningServiceNameAvailability request.
func (c POSTClient) preparerForIotDpsResourceCheckProvisioningServiceNameAvailability(ctx context.Context, id commonids.SubscriptionId, input OperationInputs) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Devices/checkProvisioningServiceNameAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotDpsResourceCheckProvisioningServiceNameAvailability handles the response to the IotDpsResourceCheckProvisioningServiceNameAvailability request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForIotDpsResourceCheckProvisioningServiceNameAvailability(resp *http.Response) (result IotDpsResourceCheckProvisioningServiceNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
