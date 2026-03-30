package baremetalmachines

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

type BareMetalMachineCordonStatus string

const (
	BareMetalMachineCordonStatusCordoned   BareMetalMachineCordonStatus = "Cordoned"
	BareMetalMachineCordonStatusUncordoned BareMetalMachineCordonStatus = "Uncordoned"
)

func PossibleValuesForBareMetalMachineCordonStatus() []string {
	return []string{
		string(BareMetalMachineCordonStatusCordoned),
		string(BareMetalMachineCordonStatusUncordoned),
	}
}

func (s *BareMetalMachineCordonStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineCordonStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineCordonStatus(input string) (*BareMetalMachineCordonStatus, error) {
	vals := map[string]BareMetalMachineCordonStatus{
		"cordoned":   BareMetalMachineCordonStatusCordoned,
		"uncordoned": BareMetalMachineCordonStatusUncordoned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineCordonStatus(input)
	return &out, nil
}

type BareMetalMachineDetailedStatus string

const (
	BareMetalMachineDetailedStatusAvailable      BareMetalMachineDetailedStatus = "Available"
	BareMetalMachineDetailedStatusDeprovisioning BareMetalMachineDetailedStatus = "Deprovisioning"
	BareMetalMachineDetailedStatusError          BareMetalMachineDetailedStatus = "Error"
	BareMetalMachineDetailedStatusPreparing      BareMetalMachineDetailedStatus = "Preparing"
	BareMetalMachineDetailedStatusProvisioned    BareMetalMachineDetailedStatus = "Provisioned"
	BareMetalMachineDetailedStatusProvisioning   BareMetalMachineDetailedStatus = "Provisioning"
)

func PossibleValuesForBareMetalMachineDetailedStatus() []string {
	return []string{
		string(BareMetalMachineDetailedStatusAvailable),
		string(BareMetalMachineDetailedStatusDeprovisioning),
		string(BareMetalMachineDetailedStatusError),
		string(BareMetalMachineDetailedStatusPreparing),
		string(BareMetalMachineDetailedStatusProvisioned),
		string(BareMetalMachineDetailedStatusProvisioning),
	}
}

func (s *BareMetalMachineDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineDetailedStatus(input string) (*BareMetalMachineDetailedStatus, error) {
	vals := map[string]BareMetalMachineDetailedStatus{
		"available":      BareMetalMachineDetailedStatusAvailable,
		"deprovisioning": BareMetalMachineDetailedStatusDeprovisioning,
		"error":          BareMetalMachineDetailedStatusError,
		"preparing":      BareMetalMachineDetailedStatusPreparing,
		"provisioned":    BareMetalMachineDetailedStatusProvisioned,
		"provisioning":   BareMetalMachineDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineDetailedStatus(input)
	return &out, nil
}

type BareMetalMachineEvacuate string

const (
	BareMetalMachineEvacuateFalse BareMetalMachineEvacuate = "False"
	BareMetalMachineEvacuateTrue  BareMetalMachineEvacuate = "True"
)

func PossibleValuesForBareMetalMachineEvacuate() []string {
	return []string{
		string(BareMetalMachineEvacuateFalse),
		string(BareMetalMachineEvacuateTrue),
	}
}

func (s *BareMetalMachineEvacuate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineEvacuate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineEvacuate(input string) (*BareMetalMachineEvacuate, error) {
	vals := map[string]BareMetalMachineEvacuate{
		"false": BareMetalMachineEvacuateFalse,
		"true":  BareMetalMachineEvacuateTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineEvacuate(input)
	return &out, nil
}

type BareMetalMachineHardwareValidationResult string

const (
	BareMetalMachineHardwareValidationResultFail BareMetalMachineHardwareValidationResult = "Fail"
	BareMetalMachineHardwareValidationResultPass BareMetalMachineHardwareValidationResult = "Pass"
)

func PossibleValuesForBareMetalMachineHardwareValidationResult() []string {
	return []string{
		string(BareMetalMachineHardwareValidationResultFail),
		string(BareMetalMachineHardwareValidationResultPass),
	}
}

func (s *BareMetalMachineHardwareValidationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineHardwareValidationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineHardwareValidationResult(input string) (*BareMetalMachineHardwareValidationResult, error) {
	vals := map[string]BareMetalMachineHardwareValidationResult{
		"fail": BareMetalMachineHardwareValidationResultFail,
		"pass": BareMetalMachineHardwareValidationResultPass,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineHardwareValidationResult(input)
	return &out, nil
}

type BareMetalMachinePowerState string

const (
	BareMetalMachinePowerStateOff BareMetalMachinePowerState = "Off"
	BareMetalMachinePowerStateOn  BareMetalMachinePowerState = "On"
)

func PossibleValuesForBareMetalMachinePowerState() []string {
	return []string{
		string(BareMetalMachinePowerStateOff),
		string(BareMetalMachinePowerStateOn),
	}
}

func (s *BareMetalMachinePowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachinePowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachinePowerState(input string) (*BareMetalMachinePowerState, error) {
	vals := map[string]BareMetalMachinePowerState{
		"off": BareMetalMachinePowerStateOff,
		"on":  BareMetalMachinePowerStateOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachinePowerState(input)
	return &out, nil
}

type BareMetalMachineProvisioningState string

const (
	BareMetalMachineProvisioningStateAccepted     BareMetalMachineProvisioningState = "Accepted"
	BareMetalMachineProvisioningStateCanceled     BareMetalMachineProvisioningState = "Canceled"
	BareMetalMachineProvisioningStateFailed       BareMetalMachineProvisioningState = "Failed"
	BareMetalMachineProvisioningStateProvisioning BareMetalMachineProvisioningState = "Provisioning"
	BareMetalMachineProvisioningStateSucceeded    BareMetalMachineProvisioningState = "Succeeded"
)

func PossibleValuesForBareMetalMachineProvisioningState() []string {
	return []string{
		string(BareMetalMachineProvisioningStateAccepted),
		string(BareMetalMachineProvisioningStateCanceled),
		string(BareMetalMachineProvisioningStateFailed),
		string(BareMetalMachineProvisioningStateProvisioning),
		string(BareMetalMachineProvisioningStateSucceeded),
	}
}

func (s *BareMetalMachineProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineProvisioningState(input string) (*BareMetalMachineProvisioningState, error) {
	vals := map[string]BareMetalMachineProvisioningState{
		"accepted":     BareMetalMachineProvisioningStateAccepted,
		"canceled":     BareMetalMachineProvisioningStateCanceled,
		"failed":       BareMetalMachineProvisioningStateFailed,
		"provisioning": BareMetalMachineProvisioningStateProvisioning,
		"succeeded":    BareMetalMachineProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineProvisioningState(input)
	return &out, nil
}

type BareMetalMachineReadyState string

const (
	BareMetalMachineReadyStateFalse BareMetalMachineReadyState = "False"
	BareMetalMachineReadyStateTrue  BareMetalMachineReadyState = "True"
)

func PossibleValuesForBareMetalMachineReadyState() []string {
	return []string{
		string(BareMetalMachineReadyStateFalse),
		string(BareMetalMachineReadyStateTrue),
	}
}

func (s *BareMetalMachineReadyState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineReadyState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineReadyState(input string) (*BareMetalMachineReadyState, error) {
	vals := map[string]BareMetalMachineReadyState{
		"false": BareMetalMachineReadyStateFalse,
		"true":  BareMetalMachineReadyStateTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineReadyState(input)
	return &out, nil
}

type BareMetalMachineReplaceSafeguardMode string

const (
	BareMetalMachineReplaceSafeguardModeAll  BareMetalMachineReplaceSafeguardMode = "All"
	BareMetalMachineReplaceSafeguardModeNone BareMetalMachineReplaceSafeguardMode = "None"
)

func PossibleValuesForBareMetalMachineReplaceSafeguardMode() []string {
	return []string{
		string(BareMetalMachineReplaceSafeguardModeAll),
		string(BareMetalMachineReplaceSafeguardModeNone),
	}
}

func (s *BareMetalMachineReplaceSafeguardMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineReplaceSafeguardMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineReplaceSafeguardMode(input string) (*BareMetalMachineReplaceSafeguardMode, error) {
	vals := map[string]BareMetalMachineReplaceSafeguardMode{
		"all":  BareMetalMachineReplaceSafeguardModeAll,
		"none": BareMetalMachineReplaceSafeguardModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineReplaceSafeguardMode(input)
	return &out, nil
}

type BareMetalMachineReplaceStoragePolicy string

const (
	BareMetalMachineReplaceStoragePolicyDiscardAll BareMetalMachineReplaceStoragePolicy = "DiscardAll"
	BareMetalMachineReplaceStoragePolicyPreserve   BareMetalMachineReplaceStoragePolicy = "Preserve"
)

func PossibleValuesForBareMetalMachineReplaceStoragePolicy() []string {
	return []string{
		string(BareMetalMachineReplaceStoragePolicyDiscardAll),
		string(BareMetalMachineReplaceStoragePolicyPreserve),
	}
}

func (s *BareMetalMachineReplaceStoragePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineReplaceStoragePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineReplaceStoragePolicy(input string) (*BareMetalMachineReplaceStoragePolicy, error) {
	vals := map[string]BareMetalMachineReplaceStoragePolicy{
		"discardall": BareMetalMachineReplaceStoragePolicyDiscardAll,
		"preserve":   BareMetalMachineReplaceStoragePolicyPreserve,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineReplaceStoragePolicy(input)
	return &out, nil
}

type BareMetalMachineSkipShutdown string

const (
	BareMetalMachineSkipShutdownFalse BareMetalMachineSkipShutdown = "False"
	BareMetalMachineSkipShutdownTrue  BareMetalMachineSkipShutdown = "True"
)

func PossibleValuesForBareMetalMachineSkipShutdown() []string {
	return []string{
		string(BareMetalMachineSkipShutdownFalse),
		string(BareMetalMachineSkipShutdownTrue),
	}
}

func (s *BareMetalMachineSkipShutdown) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineSkipShutdown(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineSkipShutdown(input string) (*BareMetalMachineSkipShutdown, error) {
	vals := map[string]BareMetalMachineSkipShutdown{
		"false": BareMetalMachineSkipShutdownFalse,
		"true":  BareMetalMachineSkipShutdownTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineSkipShutdown(input)
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
