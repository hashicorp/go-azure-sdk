package kubeenvironments

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *FrontEndServiceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrontEndServiceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *KubeEnvironmentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubeEnvironmentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *StorageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
