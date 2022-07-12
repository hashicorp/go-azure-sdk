package virtualmachinescalesetrollingupgrades

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLatestOperationResponse struct {
	HttpResponse *http.Response
	Model        *RollingUpgradeStatusInfo
}

// GetLatest ...
func (c VirtualMachineScaleSetRollingUpgradesClient) GetLatest(ctx context.Context, id VirtualMachineScaleSetId) (result GetLatestOperationResponse, err error) {
	req, err := c.preparerForGetLatest(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "GetLatest", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "GetLatest", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetLatest(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "GetLatest", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetLatest prepares the GetLatest request.
func (c VirtualMachineScaleSetRollingUpgradesClient) preparerForGetLatest(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/rollingUpgrades/latest", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetLatest handles the response to the GetLatest request. The method always
// closes the http.Response Body.
func (c VirtualMachineScaleSetRollingUpgradesClient) responderForGetLatest(resp *http.Response) (result GetLatestOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
