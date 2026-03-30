package kubernetesclusterfeatures

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterFeatureAvailabilityLifecycle string

const (
	KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable KubernetesClusterFeatureAvailabilityLifecycle = "GenerallyAvailable"
	KubernetesClusterFeatureAvailabilityLifecyclePreview            KubernetesClusterFeatureAvailabilityLifecycle = "Preview"
)

func PossibleValuesForKubernetesClusterFeatureAvailabilityLifecycle() []string {
	return []string{
		string(KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
		string(KubernetesClusterFeatureAvailabilityLifecyclePreview),
	}
}

func (s *KubernetesClusterFeatureAvailabilityLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureAvailabilityLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureAvailabilityLifecycle(input string) (*KubernetesClusterFeatureAvailabilityLifecycle, error) {
	vals := map[string]KubernetesClusterFeatureAvailabilityLifecycle{
		"generallyavailable": KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable,
		"preview":            KubernetesClusterFeatureAvailabilityLifecyclePreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureAvailabilityLifecycle(input)
	return &out, nil
}

type KubernetesClusterFeatureDetailedStatus string

const (
	KubernetesClusterFeatureDetailedStatusError        KubernetesClusterFeatureDetailedStatus = "Error"
	KubernetesClusterFeatureDetailedStatusInstalled    KubernetesClusterFeatureDetailedStatus = "Installed"
	KubernetesClusterFeatureDetailedStatusProvisioning KubernetesClusterFeatureDetailedStatus = "Provisioning"
)

func PossibleValuesForKubernetesClusterFeatureDetailedStatus() []string {
	return []string{
		string(KubernetesClusterFeatureDetailedStatusError),
		string(KubernetesClusterFeatureDetailedStatusInstalled),
		string(KubernetesClusterFeatureDetailedStatusProvisioning),
	}
}

func (s *KubernetesClusterFeatureDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureDetailedStatus(input string) (*KubernetesClusterFeatureDetailedStatus, error) {
	vals := map[string]KubernetesClusterFeatureDetailedStatus{
		"error":        KubernetesClusterFeatureDetailedStatusError,
		"installed":    KubernetesClusterFeatureDetailedStatusInstalled,
		"provisioning": KubernetesClusterFeatureDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterFeatureProvisioningState string

const (
	KubernetesClusterFeatureProvisioningStateAccepted  KubernetesClusterFeatureProvisioningState = "Accepted"
	KubernetesClusterFeatureProvisioningStateCanceled  KubernetesClusterFeatureProvisioningState = "Canceled"
	KubernetesClusterFeatureProvisioningStateDeleting  KubernetesClusterFeatureProvisioningState = "Deleting"
	KubernetesClusterFeatureProvisioningStateFailed    KubernetesClusterFeatureProvisioningState = "Failed"
	KubernetesClusterFeatureProvisioningStateSucceeded KubernetesClusterFeatureProvisioningState = "Succeeded"
	KubernetesClusterFeatureProvisioningStateUpdating  KubernetesClusterFeatureProvisioningState = "Updating"
)

func PossibleValuesForKubernetesClusterFeatureProvisioningState() []string {
	return []string{
		string(KubernetesClusterFeatureProvisioningStateAccepted),
		string(KubernetesClusterFeatureProvisioningStateCanceled),
		string(KubernetesClusterFeatureProvisioningStateDeleting),
		string(KubernetesClusterFeatureProvisioningStateFailed),
		string(KubernetesClusterFeatureProvisioningStateSucceeded),
		string(KubernetesClusterFeatureProvisioningStateUpdating),
	}
}

func (s *KubernetesClusterFeatureProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureProvisioningState(input string) (*KubernetesClusterFeatureProvisioningState, error) {
	vals := map[string]KubernetesClusterFeatureProvisioningState{
		"accepted":  KubernetesClusterFeatureProvisioningStateAccepted,
		"canceled":  KubernetesClusterFeatureProvisioningStateCanceled,
		"deleting":  KubernetesClusterFeatureProvisioningStateDeleting,
		"failed":    KubernetesClusterFeatureProvisioningStateFailed,
		"succeeded": KubernetesClusterFeatureProvisioningStateSucceeded,
		"updating":  KubernetesClusterFeatureProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureProvisioningState(input)
	return &out, nil
}

type KubernetesClusterFeatureRequired string

const (
	KubernetesClusterFeatureRequiredFalse KubernetesClusterFeatureRequired = "False"
	KubernetesClusterFeatureRequiredTrue  KubernetesClusterFeatureRequired = "True"
)

func PossibleValuesForKubernetesClusterFeatureRequired() []string {
	return []string{
		string(KubernetesClusterFeatureRequiredFalse),
		string(KubernetesClusterFeatureRequiredTrue),
	}
}

func (s *KubernetesClusterFeatureRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureRequired(input string) (*KubernetesClusterFeatureRequired, error) {
	vals := map[string]KubernetesClusterFeatureRequired{
		"false": KubernetesClusterFeatureRequiredFalse,
		"true":  KubernetesClusterFeatureRequiredTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureRequired(input)
	return &out, nil
}
