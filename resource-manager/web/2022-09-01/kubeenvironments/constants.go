package kubeenvironments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FrontEndServiceType string

const (
	FrontEndServiceTypeLoadBalancer FrontEndServiceType = "LoadBalancer"
	FrontEndServiceTypeNodePort     FrontEndServiceType = "NodePort"
)

func PossibleValuesForFrontEndServiceType() []string {
	return []string{
		string(FrontEndServiceTypeLoadBalancer),
		string(FrontEndServiceTypeNodePort),
	}
}

func parseFrontEndServiceType(input string) (*FrontEndServiceType, error) {
	vals := map[string]FrontEndServiceType{
		"loadbalancer": FrontEndServiceTypeLoadBalancer,
		"nodeport":     FrontEndServiceTypeNodePort,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrontEndServiceType(input)
	return &out, nil
}

type KubeEnvironmentProvisioningState string

const (
	KubeEnvironmentProvisioningStateCanceled                      KubeEnvironmentProvisioningState = "Canceled"
	KubeEnvironmentProvisioningStateFailed                        KubeEnvironmentProvisioningState = "Failed"
	KubeEnvironmentProvisioningStateInfrastructureSetupComplete   KubeEnvironmentProvisioningState = "InfrastructureSetupComplete"
	KubeEnvironmentProvisioningStateInfrastructureSetupInProgress KubeEnvironmentProvisioningState = "InfrastructureSetupInProgress"
	KubeEnvironmentProvisioningStateInitializationInProgress      KubeEnvironmentProvisioningState = "InitializationInProgress"
	KubeEnvironmentProvisioningStateScheduledForDelete            KubeEnvironmentProvisioningState = "ScheduledForDelete"
	KubeEnvironmentProvisioningStateSucceeded                     KubeEnvironmentProvisioningState = "Succeeded"
	KubeEnvironmentProvisioningStateUpgradeFailed                 KubeEnvironmentProvisioningState = "UpgradeFailed"
	KubeEnvironmentProvisioningStateUpgradeRequested              KubeEnvironmentProvisioningState = "UpgradeRequested"
	KubeEnvironmentProvisioningStateWaiting                       KubeEnvironmentProvisioningState = "Waiting"
)

func PossibleValuesForKubeEnvironmentProvisioningState() []string {
	return []string{
		string(KubeEnvironmentProvisioningStateCanceled),
		string(KubeEnvironmentProvisioningStateFailed),
		string(KubeEnvironmentProvisioningStateInfrastructureSetupComplete),
		string(KubeEnvironmentProvisioningStateInfrastructureSetupInProgress),
		string(KubeEnvironmentProvisioningStateInitializationInProgress),
		string(KubeEnvironmentProvisioningStateScheduledForDelete),
		string(KubeEnvironmentProvisioningStateSucceeded),
		string(KubeEnvironmentProvisioningStateUpgradeFailed),
		string(KubeEnvironmentProvisioningStateUpgradeRequested),
		string(KubeEnvironmentProvisioningStateWaiting),
	}
}

func parseKubeEnvironmentProvisioningState(input string) (*KubeEnvironmentProvisioningState, error) {
	vals := map[string]KubeEnvironmentProvisioningState{
		"canceled":                      KubeEnvironmentProvisioningStateCanceled,
		"failed":                        KubeEnvironmentProvisioningStateFailed,
		"infrastructuresetupcomplete":   KubeEnvironmentProvisioningStateInfrastructureSetupComplete,
		"infrastructuresetupinprogress": KubeEnvironmentProvisioningStateInfrastructureSetupInProgress,
		"initializationinprogress":      KubeEnvironmentProvisioningStateInitializationInProgress,
		"scheduledfordelete":            KubeEnvironmentProvisioningStateScheduledForDelete,
		"succeeded":                     KubeEnvironmentProvisioningStateSucceeded,
		"upgradefailed":                 KubeEnvironmentProvisioningStateUpgradeFailed,
		"upgraderequested":              KubeEnvironmentProvisioningStateUpgradeRequested,
		"waiting":                       KubeEnvironmentProvisioningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubeEnvironmentProvisioningState(input)
	return &out, nil
}

type StorageType string

const (
	StorageTypeLocalNode         StorageType = "LocalNode"
	StorageTypeNetworkFileSystem StorageType = "NetworkFileSystem"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypeLocalNode),
		string(StorageTypeNetworkFileSystem),
	}
}

func parseStorageType(input string) (*StorageType, error) {
	vals := map[string]StorageType{
		"localnode":         StorageTypeLocalNode,
		"networkfilesystem": StorageTypeNetworkFileSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageType(input)
	return &out, nil
}
