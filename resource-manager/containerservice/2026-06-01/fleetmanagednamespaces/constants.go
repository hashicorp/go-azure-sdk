package fleetmanagednamespaces

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

type FleetManagedNamespaceProvisioningState string

const (
	FleetManagedNamespaceProvisioningStateCanceled  FleetManagedNamespaceProvisioningState = "Canceled"
	FleetManagedNamespaceProvisioningStateCreating  FleetManagedNamespaceProvisioningState = "Creating"
	FleetManagedNamespaceProvisioningStateDeleting  FleetManagedNamespaceProvisioningState = "Deleting"
	FleetManagedNamespaceProvisioningStateFailed    FleetManagedNamespaceProvisioningState = "Failed"
	FleetManagedNamespaceProvisioningStateSucceeded FleetManagedNamespaceProvisioningState = "Succeeded"
	FleetManagedNamespaceProvisioningStateUpdating  FleetManagedNamespaceProvisioningState = "Updating"
)

func PossibleValuesForFleetManagedNamespaceProvisioningState() []string {
	return []string{
		string(FleetManagedNamespaceProvisioningStateCanceled),
		string(FleetManagedNamespaceProvisioningStateCreating),
		string(FleetManagedNamespaceProvisioningStateDeleting),
		string(FleetManagedNamespaceProvisioningStateFailed),
		string(FleetManagedNamespaceProvisioningStateSucceeded),
		string(FleetManagedNamespaceProvisioningStateUpdating),
	}
}

func (s *FleetManagedNamespaceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFleetManagedNamespaceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFleetManagedNamespaceProvisioningState(input string) (*FleetManagedNamespaceProvisioningState, error) {
	vals := map[string]FleetManagedNamespaceProvisioningState{
		"canceled":  FleetManagedNamespaceProvisioningStateCanceled,
		"creating":  FleetManagedNamespaceProvisioningStateCreating,
		"deleting":  FleetManagedNamespaceProvisioningStateDeleting,
		"failed":    FleetManagedNamespaceProvisioningStateFailed,
		"succeeded": FleetManagedNamespaceProvisioningStateSucceeded,
		"updating":  FleetManagedNamespaceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FleetManagedNamespaceProvisioningState(input)
	return &out, nil
}

type LabelSelectorOperator string

const (
	LabelSelectorOperatorDoesNotExist LabelSelectorOperator = "DoesNotExist"
	LabelSelectorOperatorExists       LabelSelectorOperator = "Exists"
	LabelSelectorOperatorIn           LabelSelectorOperator = "In"
	LabelSelectorOperatorNotIn        LabelSelectorOperator = "NotIn"
)

func PossibleValuesForLabelSelectorOperator() []string {
	return []string{
		string(LabelSelectorOperatorDoesNotExist),
		string(LabelSelectorOperatorExists),
		string(LabelSelectorOperatorIn),
		string(LabelSelectorOperatorNotIn),
	}
}

func (s *LabelSelectorOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabelSelectorOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabelSelectorOperator(input string) (*LabelSelectorOperator, error) {
	vals := map[string]LabelSelectorOperator{
		"doesnotexist": LabelSelectorOperatorDoesNotExist,
		"exists":       LabelSelectorOperatorExists,
		"in":           LabelSelectorOperatorIn,
		"notin":        LabelSelectorOperatorNotIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelSelectorOperator(input)
	return &out, nil
}

type PlacementType string

const (
	PlacementTypePickAll   PlacementType = "PickAll"
	PlacementTypePickFixed PlacementType = "PickFixed"
)

func PossibleValuesForPlacementType() []string {
	return []string{
		string(PlacementTypePickAll),
		string(PlacementTypePickFixed),
	}
}

func (s *PlacementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlacementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlacementType(input string) (*PlacementType, error) {
	vals := map[string]PlacementType{
		"pickall":   PlacementTypePickAll,
		"pickfixed": PlacementTypePickFixed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlacementType(input)
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

type PropagationType string

const (
	PropagationTypePlacement PropagationType = "Placement"
)

func PossibleValuesForPropagationType() []string {
	return []string{
		string(PropagationTypePlacement),
	}
}

func (s *PropagationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePropagationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePropagationType(input string) (*PropagationType, error) {
	vals := map[string]PropagationType{
		"placement": PropagationTypePlacement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropagationType(input)
	return &out, nil
}

type PropertySelectorOperator string

const (
	PropertySelectorOperatorEq PropertySelectorOperator = "Eq"
	PropertySelectorOperatorGe PropertySelectorOperator = "Ge"
	PropertySelectorOperatorGt PropertySelectorOperator = "Gt"
	PropertySelectorOperatorLe PropertySelectorOperator = "Le"
	PropertySelectorOperatorLt PropertySelectorOperator = "Lt"
	PropertySelectorOperatorNe PropertySelectorOperator = "Ne"
)

func PossibleValuesForPropertySelectorOperator() []string {
	return []string{
		string(PropertySelectorOperatorEq),
		string(PropertySelectorOperatorGe),
		string(PropertySelectorOperatorGt),
		string(PropertySelectorOperatorLe),
		string(PropertySelectorOperatorLt),
		string(PropertySelectorOperatorNe),
	}
}

func (s *PropertySelectorOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePropertySelectorOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePropertySelectorOperator(input string) (*PropertySelectorOperator, error) {
	vals := map[string]PropertySelectorOperator{
		"eq": PropertySelectorOperatorEq,
		"ge": PropertySelectorOperatorGe,
		"gt": PropertySelectorOperatorGt,
		"le": PropertySelectorOperatorLe,
		"lt": PropertySelectorOperatorLt,
		"ne": PropertySelectorOperatorNe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertySelectorOperator(input)
	return &out, nil
}

type TaintEffect string

const (
	TaintEffectNoSchedule TaintEffect = "NoSchedule"
)

func PossibleValuesForTaintEffect() []string {
	return []string{
		string(TaintEffectNoSchedule),
	}
}

func (s *TaintEffect) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaintEffect(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaintEffect(input string) (*TaintEffect, error) {
	vals := map[string]TaintEffect{
		"noschedule": TaintEffectNoSchedule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaintEffect(input)
	return &out, nil
}

type TolerationOperator string

const (
	TolerationOperatorEqual  TolerationOperator = "Equal"
	TolerationOperatorExists TolerationOperator = "Exists"
)

func PossibleValuesForTolerationOperator() []string {
	return []string{
		string(TolerationOperatorEqual),
		string(TolerationOperatorExists),
	}
}

func (s *TolerationOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTolerationOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTolerationOperator(input string) (*TolerationOperator, error) {
	vals := map[string]TolerationOperator{
		"equal":  TolerationOperatorEqual,
		"exists": TolerationOperatorExists,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TolerationOperator(input)
	return &out, nil
}
