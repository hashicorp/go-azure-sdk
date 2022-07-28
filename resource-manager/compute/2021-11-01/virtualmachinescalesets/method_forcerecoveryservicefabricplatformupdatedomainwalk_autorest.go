package virtualmachinescalesets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecoveryWalkResponse
}

type ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions struct {
	PlacementGroupId     *string
	PlatformUpdateDomain *int64
	Zone                 *string
}

func DefaultForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions() ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions {
	return ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions{}
}

func (o ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.PlacementGroupId != nil {
		out["placementGroupId"] = *o.PlacementGroupId
	}

	if o.PlatformUpdateDomain != nil {
		out["platformUpdateDomain"] = *o.PlatformUpdateDomain
	}

	if o.Zone != nil {
		out["zone"] = *o.Zone
	}

	return out
}

// ForceRecoveryServiceFabricPlatformUpdateDomainWalk ...
func (c VirtualMachineScaleSetsClient) ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx context.Context, id VirtualMachineScaleSetId, options ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions) (result ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationResponse, err error) {
	req, err := c.preparerForForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ForceRecoveryServiceFabricPlatformUpdateDomainWalk", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ForceRecoveryServiceFabricPlatformUpdateDomainWalk", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForceRecoveryServiceFabricPlatformUpdateDomainWalk(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ForceRecoveryServiceFabricPlatformUpdateDomainWalk", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForceRecoveryServiceFabricPlatformUpdateDomainWalk prepares the ForceRecoveryServiceFabricPlatformUpdateDomainWalk request.
func (c VirtualMachineScaleSetsClient) preparerForForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx context.Context, id VirtualMachineScaleSetId, options ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/forceRecoveryServiceFabricPlatformUpdateDomainWalk", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForForceRecoveryServiceFabricPlatformUpdateDomainWalk handles the response to the ForceRecoveryServiceFabricPlatformUpdateDomainWalk request. The method always
// closes the http.Response Body.
func (c VirtualMachineScaleSetsClient) responderForForceRecoveryServiceFabricPlatformUpdateDomainWalk(resp *http.Response) (result ForceRecoveryServiceFabricPlatformUpdateDomainWalkOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
