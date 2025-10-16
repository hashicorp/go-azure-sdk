package registryendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type RegistryEndpointAuthenticationMethod string

const (
	RegistryEndpointAuthenticationMethodAnonymous                     RegistryEndpointAuthenticationMethod = "Anonymous"
	RegistryEndpointAuthenticationMethodArtifactPullSecret            RegistryEndpointAuthenticationMethod = "ArtifactPullSecret"
	RegistryEndpointAuthenticationMethodSystemAssignedManagedIdentity RegistryEndpointAuthenticationMethod = "SystemAssignedManagedIdentity"
	RegistryEndpointAuthenticationMethodUserAssignedManagedIdentity   RegistryEndpointAuthenticationMethod = "UserAssignedManagedIdentity"
)

func PossibleValuesForRegistryEndpointAuthenticationMethod() []string {
	return []string{
		string(RegistryEndpointAuthenticationMethodAnonymous),
		string(RegistryEndpointAuthenticationMethodArtifactPullSecret),
		string(RegistryEndpointAuthenticationMethodSystemAssignedManagedIdentity),
		string(RegistryEndpointAuthenticationMethodUserAssignedManagedIdentity),
	}
}

func (s *RegistryEndpointAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryEndpointAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryEndpointAuthenticationMethod(input string) (*RegistryEndpointAuthenticationMethod, error) {
	vals := map[string]RegistryEndpointAuthenticationMethod{
		"anonymous":                     RegistryEndpointAuthenticationMethodAnonymous,
		"artifactpullsecret":            RegistryEndpointAuthenticationMethodArtifactPullSecret,
		"systemassignedmanagedidentity": RegistryEndpointAuthenticationMethodSystemAssignedManagedIdentity,
		"userassignedmanagedidentity":   RegistryEndpointAuthenticationMethodUserAssignedManagedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryEndpointAuthenticationMethod(input)
	return &out, nil
}

type RegistryEndpointTrustedSigningKeyType string

const (
	RegistryEndpointTrustedSigningKeyTypeConfigMap RegistryEndpointTrustedSigningKeyType = "ConfigMap"
	RegistryEndpointTrustedSigningKeyTypeSecret    RegistryEndpointTrustedSigningKeyType = "Secret"
)

func PossibleValuesForRegistryEndpointTrustedSigningKeyType() []string {
	return []string{
		string(RegistryEndpointTrustedSigningKeyTypeConfigMap),
		string(RegistryEndpointTrustedSigningKeyTypeSecret),
	}
}

func (s *RegistryEndpointTrustedSigningKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryEndpointTrustedSigningKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryEndpointTrustedSigningKeyType(input string) (*RegistryEndpointTrustedSigningKeyType, error) {
	vals := map[string]RegistryEndpointTrustedSigningKeyType{
		"configmap": RegistryEndpointTrustedSigningKeyTypeConfigMap,
		"secret":    RegistryEndpointTrustedSigningKeyTypeSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryEndpointTrustedSigningKeyType(input)
	return &out, nil
}

type ResourceHealthState string

const (
	ResourceHealthStateAvailable   ResourceHealthState = "Available"
	ResourceHealthStateDegraded    ResourceHealthState = "Degraded"
	ResourceHealthStateUnavailable ResourceHealthState = "Unavailable"
	ResourceHealthStateUnknown     ResourceHealthState = "Unknown"
)

func PossibleValuesForResourceHealthState() []string {
	return []string{
		string(ResourceHealthStateAvailable),
		string(ResourceHealthStateDegraded),
		string(ResourceHealthStateUnavailable),
		string(ResourceHealthStateUnknown),
	}
}

func (s *ResourceHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthState(input string) (*ResourceHealthState, error) {
	vals := map[string]ResourceHealthState{
		"available":   ResourceHealthStateAvailable,
		"degraded":    ResourceHealthStateDegraded,
		"unavailable": ResourceHealthStateUnavailable,
		"unknown":     ResourceHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthState(input)
	return &out, nil
}
