package accounts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AADObjectType string

const (
	AADObjectTypeGroup            AADObjectType = "Group"
	AADObjectTypeServicePrincipal AADObjectType = "ServicePrincipal"
	AADObjectTypeUser             AADObjectType = "User"
)

func PossibleValuesForAADObjectType() []string {
	return []string{
		string(AADObjectTypeGroup),
		string(AADObjectTypeServicePrincipal),
		string(AADObjectTypeUser),
	}
}

func (s *AADObjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAADObjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAADObjectType(input string) (*AADObjectType, error) {
	vals := map[string]AADObjectType{
		"group":            AADObjectTypeGroup,
		"serviceprincipal": AADObjectTypeServicePrincipal,
		"user":             AADObjectTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AADObjectType(input)
	return &out, nil
}

type CheckNameAvailabilityParametersType string

const (
	CheckNameAvailabilityParametersTypeMicrosoftPointDataLakeAnalyticsAccounts CheckNameAvailabilityParametersType = "Microsoft.DataLakeAnalytics/accounts"
)

func PossibleValuesForCheckNameAvailabilityParametersType() []string {
	return []string{
		string(CheckNameAvailabilityParametersTypeMicrosoftPointDataLakeAnalyticsAccounts),
	}
}

func (s *CheckNameAvailabilityParametersType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckNameAvailabilityParametersType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckNameAvailabilityParametersType(input string) (*CheckNameAvailabilityParametersType, error) {
	vals := map[string]CheckNameAvailabilityParametersType{
		"microsoft.datalakeanalytics/accounts": CheckNameAvailabilityParametersTypeMicrosoftPointDataLakeAnalyticsAccounts,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityParametersType(input)
	return &out, nil
}

type DataLakeAnalyticsAccountState string

const (
	DataLakeAnalyticsAccountStateActive    DataLakeAnalyticsAccountState = "Active"
	DataLakeAnalyticsAccountStateSuspended DataLakeAnalyticsAccountState = "Suspended"
)

func PossibleValuesForDataLakeAnalyticsAccountState() []string {
	return []string{
		string(DataLakeAnalyticsAccountStateActive),
		string(DataLakeAnalyticsAccountStateSuspended),
	}
}

func (s *DataLakeAnalyticsAccountState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataLakeAnalyticsAccountState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataLakeAnalyticsAccountState(input string) (*DataLakeAnalyticsAccountState, error) {
	vals := map[string]DataLakeAnalyticsAccountState{
		"active":    DataLakeAnalyticsAccountStateActive,
		"suspended": DataLakeAnalyticsAccountStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataLakeAnalyticsAccountState(input)
	return &out, nil
}

type DataLakeAnalyticsAccountStatus string

const (
	DataLakeAnalyticsAccountStatusCanceled   DataLakeAnalyticsAccountStatus = "Canceled"
	DataLakeAnalyticsAccountStatusCreating   DataLakeAnalyticsAccountStatus = "Creating"
	DataLakeAnalyticsAccountStatusDeleted    DataLakeAnalyticsAccountStatus = "Deleted"
	DataLakeAnalyticsAccountStatusDeleting   DataLakeAnalyticsAccountStatus = "Deleting"
	DataLakeAnalyticsAccountStatusFailed     DataLakeAnalyticsAccountStatus = "Failed"
	DataLakeAnalyticsAccountStatusPatching   DataLakeAnalyticsAccountStatus = "Patching"
	DataLakeAnalyticsAccountStatusResuming   DataLakeAnalyticsAccountStatus = "Resuming"
	DataLakeAnalyticsAccountStatusRunning    DataLakeAnalyticsAccountStatus = "Running"
	DataLakeAnalyticsAccountStatusSucceeded  DataLakeAnalyticsAccountStatus = "Succeeded"
	DataLakeAnalyticsAccountStatusSuspending DataLakeAnalyticsAccountStatus = "Suspending"
	DataLakeAnalyticsAccountStatusUndeleting DataLakeAnalyticsAccountStatus = "Undeleting"
)

func PossibleValuesForDataLakeAnalyticsAccountStatus() []string {
	return []string{
		string(DataLakeAnalyticsAccountStatusCanceled),
		string(DataLakeAnalyticsAccountStatusCreating),
		string(DataLakeAnalyticsAccountStatusDeleted),
		string(DataLakeAnalyticsAccountStatusDeleting),
		string(DataLakeAnalyticsAccountStatusFailed),
		string(DataLakeAnalyticsAccountStatusPatching),
		string(DataLakeAnalyticsAccountStatusResuming),
		string(DataLakeAnalyticsAccountStatusRunning),
		string(DataLakeAnalyticsAccountStatusSucceeded),
		string(DataLakeAnalyticsAccountStatusSuspending),
		string(DataLakeAnalyticsAccountStatusUndeleting),
	}
}

func (s *DataLakeAnalyticsAccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataLakeAnalyticsAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataLakeAnalyticsAccountStatus(input string) (*DataLakeAnalyticsAccountStatus, error) {
	vals := map[string]DataLakeAnalyticsAccountStatus{
		"canceled":   DataLakeAnalyticsAccountStatusCanceled,
		"creating":   DataLakeAnalyticsAccountStatusCreating,
		"deleted":    DataLakeAnalyticsAccountStatusDeleted,
		"deleting":   DataLakeAnalyticsAccountStatusDeleting,
		"failed":     DataLakeAnalyticsAccountStatusFailed,
		"patching":   DataLakeAnalyticsAccountStatusPatching,
		"resuming":   DataLakeAnalyticsAccountStatusResuming,
		"running":    DataLakeAnalyticsAccountStatusRunning,
		"succeeded":  DataLakeAnalyticsAccountStatusSucceeded,
		"suspending": DataLakeAnalyticsAccountStatusSuspending,
		"undeleting": DataLakeAnalyticsAccountStatusUndeleting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataLakeAnalyticsAccountStatus(input)
	return &out, nil
}

type DebugDataAccessLevel string

const (
	DebugDataAccessLevelAll      DebugDataAccessLevel = "All"
	DebugDataAccessLevelCustomer DebugDataAccessLevel = "Customer"
	DebugDataAccessLevelNone     DebugDataAccessLevel = "None"
)

func PossibleValuesForDebugDataAccessLevel() []string {
	return []string{
		string(DebugDataAccessLevelAll),
		string(DebugDataAccessLevelCustomer),
		string(DebugDataAccessLevelNone),
	}
}

func (s *DebugDataAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDebugDataAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDebugDataAccessLevel(input string) (*DebugDataAccessLevel, error) {
	vals := map[string]DebugDataAccessLevel{
		"all":      DebugDataAccessLevelAll,
		"customer": DebugDataAccessLevelCustomer,
		"none":     DebugDataAccessLevelNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DebugDataAccessLevel(input)
	return &out, nil
}

type FirewallAllowAzureIPsState string

const (
	FirewallAllowAzureIPsStateDisabled FirewallAllowAzureIPsState = "Disabled"
	FirewallAllowAzureIPsStateEnabled  FirewallAllowAzureIPsState = "Enabled"
)

func PossibleValuesForFirewallAllowAzureIPsState() []string {
	return []string{
		string(FirewallAllowAzureIPsStateDisabled),
		string(FirewallAllowAzureIPsStateEnabled),
	}
}

func (s *FirewallAllowAzureIPsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallAllowAzureIPsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallAllowAzureIPsState(input string) (*FirewallAllowAzureIPsState, error) {
	vals := map[string]FirewallAllowAzureIPsState{
		"disabled": FirewallAllowAzureIPsStateDisabled,
		"enabled":  FirewallAllowAzureIPsStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallAllowAzureIPsState(input)
	return &out, nil
}

type FirewallState string

const (
	FirewallStateDisabled FirewallState = "Disabled"
	FirewallStateEnabled  FirewallState = "Enabled"
)

func PossibleValuesForFirewallState() []string {
	return []string{
		string(FirewallStateDisabled),
		string(FirewallStateEnabled),
	}
}

func (s *FirewallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallState(input string) (*FirewallState, error) {
	vals := map[string]FirewallState{
		"disabled": FirewallStateDisabled,
		"enabled":  FirewallStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallState(input)
	return &out, nil
}

type NestedResourceProvisioningState string

const (
	NestedResourceProvisioningStateCanceled  NestedResourceProvisioningState = "Canceled"
	NestedResourceProvisioningStateFailed    NestedResourceProvisioningState = "Failed"
	NestedResourceProvisioningStateSucceeded NestedResourceProvisioningState = "Succeeded"
)

func PossibleValuesForNestedResourceProvisioningState() []string {
	return []string{
		string(NestedResourceProvisioningStateCanceled),
		string(NestedResourceProvisioningStateFailed),
		string(NestedResourceProvisioningStateSucceeded),
	}
}

func (s *NestedResourceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNestedResourceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNestedResourceProvisioningState(input string) (*NestedResourceProvisioningState, error) {
	vals := map[string]NestedResourceProvisioningState{
		"canceled":  NestedResourceProvisioningStateCanceled,
		"failed":    NestedResourceProvisioningStateFailed,
		"succeeded": NestedResourceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NestedResourceProvisioningState(input)
	return &out, nil
}

type TierType string

const (
	TierTypeCommitmentFiveHundredAUHours         TierType = "Commitment_500AUHours"
	TierTypeCommitmentFiveHundredThousandAUHours TierType = "Commitment_500000AUHours"
	TierTypeCommitmentFiveThousandAUHours        TierType = "Commitment_5000AUHours"
	TierTypeCommitmentFiveZeroThousandAUHours    TierType = "Commitment_50000AUHours"
	TierTypeCommitmentOneHundredAUHours          TierType = "Commitment_100AUHours"
	TierTypeCommitmentOneHundredThousandAUHours  TierType = "Commitment_100000AUHours"
	TierTypeCommitmentOneThousandAUHours         TierType = "Commitment_1000AUHours"
	TierTypeCommitmentOneZeroThousandAUHours     TierType = "Commitment_10000AUHours"
	TierTypeConsumption                          TierType = "Consumption"
)

func PossibleValuesForTierType() []string {
	return []string{
		string(TierTypeCommitmentFiveHundredAUHours),
		string(TierTypeCommitmentFiveHundredThousandAUHours),
		string(TierTypeCommitmentFiveThousandAUHours),
		string(TierTypeCommitmentFiveZeroThousandAUHours),
		string(TierTypeCommitmentOneHundredAUHours),
		string(TierTypeCommitmentOneHundredThousandAUHours),
		string(TierTypeCommitmentOneThousandAUHours),
		string(TierTypeCommitmentOneZeroThousandAUHours),
		string(TierTypeConsumption),
	}
}

func (s *TierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTierType(input string) (*TierType, error) {
	vals := map[string]TierType{
		"commitment_500auhours":    TierTypeCommitmentFiveHundredAUHours,
		"commitment_500000auhours": TierTypeCommitmentFiveHundredThousandAUHours,
		"commitment_5000auhours":   TierTypeCommitmentFiveThousandAUHours,
		"commitment_50000auhours":  TierTypeCommitmentFiveZeroThousandAUHours,
		"commitment_100auhours":    TierTypeCommitmentOneHundredAUHours,
		"commitment_100000auhours": TierTypeCommitmentOneHundredThousandAUHours,
		"commitment_1000auhours":   TierTypeCommitmentOneThousandAUHours,
		"commitment_10000auhours":  TierTypeCommitmentOneZeroThousandAUHours,
		"consumption":              TierTypeConsumption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TierType(input)
	return &out, nil
}

type VirtualNetworkRuleState string

const (
	VirtualNetworkRuleStateActive               VirtualNetworkRuleState = "Active"
	VirtualNetworkRuleStateFailed               VirtualNetworkRuleState = "Failed"
	VirtualNetworkRuleStateNetworkSourceDeleted VirtualNetworkRuleState = "NetworkSourceDeleted"
)

func PossibleValuesForVirtualNetworkRuleState() []string {
	return []string{
		string(VirtualNetworkRuleStateActive),
		string(VirtualNetworkRuleStateFailed),
		string(VirtualNetworkRuleStateNetworkSourceDeleted),
	}
}

func (s *VirtualNetworkRuleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkRuleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkRuleState(input string) (*VirtualNetworkRuleState, error) {
	vals := map[string]VirtualNetworkRuleState{
		"active":               VirtualNetworkRuleStateActive,
		"failed":               VirtualNetworkRuleStateFailed,
		"networksourcedeleted": VirtualNetworkRuleStateNetworkSourceDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkRuleState(input)
	return &out, nil
}
