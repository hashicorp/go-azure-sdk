package onlinedeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerType string

const (
	ContainerTypeInferenceServer    ContainerType = "InferenceServer"
	ContainerTypeStorageInitializer ContainerType = "StorageInitializer"
)

func PossibleValuesForContainerType() []string {
	return []string{
		string(ContainerTypeInferenceServer),
		string(ContainerTypeStorageInitializer),
	}
}

type DeploymentProvisioningState string

const (
	DeploymentProvisioningStateCanceled  DeploymentProvisioningState = "Canceled"
	DeploymentProvisioningStateCreating  DeploymentProvisioningState = "Creating"
	DeploymentProvisioningStateDeleting  DeploymentProvisioningState = "Deleting"
	DeploymentProvisioningStateFailed    DeploymentProvisioningState = "Failed"
	DeploymentProvisioningStateScaling   DeploymentProvisioningState = "Scaling"
	DeploymentProvisioningStateSucceeded DeploymentProvisioningState = "Succeeded"
	DeploymentProvisioningStateUpdating  DeploymentProvisioningState = "Updating"
)

func PossibleValuesForDeploymentProvisioningState() []string {
	return []string{
		string(DeploymentProvisioningStateCanceled),
		string(DeploymentProvisioningStateCreating),
		string(DeploymentProvisioningStateDeleting),
		string(DeploymentProvisioningStateFailed),
		string(DeploymentProvisioningStateScaling),
		string(DeploymentProvisioningStateSucceeded),
		string(DeploymentProvisioningStateUpdating),
	}
}

type EgressPublicNetworkAccessType string

const (
	EgressPublicNetworkAccessTypeDisabled EgressPublicNetworkAccessType = "Disabled"
	EgressPublicNetworkAccessTypeEnabled  EgressPublicNetworkAccessType = "Enabled"
)

func PossibleValuesForEgressPublicNetworkAccessType() []string {
	return []string{
		string(EgressPublicNetworkAccessTypeDisabled),
		string(EgressPublicNetworkAccessTypeEnabled),
	}
}

type EndpointComputeType string

const (
	EndpointComputeTypeAzureMLCompute EndpointComputeType = "AzureMLCompute"
	EndpointComputeTypeKubernetes     EndpointComputeType = "Kubernetes"
	EndpointComputeTypeManaged        EndpointComputeType = "Managed"
)

func PossibleValuesForEndpointComputeType() []string {
	return []string{
		string(EndpointComputeTypeAzureMLCompute),
		string(EndpointComputeTypeKubernetes),
		string(EndpointComputeTypeManaged),
	}
}

type ScaleType string

const (
	ScaleTypeDefault           ScaleType = "Default"
	ScaleTypeTargetUtilization ScaleType = "TargetUtilization"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeDefault),
		string(ScaleTypeTargetUtilization),
	}
}

type SkuScaleType string

const (
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	SkuScaleTypeManual    SkuScaleType = "Manual"
	SkuScaleTypeNone      SkuScaleType = "None"
)

func PossibleValuesForSkuScaleType() []string {
	return []string{
		string(SkuScaleTypeAutomatic),
		string(SkuScaleTypeManual),
		string(SkuScaleTypeNone),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
