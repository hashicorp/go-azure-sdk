package agentapplication

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgenticApplicationProvisioningState string

const (
	AgenticApplicationProvisioningStateCanceled  AgenticApplicationProvisioningState = "Canceled"
	AgenticApplicationProvisioningStateCreating  AgenticApplicationProvisioningState = "Creating"
	AgenticApplicationProvisioningStateDeleting  AgenticApplicationProvisioningState = "Deleting"
	AgenticApplicationProvisioningStateFailed    AgenticApplicationProvisioningState = "Failed"
	AgenticApplicationProvisioningStateSucceeded AgenticApplicationProvisioningState = "Succeeded"
	AgenticApplicationProvisioningStateUpdating  AgenticApplicationProvisioningState = "Updating"
)

func PossibleValuesForAgenticApplicationProvisioningState() []string {
	return []string{
		string(AgenticApplicationProvisioningStateCanceled),
		string(AgenticApplicationProvisioningStateCreating),
		string(AgenticApplicationProvisioningStateDeleting),
		string(AgenticApplicationProvisioningStateFailed),
		string(AgenticApplicationProvisioningStateSucceeded),
		string(AgenticApplicationProvisioningStateUpdating),
	}
}

func (s *AgenticApplicationProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgenticApplicationProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgenticApplicationProvisioningState(input string) (*AgenticApplicationProvisioningState, error) {
	vals := map[string]AgenticApplicationProvisioningState{
		"canceled":  AgenticApplicationProvisioningStateCanceled,
		"creating":  AgenticApplicationProvisioningStateCreating,
		"deleting":  AgenticApplicationProvisioningStateDeleting,
		"failed":    AgenticApplicationProvisioningStateFailed,
		"succeeded": AgenticApplicationProvisioningStateSucceeded,
		"updating":  AgenticApplicationProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgenticApplicationProvisioningState(input)
	return &out, nil
}

type BuiltInAuthorizationScheme string

const (
	BuiltInAuthorizationSchemeChannels          BuiltInAuthorizationScheme = "Channels"
	BuiltInAuthorizationSchemeCustom            BuiltInAuthorizationScheme = "Custom"
	BuiltInAuthorizationSchemeDefault           BuiltInAuthorizationScheme = "Default"
	BuiltInAuthorizationSchemeOrganizationScope BuiltInAuthorizationScheme = "OrganizationScope"
)

func PossibleValuesForBuiltInAuthorizationScheme() []string {
	return []string{
		string(BuiltInAuthorizationSchemeChannels),
		string(BuiltInAuthorizationSchemeCustom),
		string(BuiltInAuthorizationSchemeDefault),
		string(BuiltInAuthorizationSchemeOrganizationScope),
	}
}

func (s *BuiltInAuthorizationScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBuiltInAuthorizationScheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBuiltInAuthorizationScheme(input string) (*BuiltInAuthorizationScheme, error) {
	vals := map[string]BuiltInAuthorizationScheme{
		"channels":          BuiltInAuthorizationSchemeChannels,
		"custom":            BuiltInAuthorizationSchemeCustom,
		"default":           BuiltInAuthorizationSchemeDefault,
		"organizationscope": BuiltInAuthorizationSchemeOrganizationScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuiltInAuthorizationScheme(input)
	return &out, nil
}

type IdentityKind string

const (
	IdentityKindAgentBlueprint IdentityKind = "AgentBlueprint"
	IdentityKindAgentInstance  IdentityKind = "AgentInstance"
	IdentityKindAgenticUser    IdentityKind = "AgenticUser"
	IdentityKindManaged        IdentityKind = "Managed"
	IdentityKindNone           IdentityKind = "None"
)

func PossibleValuesForIdentityKind() []string {
	return []string{
		string(IdentityKindAgentBlueprint),
		string(IdentityKindAgentInstance),
		string(IdentityKindAgenticUser),
		string(IdentityKindManaged),
		string(IdentityKindNone),
	}
}

func (s *IdentityKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityKind(input string) (*IdentityKind, error) {
	vals := map[string]IdentityKind{
		"agentblueprint": IdentityKindAgentBlueprint,
		"agentinstance":  IdentityKindAgentInstance,
		"agenticuser":    IdentityKindAgenticUser,
		"managed":        IdentityKindManaged,
		"none":           IdentityKindNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityKind(input)
	return &out, nil
}

type IdentityManagementType string

const (
	IdentityManagementTypeNone   IdentityManagementType = "None"
	IdentityManagementTypeSystem IdentityManagementType = "System"
	IdentityManagementTypeUser   IdentityManagementType = "User"
)

func PossibleValuesForIdentityManagementType() []string {
	return []string{
		string(IdentityManagementTypeNone),
		string(IdentityManagementTypeSystem),
		string(IdentityManagementTypeUser),
	}
}

func (s *IdentityManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityManagementType(input string) (*IdentityManagementType, error) {
	vals := map[string]IdentityManagementType{
		"none":   IdentityManagementTypeNone,
		"system": IdentityManagementTypeSystem,
		"user":   IdentityManagementTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityManagementType(input)
	return &out, nil
}

type IdentityProvisioningState string

const (
	IdentityProvisioningStateCanceled  IdentityProvisioningState = "Canceled"
	IdentityProvisioningStateCreating  IdentityProvisioningState = "Creating"
	IdentityProvisioningStateDeleting  IdentityProvisioningState = "Deleting"
	IdentityProvisioningStateFailed    IdentityProvisioningState = "Failed"
	IdentityProvisioningStateSucceeded IdentityProvisioningState = "Succeeded"
	IdentityProvisioningStateUpdating  IdentityProvisioningState = "Updating"
)

func PossibleValuesForIdentityProvisioningState() []string {
	return []string{
		string(IdentityProvisioningStateCanceled),
		string(IdentityProvisioningStateCreating),
		string(IdentityProvisioningStateDeleting),
		string(IdentityProvisioningStateFailed),
		string(IdentityProvisioningStateSucceeded),
		string(IdentityProvisioningStateUpdating),
	}
}

func (s *IdentityProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityProvisioningState(input string) (*IdentityProvisioningState, error) {
	vals := map[string]IdentityProvisioningState{
		"canceled":  IdentityProvisioningStateCanceled,
		"creating":  IdentityProvisioningStateCreating,
		"deleting":  IdentityProvisioningStateDeleting,
		"failed":    IdentityProvisioningStateFailed,
		"succeeded": IdentityProvisioningStateSucceeded,
		"updating":  IdentityProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityProvisioningState(input)
	return &out, nil
}

type TrafficRoutingProtocol string

const (
	TrafficRoutingProtocolFixedRatio TrafficRoutingProtocol = "FixedRatio"
)

func PossibleValuesForTrafficRoutingProtocol() []string {
	return []string{
		string(TrafficRoutingProtocolFixedRatio),
	}
}

func (s *TrafficRoutingProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrafficRoutingProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrafficRoutingProtocol(input string) (*TrafficRoutingProtocol, error) {
	vals := map[string]TrafficRoutingProtocol{
		"fixedratio": TrafficRoutingProtocolFixedRatio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrafficRoutingProtocol(input)
	return &out, nil
}
