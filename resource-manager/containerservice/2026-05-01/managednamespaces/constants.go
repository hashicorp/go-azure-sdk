package managednamespaces

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdoptionPolicy string

const (
	AdoptionPolicyAlways      AdoptionPolicy = "Always"
	AdoptionPolicyIfIdentical AdoptionPolicy = "IfIdentical"
	AdoptionPolicyNever       AdoptionPolicy = "Never"
)

func PossibleValuesForAdoptionPolicy() []string {
	return []string{
		string(AdoptionPolicyAlways),
		string(AdoptionPolicyIfIdentical),
		string(AdoptionPolicyNever),
	}
}

func (s *AdoptionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdoptionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdoptionPolicy(input string) (*AdoptionPolicy, error) {
	vals := map[string]AdoptionPolicy{
		"always":      AdoptionPolicyAlways,
		"ifidentical": AdoptionPolicyIfIdentical,
		"never":       AdoptionPolicyNever,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdoptionPolicy(input)
	return &out, nil
}

type DeletePolicy string

const (
	DeletePolicyDelete DeletePolicy = "Delete"
	DeletePolicyKeep   DeletePolicy = "Keep"
)

func PossibleValuesForDeletePolicy() []string {
	return []string{
		string(DeletePolicyDelete),
		string(DeletePolicyKeep),
	}
}

func (s *DeletePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeletePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeletePolicy(input string) (*DeletePolicy, error) {
	vals := map[string]DeletePolicy{
		"delete": DeletePolicyDelete,
		"keep":   DeletePolicyKeep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeletePolicy(input)
	return &out, nil
}

type NamespaceProvisioningState string

const (
	NamespaceProvisioningStateCanceled  NamespaceProvisioningState = "Canceled"
	NamespaceProvisioningStateCreating  NamespaceProvisioningState = "Creating"
	NamespaceProvisioningStateDeleting  NamespaceProvisioningState = "Deleting"
	NamespaceProvisioningStateFailed    NamespaceProvisioningState = "Failed"
	NamespaceProvisioningStateSucceeded NamespaceProvisioningState = "Succeeded"
	NamespaceProvisioningStateUpdating  NamespaceProvisioningState = "Updating"
)

func PossibleValuesForNamespaceProvisioningState() []string {
	return []string{
		string(NamespaceProvisioningStateCanceled),
		string(NamespaceProvisioningStateCreating),
		string(NamespaceProvisioningStateDeleting),
		string(NamespaceProvisioningStateFailed),
		string(NamespaceProvisioningStateSucceeded),
		string(NamespaceProvisioningStateUpdating),
	}
}

func (s *NamespaceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNamespaceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNamespaceProvisioningState(input string) (*NamespaceProvisioningState, error) {
	vals := map[string]NamespaceProvisioningState{
		"canceled":  NamespaceProvisioningStateCanceled,
		"creating":  NamespaceProvisioningStateCreating,
		"deleting":  NamespaceProvisioningStateDeleting,
		"failed":    NamespaceProvisioningStateFailed,
		"succeeded": NamespaceProvisioningStateSucceeded,
		"updating":  NamespaceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NamespaceProvisioningState(input)
	return &out, nil
}

type PolicyRule string

const (
	PolicyRuleAllowAll           PolicyRule = "AllowAll"
	PolicyRuleAllowSameNamespace PolicyRule = "AllowSameNamespace"
	PolicyRuleDenyAll            PolicyRule = "DenyAll"
)

func PossibleValuesForPolicyRule() []string {
	return []string{
		string(PolicyRuleAllowAll),
		string(PolicyRuleAllowSameNamespace),
		string(PolicyRuleDenyAll),
	}
}

func (s *PolicyRule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyRule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyRule(input string) (*PolicyRule, error) {
	vals := map[string]PolicyRule{
		"allowall":           PolicyRuleAllowAll,
		"allowsamenamespace": PolicyRuleAllowSameNamespace,
		"denyall":            PolicyRuleDenyAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyRule(input)
	return &out, nil
}
