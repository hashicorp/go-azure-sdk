package serverlessendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentSafetyStatus string

const (
	ContentSafetyStatusDisabled ContentSafetyStatus = "Disabled"
	ContentSafetyStatusEnabled  ContentSafetyStatus = "Enabled"
)

func PossibleValuesForContentSafetyStatus() []string {
	return []string{
		string(ContentSafetyStatusDisabled),
		string(ContentSafetyStatusEnabled),
	}
}

func (s *ContentSafetyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentSafetyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentSafetyStatus(input string) (*ContentSafetyStatus, error) {
	vals := map[string]ContentSafetyStatus{
		"disabled": ContentSafetyStatusDisabled,
		"enabled":  ContentSafetyStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentSafetyStatus(input)
	return &out, nil
}

type EndpointProvisioningState string

const (
	EndpointProvisioningStateCanceled  EndpointProvisioningState = "Canceled"
	EndpointProvisioningStateCreating  EndpointProvisioningState = "Creating"
	EndpointProvisioningStateDeleting  EndpointProvisioningState = "Deleting"
	EndpointProvisioningStateFailed    EndpointProvisioningState = "Failed"
	EndpointProvisioningStateSucceeded EndpointProvisioningState = "Succeeded"
	EndpointProvisioningStateUpdating  EndpointProvisioningState = "Updating"
)

func PossibleValuesForEndpointProvisioningState() []string {
	return []string{
		string(EndpointProvisioningStateCanceled),
		string(EndpointProvisioningStateCreating),
		string(EndpointProvisioningStateDeleting),
		string(EndpointProvisioningStateFailed),
		string(EndpointProvisioningStateSucceeded),
		string(EndpointProvisioningStateUpdating),
	}
}

func (s *EndpointProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointProvisioningState(input string) (*EndpointProvisioningState, error) {
	vals := map[string]EndpointProvisioningState{
		"canceled":  EndpointProvisioningStateCanceled,
		"creating":  EndpointProvisioningStateCreating,
		"deleting":  EndpointProvisioningStateDeleting,
		"failed":    EndpointProvisioningStateFailed,
		"succeeded": EndpointProvisioningStateSucceeded,
		"updating":  EndpointProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointProvisioningState(input)
	return &out, nil
}

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSecondary KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

func (s *KeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeyType(input string) (*KeyType, error) {
	vals := map[string]KeyType{
		"primary":   KeyTypePrimary,
		"secondary": KeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyType(input)
	return &out, nil
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeNone),
		string(ManagedServiceIdentityTypeSystemAssigned),
		string(ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		string(ManagedServiceIdentityTypeUserAssigned),
	}
}

func (s *ManagedServiceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedServiceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedServiceIdentityType(input string) (*ManagedServiceIdentityType, error) {
	vals := map[string]ManagedServiceIdentityType{
		"none":                        ManagedServiceIdentityTypeNone,
		"systemassigned":              ManagedServiceIdentityTypeSystemAssigned,
		"systemassigned,userassigned": ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		"userassigned":                ManagedServiceIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedServiceIdentityType(input)
	return &out, nil
}

type ServerlessEndpointState string

const (
	ServerlessEndpointStateCreating       ServerlessEndpointState = "Creating"
	ServerlessEndpointStateCreationFailed ServerlessEndpointState = "CreationFailed"
	ServerlessEndpointStateDeleting       ServerlessEndpointState = "Deleting"
	ServerlessEndpointStateDeletionFailed ServerlessEndpointState = "DeletionFailed"
	ServerlessEndpointStateOnline         ServerlessEndpointState = "Online"
	ServerlessEndpointStateReinstating    ServerlessEndpointState = "Reinstating"
	ServerlessEndpointStateSuspended      ServerlessEndpointState = "Suspended"
	ServerlessEndpointStateSuspending     ServerlessEndpointState = "Suspending"
	ServerlessEndpointStateUnknown        ServerlessEndpointState = "Unknown"
)

func PossibleValuesForServerlessEndpointState() []string {
	return []string{
		string(ServerlessEndpointStateCreating),
		string(ServerlessEndpointStateCreationFailed),
		string(ServerlessEndpointStateDeleting),
		string(ServerlessEndpointStateDeletionFailed),
		string(ServerlessEndpointStateOnline),
		string(ServerlessEndpointStateReinstating),
		string(ServerlessEndpointStateSuspended),
		string(ServerlessEndpointStateSuspending),
		string(ServerlessEndpointStateUnknown),
	}
}

func (s *ServerlessEndpointState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerlessEndpointState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerlessEndpointState(input string) (*ServerlessEndpointState, error) {
	vals := map[string]ServerlessEndpointState{
		"creating":       ServerlessEndpointStateCreating,
		"creationfailed": ServerlessEndpointStateCreationFailed,
		"deleting":       ServerlessEndpointStateDeleting,
		"deletionfailed": ServerlessEndpointStateDeletionFailed,
		"online":         ServerlessEndpointStateOnline,
		"reinstating":    ServerlessEndpointStateReinstating,
		"suspended":      ServerlessEndpointStateSuspended,
		"suspending":     ServerlessEndpointStateSuspending,
		"unknown":        ServerlessEndpointStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerlessEndpointState(input)
	return &out, nil
}

type ServerlessInferenceEndpointAuthMode string

const (
	ServerlessInferenceEndpointAuthModeKey ServerlessInferenceEndpointAuthMode = "Key"
)

func PossibleValuesForServerlessInferenceEndpointAuthMode() []string {
	return []string{
		string(ServerlessInferenceEndpointAuthModeKey),
	}
}

func (s *ServerlessInferenceEndpointAuthMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerlessInferenceEndpointAuthMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerlessInferenceEndpointAuthMode(input string) (*ServerlessInferenceEndpointAuthMode, error) {
	vals := map[string]ServerlessInferenceEndpointAuthMode{
		"key": ServerlessInferenceEndpointAuthModeKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerlessInferenceEndpointAuthMode(input)
	return &out, nil
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

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
