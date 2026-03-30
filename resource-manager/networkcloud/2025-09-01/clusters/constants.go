package clusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionStateStatus string

const (
	ActionStateStatusCompleted  ActionStateStatus = "Completed"
	ActionStateStatusFailed     ActionStateStatus = "Failed"
	ActionStateStatusInProgress ActionStateStatus = "InProgress"
)

func PossibleValuesForActionStateStatus() []string {
	return []string{
		string(ActionStateStatusCompleted),
		string(ActionStateStatusFailed),
		string(ActionStateStatusInProgress),
	}
}

func (s *ActionStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionStateStatus(input string) (*ActionStateStatus, error) {
	vals := map[string]ActionStateStatus{
		"completed":  ActionStateStatusCompleted,
		"failed":     ActionStateStatusFailed,
		"inprogress": ActionStateStatusInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionStateStatus(input)
	return &out, nil
}

type ClusterConnectionStatus string

const (
	ClusterConnectionStatusConnected    ClusterConnectionStatus = "Connected"
	ClusterConnectionStatusDisconnected ClusterConnectionStatus = "Disconnected"
	ClusterConnectionStatusTimeout      ClusterConnectionStatus = "Timeout"
	ClusterConnectionStatusUndefined    ClusterConnectionStatus = "Undefined"
)

func PossibleValuesForClusterConnectionStatus() []string {
	return []string{
		string(ClusterConnectionStatusConnected),
		string(ClusterConnectionStatusDisconnected),
		string(ClusterConnectionStatusTimeout),
		string(ClusterConnectionStatusUndefined),
	}
}

func (s *ClusterConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterConnectionStatus(input string) (*ClusterConnectionStatus, error) {
	vals := map[string]ClusterConnectionStatus{
		"connected":    ClusterConnectionStatusConnected,
		"disconnected": ClusterConnectionStatusDisconnected,
		"timeout":      ClusterConnectionStatusTimeout,
		"undefined":    ClusterConnectionStatusUndefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterConnectionStatus(input)
	return &out, nil
}

type ClusterContinueUpdateVersionMachineGroupTargetingMode string

const (
	ClusterContinueUpdateVersionMachineGroupTargetingModeAlphaByRack ClusterContinueUpdateVersionMachineGroupTargetingMode = "AlphaByRack"
)

func PossibleValuesForClusterContinueUpdateVersionMachineGroupTargetingMode() []string {
	return []string{
		string(ClusterContinueUpdateVersionMachineGroupTargetingModeAlphaByRack),
	}
}

func (s *ClusterContinueUpdateVersionMachineGroupTargetingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterContinueUpdateVersionMachineGroupTargetingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterContinueUpdateVersionMachineGroupTargetingMode(input string) (*ClusterContinueUpdateVersionMachineGroupTargetingMode, error) {
	vals := map[string]ClusterContinueUpdateVersionMachineGroupTargetingMode{
		"alphabyrack": ClusterContinueUpdateVersionMachineGroupTargetingModeAlphaByRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterContinueUpdateVersionMachineGroupTargetingMode(input)
	return &out, nil
}

type ClusterDetailedStatus string

const (
	ClusterDetailedStatusDegraded          ClusterDetailedStatus = "Degraded"
	ClusterDetailedStatusDeleting          ClusterDetailedStatus = "Deleting"
	ClusterDetailedStatusDeploying         ClusterDetailedStatus = "Deploying"
	ClusterDetailedStatusDisconnected      ClusterDetailedStatus = "Disconnected"
	ClusterDetailedStatusFailed            ClusterDetailedStatus = "Failed"
	ClusterDetailedStatusPendingDeployment ClusterDetailedStatus = "PendingDeployment"
	ClusterDetailedStatusRunning           ClusterDetailedStatus = "Running"
	ClusterDetailedStatusUpdatePaused      ClusterDetailedStatus = "UpdatePaused"
	ClusterDetailedStatusUpdating          ClusterDetailedStatus = "Updating"
)

func PossibleValuesForClusterDetailedStatus() []string {
	return []string{
		string(ClusterDetailedStatusDegraded),
		string(ClusterDetailedStatusDeleting),
		string(ClusterDetailedStatusDeploying),
		string(ClusterDetailedStatusDisconnected),
		string(ClusterDetailedStatusFailed),
		string(ClusterDetailedStatusPendingDeployment),
		string(ClusterDetailedStatusRunning),
		string(ClusterDetailedStatusUpdatePaused),
		string(ClusterDetailedStatusUpdating),
	}
}

func (s *ClusterDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterDetailedStatus(input string) (*ClusterDetailedStatus, error) {
	vals := map[string]ClusterDetailedStatus{
		"degraded":          ClusterDetailedStatusDegraded,
		"deleting":          ClusterDetailedStatusDeleting,
		"deploying":         ClusterDetailedStatusDeploying,
		"disconnected":      ClusterDetailedStatusDisconnected,
		"failed":            ClusterDetailedStatusFailed,
		"pendingdeployment": ClusterDetailedStatusPendingDeployment,
		"running":           ClusterDetailedStatusRunning,
		"updatepaused":      ClusterDetailedStatusUpdatePaused,
		"updating":          ClusterDetailedStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterDetailedStatus(input)
	return &out, nil
}

type ClusterManagerConnectionStatus string

const (
	ClusterManagerConnectionStatusConnected   ClusterManagerConnectionStatus = "Connected"
	ClusterManagerConnectionStatusUnreachable ClusterManagerConnectionStatus = "Unreachable"
)

func PossibleValuesForClusterManagerConnectionStatus() []string {
	return []string{
		string(ClusterManagerConnectionStatusConnected),
		string(ClusterManagerConnectionStatusUnreachable),
	}
}

func (s *ClusterManagerConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterManagerConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterManagerConnectionStatus(input string) (*ClusterManagerConnectionStatus, error) {
	vals := map[string]ClusterManagerConnectionStatus{
		"connected":   ClusterManagerConnectionStatusConnected,
		"unreachable": ClusterManagerConnectionStatusUnreachable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerConnectionStatus(input)
	return &out, nil
}

type ClusterProvisioningState string

const (
	ClusterProvisioningStateAccepted   ClusterProvisioningState = "Accepted"
	ClusterProvisioningStateCanceled   ClusterProvisioningState = "Canceled"
	ClusterProvisioningStateFailed     ClusterProvisioningState = "Failed"
	ClusterProvisioningStateSucceeded  ClusterProvisioningState = "Succeeded"
	ClusterProvisioningStateUpdating   ClusterProvisioningState = "Updating"
	ClusterProvisioningStateValidating ClusterProvisioningState = "Validating"
)

func PossibleValuesForClusterProvisioningState() []string {
	return []string{
		string(ClusterProvisioningStateAccepted),
		string(ClusterProvisioningStateCanceled),
		string(ClusterProvisioningStateFailed),
		string(ClusterProvisioningStateSucceeded),
		string(ClusterProvisioningStateUpdating),
		string(ClusterProvisioningStateValidating),
	}
}

func (s *ClusterProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterProvisioningState(input string) (*ClusterProvisioningState, error) {
	vals := map[string]ClusterProvisioningState{
		"accepted":   ClusterProvisioningStateAccepted,
		"canceled":   ClusterProvisioningStateCanceled,
		"failed":     ClusterProvisioningStateFailed,
		"succeeded":  ClusterProvisioningStateSucceeded,
		"updating":   ClusterProvisioningStateUpdating,
		"validating": ClusterProvisioningStateValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterProvisioningState(input)
	return &out, nil
}

type ClusterScanRuntimeParametersScanActivity string

const (
	ClusterScanRuntimeParametersScanActivityScan ClusterScanRuntimeParametersScanActivity = "Scan"
	ClusterScanRuntimeParametersScanActivitySkip ClusterScanRuntimeParametersScanActivity = "Skip"
)

func PossibleValuesForClusterScanRuntimeParametersScanActivity() []string {
	return []string{
		string(ClusterScanRuntimeParametersScanActivityScan),
		string(ClusterScanRuntimeParametersScanActivitySkip),
	}
}

func (s *ClusterScanRuntimeParametersScanActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterScanRuntimeParametersScanActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterScanRuntimeParametersScanActivity(input string) (*ClusterScanRuntimeParametersScanActivity, error) {
	vals := map[string]ClusterScanRuntimeParametersScanActivity{
		"scan": ClusterScanRuntimeParametersScanActivityScan,
		"skip": ClusterScanRuntimeParametersScanActivitySkip,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterScanRuntimeParametersScanActivity(input)
	return &out, nil
}

type ClusterSecretArchiveEnabled string

const (
	ClusterSecretArchiveEnabledFalse ClusterSecretArchiveEnabled = "False"
	ClusterSecretArchiveEnabledTrue  ClusterSecretArchiveEnabled = "True"
)

func PossibleValuesForClusterSecretArchiveEnabled() []string {
	return []string{
		string(ClusterSecretArchiveEnabledFalse),
		string(ClusterSecretArchiveEnabledTrue),
	}
}

func (s *ClusterSecretArchiveEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterSecretArchiveEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterSecretArchiveEnabled(input string) (*ClusterSecretArchiveEnabled, error) {
	vals := map[string]ClusterSecretArchiveEnabled{
		"false": ClusterSecretArchiveEnabledFalse,
		"true":  ClusterSecretArchiveEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterSecretArchiveEnabled(input)
	return &out, nil
}

type ClusterType string

const (
	ClusterTypeMultiRack  ClusterType = "MultiRack"
	ClusterTypeSingleRack ClusterType = "SingleRack"
)

func PossibleValuesForClusterType() []string {
	return []string{
		string(ClusterTypeMultiRack),
		string(ClusterTypeSingleRack),
	}
}

func (s *ClusterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterType(input string) (*ClusterType, error) {
	vals := map[string]ClusterType{
		"multirack":  ClusterTypeMultiRack,
		"singlerack": ClusterTypeSingleRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterType(input)
	return &out, nil
}

type ClusterUpdateStrategyType string

const (
	ClusterUpdateStrategyTypePauseAfterRack ClusterUpdateStrategyType = "PauseAfterRack"
	ClusterUpdateStrategyTypeRack           ClusterUpdateStrategyType = "Rack"
)

func PossibleValuesForClusterUpdateStrategyType() []string {
	return []string{
		string(ClusterUpdateStrategyTypePauseAfterRack),
		string(ClusterUpdateStrategyTypeRack),
	}
}

func (s *ClusterUpdateStrategyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterUpdateStrategyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterUpdateStrategyType(input string) (*ClusterUpdateStrategyType, error) {
	vals := map[string]ClusterUpdateStrategyType{
		"pauseafterrack": ClusterUpdateStrategyTypePauseAfterRack,
		"rack":           ClusterUpdateStrategyTypeRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterUpdateStrategyType(input)
	return &out, nil
}

type CommandOutputType string

const (
	CommandOutputTypeBareMetalMachineRunCommand                CommandOutputType = "BareMetalMachineRunCommand"
	CommandOutputTypeBareMetalMachineRunDataExtracts           CommandOutputType = "BareMetalMachineRunDataExtracts"
	CommandOutputTypeBareMetalMachineRunDataExtractsRestricted CommandOutputType = "BareMetalMachineRunDataExtractsRestricted"
	CommandOutputTypeBareMetalMachineRunReadCommands           CommandOutputType = "BareMetalMachineRunReadCommands"
	CommandOutputTypeStorageRunReadCommands                    CommandOutputType = "StorageRunReadCommands"
)

func PossibleValuesForCommandOutputType() []string {
	return []string{
		string(CommandOutputTypeBareMetalMachineRunCommand),
		string(CommandOutputTypeBareMetalMachineRunDataExtracts),
		string(CommandOutputTypeBareMetalMachineRunDataExtractsRestricted),
		string(CommandOutputTypeBareMetalMachineRunReadCommands),
		string(CommandOutputTypeStorageRunReadCommands),
	}
}

func (s *CommandOutputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommandOutputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommandOutputType(input string) (*CommandOutputType, error) {
	vals := map[string]CommandOutputType{
		"baremetalmachineruncommand":                CommandOutputTypeBareMetalMachineRunCommand,
		"baremetalmachinerundataextracts":           CommandOutputTypeBareMetalMachineRunDataExtracts,
		"baremetalmachinerundataextractsrestricted": CommandOutputTypeBareMetalMachineRunDataExtractsRestricted,
		"baremetalmachinerunreadcommands":           CommandOutputTypeBareMetalMachineRunReadCommands,
		"storagerunreadcommands":                    CommandOutputTypeStorageRunReadCommands,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommandOutputType(input)
	return &out, nil
}

type ControlImpact string

const (
	ControlImpactFalse ControlImpact = "False"
	ControlImpactTrue  ControlImpact = "True"
)

func PossibleValuesForControlImpact() []string {
	return []string{
		string(ControlImpactFalse),
		string(ControlImpactTrue),
	}
}

func (s *ControlImpact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseControlImpact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseControlImpact(input string) (*ControlImpact, error) {
	vals := map[string]ControlImpact{
		"false": ControlImpactFalse,
		"true":  ControlImpactTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ControlImpact(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
	ExtendedLocationTypeEdgeZone       ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
		string(ExtendedLocationTypeEdgeZone),
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
		"edgezone":       ExtendedLocationTypeEdgeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type ManagedServiceIdentitySelectorType string

const (
	ManagedServiceIdentitySelectorTypeSystemAssignedIdentity ManagedServiceIdentitySelectorType = "SystemAssignedIdentity"
	ManagedServiceIdentitySelectorTypeUserAssignedIdentity   ManagedServiceIdentitySelectorType = "UserAssignedIdentity"
)

func PossibleValuesForManagedServiceIdentitySelectorType() []string {
	return []string{
		string(ManagedServiceIdentitySelectorTypeSystemAssignedIdentity),
		string(ManagedServiceIdentitySelectorTypeUserAssignedIdentity),
	}
}

func (s *ManagedServiceIdentitySelectorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedServiceIdentitySelectorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedServiceIdentitySelectorType(input string) (*ManagedServiceIdentitySelectorType, error) {
	vals := map[string]ManagedServiceIdentitySelectorType{
		"systemassignedidentity": ManagedServiceIdentitySelectorTypeSystemAssignedIdentity,
		"userassignedidentity":   ManagedServiceIdentitySelectorTypeUserAssignedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedServiceIdentitySelectorType(input)
	return &out, nil
}

type RuntimeProtectionEnforcementLevel string

const (
	RuntimeProtectionEnforcementLevelAudit    RuntimeProtectionEnforcementLevel = "Audit"
	RuntimeProtectionEnforcementLevelDisabled RuntimeProtectionEnforcementLevel = "Disabled"
	RuntimeProtectionEnforcementLevelOnDemand RuntimeProtectionEnforcementLevel = "OnDemand"
	RuntimeProtectionEnforcementLevelPassive  RuntimeProtectionEnforcementLevel = "Passive"
	RuntimeProtectionEnforcementLevelRealTime RuntimeProtectionEnforcementLevel = "RealTime"
)

func PossibleValuesForRuntimeProtectionEnforcementLevel() []string {
	return []string{
		string(RuntimeProtectionEnforcementLevelAudit),
		string(RuntimeProtectionEnforcementLevelDisabled),
		string(RuntimeProtectionEnforcementLevelOnDemand),
		string(RuntimeProtectionEnforcementLevelPassive),
		string(RuntimeProtectionEnforcementLevelRealTime),
	}
}

func (s *RuntimeProtectionEnforcementLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuntimeProtectionEnforcementLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuntimeProtectionEnforcementLevel(input string) (*RuntimeProtectionEnforcementLevel, error) {
	vals := map[string]RuntimeProtectionEnforcementLevel{
		"audit":    RuntimeProtectionEnforcementLevelAudit,
		"disabled": RuntimeProtectionEnforcementLevelDisabled,
		"ondemand": RuntimeProtectionEnforcementLevelOnDemand,
		"passive":  RuntimeProtectionEnforcementLevelPassive,
		"realtime": RuntimeProtectionEnforcementLevelRealTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuntimeProtectionEnforcementLevel(input)
	return &out, nil
}

type StepStateStatus string

const (
	StepStateStatusCompleted  StepStateStatus = "Completed"
	StepStateStatusFailed     StepStateStatus = "Failed"
	StepStateStatusInProgress StepStateStatus = "InProgress"
	StepStateStatusNotStarted StepStateStatus = "NotStarted"
)

func PossibleValuesForStepStateStatus() []string {
	return []string{
		string(StepStateStatusCompleted),
		string(StepStateStatusFailed),
		string(StepStateStatusInProgress),
		string(StepStateStatusNotStarted),
	}
}

func (s *StepStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStepStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStepStateStatus(input string) (*StepStateStatus, error) {
	vals := map[string]StepStateStatus{
		"completed":  StepStateStatusCompleted,
		"failed":     StepStateStatusFailed,
		"inprogress": StepStateStatusInProgress,
		"notstarted": StepStateStatusNotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StepStateStatus(input)
	return &out, nil
}

type ValidationThresholdGrouping string

const (
	ValidationThresholdGroupingPerCluster ValidationThresholdGrouping = "PerCluster"
	ValidationThresholdGroupingPerRack    ValidationThresholdGrouping = "PerRack"
)

func PossibleValuesForValidationThresholdGrouping() []string {
	return []string{
		string(ValidationThresholdGroupingPerCluster),
		string(ValidationThresholdGroupingPerRack),
	}
}

func (s *ValidationThresholdGrouping) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidationThresholdGrouping(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidationThresholdGrouping(input string) (*ValidationThresholdGrouping, error) {
	vals := map[string]ValidationThresholdGrouping{
		"percluster": ValidationThresholdGroupingPerCluster,
		"perrack":    ValidationThresholdGroupingPerRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationThresholdGrouping(input)
	return &out, nil
}

type ValidationThresholdType string

const (
	ValidationThresholdTypeCountSuccess   ValidationThresholdType = "CountSuccess"
	ValidationThresholdTypePercentSuccess ValidationThresholdType = "PercentSuccess"
)

func PossibleValuesForValidationThresholdType() []string {
	return []string{
		string(ValidationThresholdTypeCountSuccess),
		string(ValidationThresholdTypePercentSuccess),
	}
}

func (s *ValidationThresholdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidationThresholdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidationThresholdType(input string) (*ValidationThresholdType, error) {
	vals := map[string]ValidationThresholdType{
		"countsuccess":   ValidationThresholdTypeCountSuccess,
		"percentsuccess": ValidationThresholdTypePercentSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationThresholdType(input)
	return &out, nil
}

type VulnerabilityScanningSettingsContainerScan string

const (
	VulnerabilityScanningSettingsContainerScanDisabled VulnerabilityScanningSettingsContainerScan = "Disabled"
	VulnerabilityScanningSettingsContainerScanEnabled  VulnerabilityScanningSettingsContainerScan = "Enabled"
)

func PossibleValuesForVulnerabilityScanningSettingsContainerScan() []string {
	return []string{
		string(VulnerabilityScanningSettingsContainerScanDisabled),
		string(VulnerabilityScanningSettingsContainerScanEnabled),
	}
}

func (s *VulnerabilityScanningSettingsContainerScan) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVulnerabilityScanningSettingsContainerScan(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVulnerabilityScanningSettingsContainerScan(input string) (*VulnerabilityScanningSettingsContainerScan, error) {
	vals := map[string]VulnerabilityScanningSettingsContainerScan{
		"disabled": VulnerabilityScanningSettingsContainerScanDisabled,
		"enabled":  VulnerabilityScanningSettingsContainerScanEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VulnerabilityScanningSettingsContainerScan(input)
	return &out, nil
}

type WorkloadImpact string

const (
	WorkloadImpactFalse WorkloadImpact = "False"
	WorkloadImpactTrue  WorkloadImpact = "True"
)

func PossibleValuesForWorkloadImpact() []string {
	return []string{
		string(WorkloadImpactFalse),
		string(WorkloadImpactTrue),
	}
}

func (s *WorkloadImpact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadImpact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadImpact(input string) (*WorkloadImpact, error) {
	vals := map[string]WorkloadImpact{
		"false": WorkloadImpactFalse,
		"true":  WorkloadImpactTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadImpact(input)
	return &out, nil
}
