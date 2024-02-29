package datacollectionendpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KnownDataCollectionEndpointProvisioningState string

const (
	KnownDataCollectionEndpointProvisioningStateCreating  KnownDataCollectionEndpointProvisioningState = "Creating"
	KnownDataCollectionEndpointProvisioningStateDeleting  KnownDataCollectionEndpointProvisioningState = "Deleting"
	KnownDataCollectionEndpointProvisioningStateFailed    KnownDataCollectionEndpointProvisioningState = "Failed"
	KnownDataCollectionEndpointProvisioningStateSucceeded KnownDataCollectionEndpointProvisioningState = "Succeeded"
	KnownDataCollectionEndpointProvisioningStateUpdating  KnownDataCollectionEndpointProvisioningState = "Updating"
)

func PossibleValuesForKnownDataCollectionEndpointProvisioningState() []string {
	return []string{
		string(KnownDataCollectionEndpointProvisioningStateCreating),
		string(KnownDataCollectionEndpointProvisioningStateDeleting),
		string(KnownDataCollectionEndpointProvisioningStateFailed),
		string(KnownDataCollectionEndpointProvisioningStateSucceeded),
		string(KnownDataCollectionEndpointProvisioningStateUpdating),
	}
}

func (s *KnownDataCollectionEndpointProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKnownDataCollectionEndpointProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKnownDataCollectionEndpointProvisioningState(input string) (*KnownDataCollectionEndpointProvisioningState, error) {
	vals := map[string]KnownDataCollectionEndpointProvisioningState{
		"creating":  KnownDataCollectionEndpointProvisioningStateCreating,
		"deleting":  KnownDataCollectionEndpointProvisioningStateDeleting,
		"failed":    KnownDataCollectionEndpointProvisioningStateFailed,
		"succeeded": KnownDataCollectionEndpointProvisioningStateSucceeded,
		"updating":  KnownDataCollectionEndpointProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KnownDataCollectionEndpointProvisioningState(input)
	return &out, nil
}

type KnownDataCollectionEndpointResourceKind string

const (
	KnownDataCollectionEndpointResourceKindLinux   KnownDataCollectionEndpointResourceKind = "Linux"
	KnownDataCollectionEndpointResourceKindWindows KnownDataCollectionEndpointResourceKind = "Windows"
)

func PossibleValuesForKnownDataCollectionEndpointResourceKind() []string {
	return []string{
		string(KnownDataCollectionEndpointResourceKindLinux),
		string(KnownDataCollectionEndpointResourceKindWindows),
	}
}

func (s *KnownDataCollectionEndpointResourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKnownDataCollectionEndpointResourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKnownDataCollectionEndpointResourceKind(input string) (*KnownDataCollectionEndpointResourceKind, error) {
	vals := map[string]KnownDataCollectionEndpointResourceKind{
		"linux":   KnownDataCollectionEndpointResourceKindLinux,
		"windows": KnownDataCollectionEndpointResourceKindWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KnownDataCollectionEndpointResourceKind(input)
	return &out, nil
}

type KnownPublicNetworkAccessOptions string

const (
	KnownPublicNetworkAccessOptionsDisabled KnownPublicNetworkAccessOptions = "Disabled"
	KnownPublicNetworkAccessOptionsEnabled  KnownPublicNetworkAccessOptions = "Enabled"
)

func PossibleValuesForKnownPublicNetworkAccessOptions() []string {
	return []string{
		string(KnownPublicNetworkAccessOptionsDisabled),
		string(KnownPublicNetworkAccessOptionsEnabled),
	}
}

func (s *KnownPublicNetworkAccessOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKnownPublicNetworkAccessOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKnownPublicNetworkAccessOptions(input string) (*KnownPublicNetworkAccessOptions, error) {
	vals := map[string]KnownPublicNetworkAccessOptions{
		"disabled": KnownPublicNetworkAccessOptionsDisabled,
		"enabled":  KnownPublicNetworkAccessOptionsEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KnownPublicNetworkAccessOptions(input)
	return &out, nil
}
