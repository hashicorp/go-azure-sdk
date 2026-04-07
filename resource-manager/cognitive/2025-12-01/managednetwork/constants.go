package managednetwork

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallSku string

const (
	FirewallSkuBasic    FirewallSku = "Basic"
	FirewallSkuStandard FirewallSku = "Standard"
)

func PossibleValuesForFirewallSku() []string {
	return []string{
		string(FirewallSkuBasic),
		string(FirewallSkuStandard),
	}
}

func (s *FirewallSku) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallSku(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallSku(input string) (*FirewallSku, error) {
	vals := map[string]FirewallSku{
		"basic":    FirewallSkuBasic,
		"standard": FirewallSkuStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallSku(input)
	return &out, nil
}

type IsolationMode string

const (
	IsolationModeAllowInternetOutbound     IsolationMode = "AllowInternetOutbound"
	IsolationModeAllowOnlyApprovedOutbound IsolationMode = "AllowOnlyApprovedOutbound"
	IsolationModeDisabled                  IsolationMode = "Disabled"
)

func PossibleValuesForIsolationMode() []string {
	return []string{
		string(IsolationModeAllowInternetOutbound),
		string(IsolationModeAllowOnlyApprovedOutbound),
		string(IsolationModeDisabled),
	}
}

func (s *IsolationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIsolationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIsolationMode(input string) (*IsolationMode, error) {
	vals := map[string]IsolationMode{
		"allowinternetoutbound":     IsolationModeAllowInternetOutbound,
		"allowonlyapprovedoutbound": IsolationModeAllowOnlyApprovedOutbound,
		"disabled":                  IsolationModeDisabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IsolationMode(input)
	return &out, nil
}

type ManagedNetworkKind string

const (
	ManagedNetworkKindVOne ManagedNetworkKind = "V1"
	ManagedNetworkKindVTwo ManagedNetworkKind = "V2"
)

func PossibleValuesForManagedNetworkKind() []string {
	return []string{
		string(ManagedNetworkKindVOne),
		string(ManagedNetworkKindVTwo),
	}
}

func (s *ManagedNetworkKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedNetworkKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedNetworkKind(input string) (*ManagedNetworkKind, error) {
	vals := map[string]ManagedNetworkKind{
		"v1": ManagedNetworkKindVOne,
		"v2": ManagedNetworkKindVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedNetworkKind(input)
	return &out, nil
}

type ManagedNetworkProvisioningState string

const (
	ManagedNetworkProvisioningStateDeferred  ManagedNetworkProvisioningState = "Deferred"
	ManagedNetworkProvisioningStateDeleted   ManagedNetworkProvisioningState = "Deleted"
	ManagedNetworkProvisioningStateDeleting  ManagedNetworkProvisioningState = "Deleting"
	ManagedNetworkProvisioningStateFailed    ManagedNetworkProvisioningState = "Failed"
	ManagedNetworkProvisioningStateSucceeded ManagedNetworkProvisioningState = "Succeeded"
	ManagedNetworkProvisioningStateUpdating  ManagedNetworkProvisioningState = "Updating"
)

func PossibleValuesForManagedNetworkProvisioningState() []string {
	return []string{
		string(ManagedNetworkProvisioningStateDeferred),
		string(ManagedNetworkProvisioningStateDeleted),
		string(ManagedNetworkProvisioningStateDeleting),
		string(ManagedNetworkProvisioningStateFailed),
		string(ManagedNetworkProvisioningStateSucceeded),
		string(ManagedNetworkProvisioningStateUpdating),
	}
}

func (s *ManagedNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedNetworkProvisioningState(input string) (*ManagedNetworkProvisioningState, error) {
	vals := map[string]ManagedNetworkProvisioningState{
		"deferred":  ManagedNetworkProvisioningStateDeferred,
		"deleted":   ManagedNetworkProvisioningStateDeleted,
		"deleting":  ManagedNetworkProvisioningStateDeleting,
		"failed":    ManagedNetworkProvisioningStateFailed,
		"succeeded": ManagedNetworkProvisioningStateSucceeded,
		"updating":  ManagedNetworkProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedNetworkProvisioningState(input)
	return &out, nil
}

type ManagedNetworkStatus string

const (
	ManagedNetworkStatusActive   ManagedNetworkStatus = "Active"
	ManagedNetworkStatusInactive ManagedNetworkStatus = "Inactive"
)

func PossibleValuesForManagedNetworkStatus() []string {
	return []string{
		string(ManagedNetworkStatusActive),
		string(ManagedNetworkStatusInactive),
	}
}

func (s *ManagedNetworkStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedNetworkStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedNetworkStatus(input string) (*ManagedNetworkStatus, error) {
	vals := map[string]ManagedNetworkStatus{
		"active":   ManagedNetworkStatusActive,
		"inactive": ManagedNetworkStatusInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedNetworkStatus(input)
	return &out, nil
}

type RuleAction string

const (
	RuleActionAllow RuleAction = "Allow"
	RuleActionDeny  RuleAction = "Deny"
)

func PossibleValuesForRuleAction() []string {
	return []string{
		string(RuleActionAllow),
		string(RuleActionDeny),
	}
}

func (s *RuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleAction(input string) (*RuleAction, error) {
	vals := map[string]RuleAction{
		"allow": RuleActionAllow,
		"deny":  RuleActionDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleAction(input)
	return &out, nil
}

type RuleCategory string

const (
	RuleCategoryDependency  RuleCategory = "Dependency"
	RuleCategoryRecommended RuleCategory = "Recommended"
	RuleCategoryRequired    RuleCategory = "Required"
	RuleCategoryUserDefined RuleCategory = "UserDefined"
)

func PossibleValuesForRuleCategory() []string {
	return []string{
		string(RuleCategoryDependency),
		string(RuleCategoryRecommended),
		string(RuleCategoryRequired),
		string(RuleCategoryUserDefined),
	}
}

func (s *RuleCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleCategory(input string) (*RuleCategory, error) {
	vals := map[string]RuleCategory{
		"dependency":  RuleCategoryDependency,
		"recommended": RuleCategoryRecommended,
		"required":    RuleCategoryRequired,
		"userdefined": RuleCategoryUserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleCategory(input)
	return &out, nil
}

type RuleStatus string

const (
	RuleStatusActive       RuleStatus = "Active"
	RuleStatusDeleting     RuleStatus = "Deleting"
	RuleStatusFailed       RuleStatus = "Failed"
	RuleStatusInactive     RuleStatus = "Inactive"
	RuleStatusProvisioning RuleStatus = "Provisioning"
)

func PossibleValuesForRuleStatus() []string {
	return []string{
		string(RuleStatusActive),
		string(RuleStatusDeleting),
		string(RuleStatusFailed),
		string(RuleStatusInactive),
		string(RuleStatusProvisioning),
	}
}

func (s *RuleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleStatus(input string) (*RuleStatus, error) {
	vals := map[string]RuleStatus{
		"active":       RuleStatusActive,
		"deleting":     RuleStatusDeleting,
		"failed":       RuleStatusFailed,
		"inactive":     RuleStatusInactive,
		"provisioning": RuleStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleStatus(input)
	return &out, nil
}

type RuleType string

const (
	RuleTypeFQDN            RuleType = "FQDN"
	RuleTypePrivateEndpoint RuleType = "PrivateEndpoint"
	RuleTypeServiceTag      RuleType = "ServiceTag"
)

func PossibleValuesForRuleType() []string {
	return []string{
		string(RuleTypeFQDN),
		string(RuleTypePrivateEndpoint),
		string(RuleTypeServiceTag),
	}
}

func (s *RuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleType(input string) (*RuleType, error) {
	vals := map[string]RuleType{
		"fqdn":            RuleTypeFQDN,
		"privateendpoint": RuleTypePrivateEndpoint,
		"servicetag":      RuleTypeServiceTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleType(input)
	return &out, nil
}
