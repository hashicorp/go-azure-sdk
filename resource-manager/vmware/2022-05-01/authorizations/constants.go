package authorizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteAuthorizationProvisioningState string

const (
	ExpressRouteAuthorizationProvisioningStateCanceled  ExpressRouteAuthorizationProvisioningState = "Canceled"
	ExpressRouteAuthorizationProvisioningStateFailed    ExpressRouteAuthorizationProvisioningState = "Failed"
	ExpressRouteAuthorizationProvisioningStateSucceeded ExpressRouteAuthorizationProvisioningState = "Succeeded"
	ExpressRouteAuthorizationProvisioningStateUpdating  ExpressRouteAuthorizationProvisioningState = "Updating"
)

func PossibleValuesForExpressRouteAuthorizationProvisioningState() []string {
	return []string{
		string(ExpressRouteAuthorizationProvisioningStateCanceled),
		string(ExpressRouteAuthorizationProvisioningStateFailed),
		string(ExpressRouteAuthorizationProvisioningStateSucceeded),
		string(ExpressRouteAuthorizationProvisioningStateUpdating),
	}
}
