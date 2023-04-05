package customresourceprovider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionRouting string

const (
	ActionRoutingProxy ActionRouting = "Proxy"
)

func PossibleValuesForActionRouting() []string {
	return []string{
		string(ActionRoutingProxy),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
	}
}

type ResourceTypeRouting string

const (
	ResourceTypeRoutingProxy      ResourceTypeRouting = "Proxy"
	ResourceTypeRoutingProxyCache ResourceTypeRouting = "Proxy,Cache"
)

func PossibleValuesForResourceTypeRouting() []string {
	return []string{
		string(ResourceTypeRoutingProxy),
		string(ResourceTypeRoutingProxyCache),
	}
}

type ValidationType string

const (
	ValidationTypeSwagger ValidationType = "Swagger"
)

func PossibleValuesForValidationType() []string {
	return []string{
		string(ValidationTypeSwagger),
	}
}
