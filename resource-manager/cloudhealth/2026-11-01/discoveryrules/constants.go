package discoveryrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveryRuleKind string

const (
	DiscoveryRuleKindApplicationInsightsTopology DiscoveryRuleKind = "ApplicationInsightsTopology"
	DiscoveryRuleKindResourceGraphQuery          DiscoveryRuleKind = "ResourceGraphQuery"
)

func PossibleValuesForDiscoveryRuleKind() []string {
	return []string{
		string(DiscoveryRuleKindApplicationInsightsTopology),
		string(DiscoveryRuleKindResourceGraphQuery),
	}
}

func (s *DiscoveryRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiscoveryRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiscoveryRuleKind(input string) (*DiscoveryRuleKind, error) {
	vals := map[string]DiscoveryRuleKind{
		"applicationinsightstopology": DiscoveryRuleKindApplicationInsightsTopology,
		"resourcegraphquery":          DiscoveryRuleKindResourceGraphQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiscoveryRuleKind(input)
	return &out, nil
}

type DiscoveryRuleRecommendedSignalsBehavior string

const (
	DiscoveryRuleRecommendedSignalsBehaviorDisabled DiscoveryRuleRecommendedSignalsBehavior = "Disabled"
	DiscoveryRuleRecommendedSignalsBehaviorEnabled  DiscoveryRuleRecommendedSignalsBehavior = "Enabled"
)

func PossibleValuesForDiscoveryRuleRecommendedSignalsBehavior() []string {
	return []string{
		string(DiscoveryRuleRecommendedSignalsBehaviorDisabled),
		string(DiscoveryRuleRecommendedSignalsBehaviorEnabled),
	}
}

func (s *DiscoveryRuleRecommendedSignalsBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiscoveryRuleRecommendedSignalsBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiscoveryRuleRecommendedSignalsBehavior(input string) (*DiscoveryRuleRecommendedSignalsBehavior, error) {
	vals := map[string]DiscoveryRuleRecommendedSignalsBehavior{
		"disabled": DiscoveryRuleRecommendedSignalsBehaviorDisabled,
		"enabled":  DiscoveryRuleRecommendedSignalsBehaviorEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiscoveryRuleRecommendedSignalsBehavior(input)
	return &out, nil
}

type DiscoveryRuleRelationshipDiscoveryBehavior string

const (
	DiscoveryRuleRelationshipDiscoveryBehaviorDisabled DiscoveryRuleRelationshipDiscoveryBehavior = "Disabled"
	DiscoveryRuleRelationshipDiscoveryBehaviorEnabled  DiscoveryRuleRelationshipDiscoveryBehavior = "Enabled"
)

func PossibleValuesForDiscoveryRuleRelationshipDiscoveryBehavior() []string {
	return []string{
		string(DiscoveryRuleRelationshipDiscoveryBehaviorDisabled),
		string(DiscoveryRuleRelationshipDiscoveryBehaviorEnabled),
	}
}

func (s *DiscoveryRuleRelationshipDiscoveryBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiscoveryRuleRelationshipDiscoveryBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiscoveryRuleRelationshipDiscoveryBehavior(input string) (*DiscoveryRuleRelationshipDiscoveryBehavior, error) {
	vals := map[string]DiscoveryRuleRelationshipDiscoveryBehavior{
		"disabled": DiscoveryRuleRelationshipDiscoveryBehaviorDisabled,
		"enabled":  DiscoveryRuleRelationshipDiscoveryBehaviorEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiscoveryRuleRelationshipDiscoveryBehavior(input)
	return &out, nil
}

type HealthModelProvisioningState string

const (
	HealthModelProvisioningStateCanceled  HealthModelProvisioningState = "Canceled"
	HealthModelProvisioningStateCreating  HealthModelProvisioningState = "Creating"
	HealthModelProvisioningStateDeleting  HealthModelProvisioningState = "Deleting"
	HealthModelProvisioningStateFailed    HealthModelProvisioningState = "Failed"
	HealthModelProvisioningStateSucceeded HealthModelProvisioningState = "Succeeded"
)

func PossibleValuesForHealthModelProvisioningState() []string {
	return []string{
		string(HealthModelProvisioningStateCanceled),
		string(HealthModelProvisioningStateCreating),
		string(HealthModelProvisioningStateDeleting),
		string(HealthModelProvisioningStateFailed),
		string(HealthModelProvisioningStateSucceeded),
	}
}

func (s *HealthModelProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthModelProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthModelProvisioningState(input string) (*HealthModelProvisioningState, error) {
	vals := map[string]HealthModelProvisioningState{
		"canceled":  HealthModelProvisioningStateCanceled,
		"creating":  HealthModelProvisioningStateCreating,
		"deleting":  HealthModelProvisioningStateDeleting,
		"failed":    HealthModelProvisioningStateFailed,
		"succeeded": HealthModelProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthModelProvisioningState(input)
	return &out, nil
}

type ResourceHealthAvailabilityStateSignalBehavior string

const (
	ResourceHealthAvailabilityStateSignalBehaviorDisabled ResourceHealthAvailabilityStateSignalBehavior = "Disabled"
	ResourceHealthAvailabilityStateSignalBehaviorEnabled  ResourceHealthAvailabilityStateSignalBehavior = "Enabled"
)

func PossibleValuesForResourceHealthAvailabilityStateSignalBehavior() []string {
	return []string{
		string(ResourceHealthAvailabilityStateSignalBehaviorDisabled),
		string(ResourceHealthAvailabilityStateSignalBehaviorEnabled),
	}
}

func (s *ResourceHealthAvailabilityStateSignalBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthAvailabilityStateSignalBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthAvailabilityStateSignalBehavior(input string) (*ResourceHealthAvailabilityStateSignalBehavior, error) {
	vals := map[string]ResourceHealthAvailabilityStateSignalBehavior{
		"disabled": ResourceHealthAvailabilityStateSignalBehaviorDisabled,
		"enabled":  ResourceHealthAvailabilityStateSignalBehaviorEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthAvailabilityStateSignalBehavior(input)
	return &out, nil
}
