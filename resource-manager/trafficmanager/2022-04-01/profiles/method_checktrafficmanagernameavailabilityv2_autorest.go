package profiles

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

type CheckTrafficManagerNameAvailabilityV2OperationResponse struct {
	HttpResponse *http.Response
	Model        *TrafficManagerNameAvailability
}

// CheckTrafficManagerNameAvailabilityV2 ...
func (c ProfilesClient) CheckTrafficManagerNameAvailabilityV2(ctx context.Context, id commonids.SubscriptionId, input CheckTrafficManagerRelativeDnsNameAvailabilityParameters) (result CheckTrafficManagerNameAvailabilityV2OperationResponse, err error) {
	req, err := c.preparerForCheckTrafficManagerNameAvailabilityV2(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "profiles.ProfilesClient", "CheckTrafficManagerNameAvailabilityV2", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "profiles.ProfilesClient", "CheckTrafficManagerNameAvailabilityV2", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckTrafficManagerNameAvailabilityV2(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "profiles.ProfilesClient", "CheckTrafficManagerNameAvailabilityV2", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckTrafficManagerNameAvailabilityV2 prepares the CheckTrafficManagerNameAvailabilityV2 request.
func (c ProfilesClient) preparerForCheckTrafficManagerNameAvailabilityV2(ctx context.Context, id commonids.SubscriptionId, input CheckTrafficManagerRelativeDnsNameAvailabilityParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Network/checkTrafficManagerNameAvailabilityV2", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckTrafficManagerNameAvailabilityV2 handles the response to the CheckTrafficManagerNameAvailabilityV2 request. The method always
// closes the http.Response Body.
func (c ProfilesClient) responderForCheckTrafficManagerNameAvailabilityV2(resp *http.Response) (result CheckTrafficManagerNameAvailabilityV2OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
