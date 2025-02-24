package networkclouds

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvertiseToFabric string

const (
	AdvertiseToFabricFalse AdvertiseToFabric = "False"
	AdvertiseToFabricTrue  AdvertiseToFabric = "True"
)

func PossibleValuesForAdvertiseToFabric() []string {
	return []string{
		string(AdvertiseToFabricFalse),
		string(AdvertiseToFabricTrue),
	}
}

func (s *AdvertiseToFabric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvertiseToFabric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvertiseToFabric(input string) (*AdvertiseToFabric, error) {
	vals := map[string]AdvertiseToFabric{
		"false": AdvertiseToFabricFalse,
		"true":  AdvertiseToFabricTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvertiseToFabric(input)
	return &out, nil
}

type AgentPoolDetailedStatus string

const (
	AgentPoolDetailedStatusAvailable    AgentPoolDetailedStatus = "Available"
	AgentPoolDetailedStatusError        AgentPoolDetailedStatus = "Error"
	AgentPoolDetailedStatusProvisioning AgentPoolDetailedStatus = "Provisioning"
)

func PossibleValuesForAgentPoolDetailedStatus() []string {
	return []string{
		string(AgentPoolDetailedStatusAvailable),
		string(AgentPoolDetailedStatusError),
		string(AgentPoolDetailedStatusProvisioning),
	}
}

func (s *AgentPoolDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolDetailedStatus(input string) (*AgentPoolDetailedStatus, error) {
	vals := map[string]AgentPoolDetailedStatus{
		"available":    AgentPoolDetailedStatusAvailable,
		"error":        AgentPoolDetailedStatusError,
		"provisioning": AgentPoolDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolDetailedStatus(input)
	return &out, nil
}

type AgentPoolMode string

const (
	AgentPoolModeNotApplicable AgentPoolMode = "NotApplicable"
	AgentPoolModeSystem        AgentPoolMode = "System"
	AgentPoolModeUser          AgentPoolMode = "User"
)

func PossibleValuesForAgentPoolMode() []string {
	return []string{
		string(AgentPoolModeNotApplicable),
		string(AgentPoolModeSystem),
		string(AgentPoolModeUser),
	}
}

func (s *AgentPoolMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolMode(input string) (*AgentPoolMode, error) {
	vals := map[string]AgentPoolMode{
		"notapplicable": AgentPoolModeNotApplicable,
		"system":        AgentPoolModeSystem,
		"user":          AgentPoolModeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolMode(input)
	return &out, nil
}

type AgentPoolProvisioningState string

const (
	AgentPoolProvisioningStateAccepted   AgentPoolProvisioningState = "Accepted"
	AgentPoolProvisioningStateCanceled   AgentPoolProvisioningState = "Canceled"
	AgentPoolProvisioningStateDeleting   AgentPoolProvisioningState = "Deleting"
	AgentPoolProvisioningStateFailed     AgentPoolProvisioningState = "Failed"
	AgentPoolProvisioningStateInProgress AgentPoolProvisioningState = "InProgress"
	AgentPoolProvisioningStateSucceeded  AgentPoolProvisioningState = "Succeeded"
	AgentPoolProvisioningStateUpdating   AgentPoolProvisioningState = "Updating"
)

func PossibleValuesForAgentPoolProvisioningState() []string {
	return []string{
		string(AgentPoolProvisioningStateAccepted),
		string(AgentPoolProvisioningStateCanceled),
		string(AgentPoolProvisioningStateDeleting),
		string(AgentPoolProvisioningStateFailed),
		string(AgentPoolProvisioningStateInProgress),
		string(AgentPoolProvisioningStateSucceeded),
		string(AgentPoolProvisioningStateUpdating),
	}
}

func (s *AgentPoolProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentPoolProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentPoolProvisioningState(input string) (*AgentPoolProvisioningState, error) {
	vals := map[string]AgentPoolProvisioningState{
		"accepted":   AgentPoolProvisioningStateAccepted,
		"canceled":   AgentPoolProvisioningStateCanceled,
		"deleting":   AgentPoolProvisioningStateDeleting,
		"failed":     AgentPoolProvisioningStateFailed,
		"inprogress": AgentPoolProvisioningStateInProgress,
		"succeeded":  AgentPoolProvisioningStateSucceeded,
		"updating":   AgentPoolProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolProvisioningState(input)
	return &out, nil
}

type AvailabilityLifecycle string

const (
	AvailabilityLifecycleGenerallyAvailable AvailabilityLifecycle = "GenerallyAvailable"
	AvailabilityLifecyclePreview            AvailabilityLifecycle = "Preview"
)

func PossibleValuesForAvailabilityLifecycle() []string {
	return []string{
		string(AvailabilityLifecycleGenerallyAvailable),
		string(AvailabilityLifecyclePreview),
	}
}

func (s *AvailabilityLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAvailabilityLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAvailabilityLifecycle(input string) (*AvailabilityLifecycle, error) {
	vals := map[string]AvailabilityLifecycle{
		"generallyavailable": AvailabilityLifecycleGenerallyAvailable,
		"preview":            AvailabilityLifecyclePreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AvailabilityLifecycle(input)
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

type BareMetalMachineKeySetDetailedStatus string

const (
	BareMetalMachineKeySetDetailedStatusAllActive   BareMetalMachineKeySetDetailedStatus = "AllActive"
	BareMetalMachineKeySetDetailedStatusAllInvalid  BareMetalMachineKeySetDetailedStatus = "AllInvalid"
	BareMetalMachineKeySetDetailedStatusSomeInvalid BareMetalMachineKeySetDetailedStatus = "SomeInvalid"
	BareMetalMachineKeySetDetailedStatusValidating  BareMetalMachineKeySetDetailedStatus = "Validating"
)

func PossibleValuesForBareMetalMachineKeySetDetailedStatus() []string {
	return []string{
		string(BareMetalMachineKeySetDetailedStatusAllActive),
		string(BareMetalMachineKeySetDetailedStatusAllInvalid),
		string(BareMetalMachineKeySetDetailedStatusSomeInvalid),
		string(BareMetalMachineKeySetDetailedStatusValidating),
	}
}

func (s *BareMetalMachineKeySetDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetDetailedStatus(input string) (*BareMetalMachineKeySetDetailedStatus, error) {
	vals := map[string]BareMetalMachineKeySetDetailedStatus{
		"allactive":   BareMetalMachineKeySetDetailedStatusAllActive,
		"allinvalid":  BareMetalMachineKeySetDetailedStatusAllInvalid,
		"someinvalid": BareMetalMachineKeySetDetailedStatusSomeInvalid,
		"validating":  BareMetalMachineKeySetDetailedStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetDetailedStatus(input)
	return &out, nil
}

type BareMetalMachineKeySetPrivilegeLevel string

const (
	BareMetalMachineKeySetPrivilegeLevelStandard  BareMetalMachineKeySetPrivilegeLevel = "Standard"
	BareMetalMachineKeySetPrivilegeLevelSuperuser BareMetalMachineKeySetPrivilegeLevel = "Superuser"
)

func PossibleValuesForBareMetalMachineKeySetPrivilegeLevel() []string {
	return []string{
		string(BareMetalMachineKeySetPrivilegeLevelStandard),
		string(BareMetalMachineKeySetPrivilegeLevelSuperuser),
	}
}

func (s *BareMetalMachineKeySetPrivilegeLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetPrivilegeLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetPrivilegeLevel(input string) (*BareMetalMachineKeySetPrivilegeLevel, error) {
	vals := map[string]BareMetalMachineKeySetPrivilegeLevel{
		"standard":  BareMetalMachineKeySetPrivilegeLevelStandard,
		"superuser": BareMetalMachineKeySetPrivilegeLevelSuperuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetPrivilegeLevel(input)
	return &out, nil
}

type BareMetalMachineKeySetProvisioningState string

const (
	BareMetalMachineKeySetProvisioningStateAccepted     BareMetalMachineKeySetProvisioningState = "Accepted"
	BareMetalMachineKeySetProvisioningStateCanceled     BareMetalMachineKeySetProvisioningState = "Canceled"
	BareMetalMachineKeySetProvisioningStateFailed       BareMetalMachineKeySetProvisioningState = "Failed"
	BareMetalMachineKeySetProvisioningStateProvisioning BareMetalMachineKeySetProvisioningState = "Provisioning"
	BareMetalMachineKeySetProvisioningStateSucceeded    BareMetalMachineKeySetProvisioningState = "Succeeded"
)

func PossibleValuesForBareMetalMachineKeySetProvisioningState() []string {
	return []string{
		string(BareMetalMachineKeySetProvisioningStateAccepted),
		string(BareMetalMachineKeySetProvisioningStateCanceled),
		string(BareMetalMachineKeySetProvisioningStateFailed),
		string(BareMetalMachineKeySetProvisioningStateProvisioning),
		string(BareMetalMachineKeySetProvisioningStateSucceeded),
	}
}

func (s *BareMetalMachineKeySetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetProvisioningState(input string) (*BareMetalMachineKeySetProvisioningState, error) {
	vals := map[string]BareMetalMachineKeySetProvisioningState{
		"accepted":     BareMetalMachineKeySetProvisioningStateAccepted,
		"canceled":     BareMetalMachineKeySetProvisioningStateCanceled,
		"failed":       BareMetalMachineKeySetProvisioningStateFailed,
		"provisioning": BareMetalMachineKeySetProvisioningStateProvisioning,
		"succeeded":    BareMetalMachineKeySetProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetProvisioningState(input)
	return &out, nil
}

type BareMetalMachineKeySetUserSetupStatus string

const (
	BareMetalMachineKeySetUserSetupStatusActive  BareMetalMachineKeySetUserSetupStatus = "Active"
	BareMetalMachineKeySetUserSetupStatusInvalid BareMetalMachineKeySetUserSetupStatus = "Invalid"
)

func PossibleValuesForBareMetalMachineKeySetUserSetupStatus() []string {
	return []string{
		string(BareMetalMachineKeySetUserSetupStatusActive),
		string(BareMetalMachineKeySetUserSetupStatusInvalid),
	}
}

func (s *BareMetalMachineKeySetUserSetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetUserSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetUserSetupStatus(input string) (*BareMetalMachineKeySetUserSetupStatus, error) {
	vals := map[string]BareMetalMachineKeySetUserSetupStatus{
		"active":  BareMetalMachineKeySetUserSetupStatusActive,
		"invalid": BareMetalMachineKeySetUserSetupStatusInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetUserSetupStatus(input)
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

type BfdEnabled string

const (
	BfdEnabledFalse BfdEnabled = "False"
	BfdEnabledTrue  BfdEnabled = "True"
)

func PossibleValuesForBfdEnabled() []string {
	return []string{
		string(BfdEnabledFalse),
		string(BfdEnabledTrue),
	}
}

func (s *BfdEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBfdEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBfdEnabled(input string) (*BfdEnabled, error) {
	vals := map[string]BfdEnabled{
		"false": BfdEnabledFalse,
		"true":  BfdEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BfdEnabled(input)
	return &out, nil
}

type BgpMultiHop string

const (
	BgpMultiHopFalse BgpMultiHop = "False"
	BgpMultiHopTrue  BgpMultiHop = "True"
)

func PossibleValuesForBgpMultiHop() []string {
	return []string{
		string(BgpMultiHopFalse),
		string(BgpMultiHopTrue),
	}
}

func (s *BgpMultiHop) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBgpMultiHop(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBgpMultiHop(input string) (*BgpMultiHop, error) {
	vals := map[string]BgpMultiHop{
		"false": BgpMultiHopFalse,
		"true":  BgpMultiHopTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BgpMultiHop(input)
	return &out, nil
}

type BmcKeySetDetailedStatus string

const (
	BmcKeySetDetailedStatusAllActive   BmcKeySetDetailedStatus = "AllActive"
	BmcKeySetDetailedStatusAllInvalid  BmcKeySetDetailedStatus = "AllInvalid"
	BmcKeySetDetailedStatusSomeInvalid BmcKeySetDetailedStatus = "SomeInvalid"
	BmcKeySetDetailedStatusValidating  BmcKeySetDetailedStatus = "Validating"
)

func PossibleValuesForBmcKeySetDetailedStatus() []string {
	return []string{
		string(BmcKeySetDetailedStatusAllActive),
		string(BmcKeySetDetailedStatusAllInvalid),
		string(BmcKeySetDetailedStatusSomeInvalid),
		string(BmcKeySetDetailedStatusValidating),
	}
}

func (s *BmcKeySetDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetDetailedStatus(input string) (*BmcKeySetDetailedStatus, error) {
	vals := map[string]BmcKeySetDetailedStatus{
		"allactive":   BmcKeySetDetailedStatusAllActive,
		"allinvalid":  BmcKeySetDetailedStatusAllInvalid,
		"someinvalid": BmcKeySetDetailedStatusSomeInvalid,
		"validating":  BmcKeySetDetailedStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetDetailedStatus(input)
	return &out, nil
}

type BmcKeySetPrivilegeLevel string

const (
	BmcKeySetPrivilegeLevelAdministrator BmcKeySetPrivilegeLevel = "Administrator"
	BmcKeySetPrivilegeLevelReadOnly      BmcKeySetPrivilegeLevel = "ReadOnly"
)

func PossibleValuesForBmcKeySetPrivilegeLevel() []string {
	return []string{
		string(BmcKeySetPrivilegeLevelAdministrator),
		string(BmcKeySetPrivilegeLevelReadOnly),
	}
}

func (s *BmcKeySetPrivilegeLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetPrivilegeLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetPrivilegeLevel(input string) (*BmcKeySetPrivilegeLevel, error) {
	vals := map[string]BmcKeySetPrivilegeLevel{
		"administrator": BmcKeySetPrivilegeLevelAdministrator,
		"readonly":      BmcKeySetPrivilegeLevelReadOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetPrivilegeLevel(input)
	return &out, nil
}

type BmcKeySetProvisioningState string

const (
	BmcKeySetProvisioningStateAccepted     BmcKeySetProvisioningState = "Accepted"
	BmcKeySetProvisioningStateCanceled     BmcKeySetProvisioningState = "Canceled"
	BmcKeySetProvisioningStateFailed       BmcKeySetProvisioningState = "Failed"
	BmcKeySetProvisioningStateProvisioning BmcKeySetProvisioningState = "Provisioning"
	BmcKeySetProvisioningStateSucceeded    BmcKeySetProvisioningState = "Succeeded"
)

func PossibleValuesForBmcKeySetProvisioningState() []string {
	return []string{
		string(BmcKeySetProvisioningStateAccepted),
		string(BmcKeySetProvisioningStateCanceled),
		string(BmcKeySetProvisioningStateFailed),
		string(BmcKeySetProvisioningStateProvisioning),
		string(BmcKeySetProvisioningStateSucceeded),
	}
}

func (s *BmcKeySetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetProvisioningState(input string) (*BmcKeySetProvisioningState, error) {
	vals := map[string]BmcKeySetProvisioningState{
		"accepted":     BmcKeySetProvisioningStateAccepted,
		"canceled":     BmcKeySetProvisioningStateCanceled,
		"failed":       BmcKeySetProvisioningStateFailed,
		"provisioning": BmcKeySetProvisioningStateProvisioning,
		"succeeded":    BmcKeySetProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetProvisioningState(input)
	return &out, nil
}

type BootstrapProtocol string

const (
	BootstrapProtocolPXE BootstrapProtocol = "PXE"
)

func PossibleValuesForBootstrapProtocol() []string {
	return []string{
		string(BootstrapProtocolPXE),
	}
}

func (s *BootstrapProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBootstrapProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBootstrapProtocol(input string) (*BootstrapProtocol, error) {
	vals := map[string]BootstrapProtocol{
		"pxe": BootstrapProtocolPXE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BootstrapProtocol(input)
	return &out, nil
}

type CloudServicesNetworkDetailedStatus string

const (
	CloudServicesNetworkDetailedStatusAvailable    CloudServicesNetworkDetailedStatus = "Available"
	CloudServicesNetworkDetailedStatusError        CloudServicesNetworkDetailedStatus = "Error"
	CloudServicesNetworkDetailedStatusProvisioning CloudServicesNetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForCloudServicesNetworkDetailedStatus() []string {
	return []string{
		string(CloudServicesNetworkDetailedStatusAvailable),
		string(CloudServicesNetworkDetailedStatusError),
		string(CloudServicesNetworkDetailedStatusProvisioning),
	}
}

func (s *CloudServicesNetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkDetailedStatus(input string) (*CloudServicesNetworkDetailedStatus, error) {
	vals := map[string]CloudServicesNetworkDetailedStatus{
		"available":    CloudServicesNetworkDetailedStatusAvailable,
		"error":        CloudServicesNetworkDetailedStatusError,
		"provisioning": CloudServicesNetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkDetailedStatus(input)
	return &out, nil
}

type CloudServicesNetworkEnableDefaultEgressEndpoints string

const (
	CloudServicesNetworkEnableDefaultEgressEndpointsFalse CloudServicesNetworkEnableDefaultEgressEndpoints = "False"
	CloudServicesNetworkEnableDefaultEgressEndpointsTrue  CloudServicesNetworkEnableDefaultEgressEndpoints = "True"
)

func PossibleValuesForCloudServicesNetworkEnableDefaultEgressEndpoints() []string {
	return []string{
		string(CloudServicesNetworkEnableDefaultEgressEndpointsFalse),
		string(CloudServicesNetworkEnableDefaultEgressEndpointsTrue),
	}
}

func (s *CloudServicesNetworkEnableDefaultEgressEndpoints) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkEnableDefaultEgressEndpoints(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkEnableDefaultEgressEndpoints(input string) (*CloudServicesNetworkEnableDefaultEgressEndpoints, error) {
	vals := map[string]CloudServicesNetworkEnableDefaultEgressEndpoints{
		"false": CloudServicesNetworkEnableDefaultEgressEndpointsFalse,
		"true":  CloudServicesNetworkEnableDefaultEgressEndpointsTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkEnableDefaultEgressEndpoints(input)
	return &out, nil
}

type CloudServicesNetworkProvisioningState string

const (
	CloudServicesNetworkProvisioningStateAccepted     CloudServicesNetworkProvisioningState = "Accepted"
	CloudServicesNetworkProvisioningStateCanceled     CloudServicesNetworkProvisioningState = "Canceled"
	CloudServicesNetworkProvisioningStateFailed       CloudServicesNetworkProvisioningState = "Failed"
	CloudServicesNetworkProvisioningStateProvisioning CloudServicesNetworkProvisioningState = "Provisioning"
	CloudServicesNetworkProvisioningStateSucceeded    CloudServicesNetworkProvisioningState = "Succeeded"
)

func PossibleValuesForCloudServicesNetworkProvisioningState() []string {
	return []string{
		string(CloudServicesNetworkProvisioningStateAccepted),
		string(CloudServicesNetworkProvisioningStateCanceled),
		string(CloudServicesNetworkProvisioningStateFailed),
		string(CloudServicesNetworkProvisioningStateProvisioning),
		string(CloudServicesNetworkProvisioningStateSucceeded),
	}
}

func (s *CloudServicesNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkProvisioningState(input string) (*CloudServicesNetworkProvisioningState, error) {
	vals := map[string]CloudServicesNetworkProvisioningState{
		"accepted":     CloudServicesNetworkProvisioningStateAccepted,
		"canceled":     CloudServicesNetworkProvisioningStateCanceled,
		"failed":       CloudServicesNetworkProvisioningStateFailed,
		"provisioning": CloudServicesNetworkProvisioningStateProvisioning,
		"succeeded":    CloudServicesNetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkProvisioningState(input)
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

type ClusterManagerDetailedStatus string

const (
	ClusterManagerDetailedStatusAvailable          ClusterManagerDetailedStatus = "Available"
	ClusterManagerDetailedStatusError              ClusterManagerDetailedStatus = "Error"
	ClusterManagerDetailedStatusProvisioning       ClusterManagerDetailedStatus = "Provisioning"
	ClusterManagerDetailedStatusProvisioningFailed ClusterManagerDetailedStatus = "ProvisioningFailed"
	ClusterManagerDetailedStatusUpdateFailed       ClusterManagerDetailedStatus = "UpdateFailed"
	ClusterManagerDetailedStatusUpdating           ClusterManagerDetailedStatus = "Updating"
)

func PossibleValuesForClusterManagerDetailedStatus() []string {
	return []string{
		string(ClusterManagerDetailedStatusAvailable),
		string(ClusterManagerDetailedStatusError),
		string(ClusterManagerDetailedStatusProvisioning),
		string(ClusterManagerDetailedStatusProvisioningFailed),
		string(ClusterManagerDetailedStatusUpdateFailed),
		string(ClusterManagerDetailedStatusUpdating),
	}
}

func (s *ClusterManagerDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterManagerDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterManagerDetailedStatus(input string) (*ClusterManagerDetailedStatus, error) {
	vals := map[string]ClusterManagerDetailedStatus{
		"available":          ClusterManagerDetailedStatusAvailable,
		"error":              ClusterManagerDetailedStatusError,
		"provisioning":       ClusterManagerDetailedStatusProvisioning,
		"provisioningfailed": ClusterManagerDetailedStatusProvisioningFailed,
		"updatefailed":       ClusterManagerDetailedStatusUpdateFailed,
		"updating":           ClusterManagerDetailedStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerDetailedStatus(input)
	return &out, nil
}

type ClusterManagerProvisioningState string

const (
	ClusterManagerProvisioningStateAccepted     ClusterManagerProvisioningState = "Accepted"
	ClusterManagerProvisioningStateCanceled     ClusterManagerProvisioningState = "Canceled"
	ClusterManagerProvisioningStateFailed       ClusterManagerProvisioningState = "Failed"
	ClusterManagerProvisioningStateProvisioning ClusterManagerProvisioningState = "Provisioning"
	ClusterManagerProvisioningStateSucceeded    ClusterManagerProvisioningState = "Succeeded"
	ClusterManagerProvisioningStateUpdating     ClusterManagerProvisioningState = "Updating"
)

func PossibleValuesForClusterManagerProvisioningState() []string {
	return []string{
		string(ClusterManagerProvisioningStateAccepted),
		string(ClusterManagerProvisioningStateCanceled),
		string(ClusterManagerProvisioningStateFailed),
		string(ClusterManagerProvisioningStateProvisioning),
		string(ClusterManagerProvisioningStateSucceeded),
		string(ClusterManagerProvisioningStateUpdating),
	}
}

func (s *ClusterManagerProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterManagerProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterManagerProvisioningState(input string) (*ClusterManagerProvisioningState, error) {
	vals := map[string]ClusterManagerProvisioningState{
		"accepted":     ClusterManagerProvisioningStateAccepted,
		"canceled":     ClusterManagerProvisioningStateCanceled,
		"failed":       ClusterManagerProvisioningStateFailed,
		"provisioning": ClusterManagerProvisioningStateProvisioning,
		"succeeded":    ClusterManagerProvisioningStateSucceeded,
		"updating":     ClusterManagerProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerProvisioningState(input)
	return &out, nil
}

type ClusterMetricsConfigurationDetailedStatus string

const (
	ClusterMetricsConfigurationDetailedStatusApplied    ClusterMetricsConfigurationDetailedStatus = "Applied"
	ClusterMetricsConfigurationDetailedStatusError      ClusterMetricsConfigurationDetailedStatus = "Error"
	ClusterMetricsConfigurationDetailedStatusProcessing ClusterMetricsConfigurationDetailedStatus = "Processing"
)

func PossibleValuesForClusterMetricsConfigurationDetailedStatus() []string {
	return []string{
		string(ClusterMetricsConfigurationDetailedStatusApplied),
		string(ClusterMetricsConfigurationDetailedStatusError),
		string(ClusterMetricsConfigurationDetailedStatusProcessing),
	}
}

func (s *ClusterMetricsConfigurationDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterMetricsConfigurationDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterMetricsConfigurationDetailedStatus(input string) (*ClusterMetricsConfigurationDetailedStatus, error) {
	vals := map[string]ClusterMetricsConfigurationDetailedStatus{
		"applied":    ClusterMetricsConfigurationDetailedStatusApplied,
		"error":      ClusterMetricsConfigurationDetailedStatusError,
		"processing": ClusterMetricsConfigurationDetailedStatusProcessing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterMetricsConfigurationDetailedStatus(input)
	return &out, nil
}

type ClusterMetricsConfigurationProvisioningState string

const (
	ClusterMetricsConfigurationProvisioningStateAccepted     ClusterMetricsConfigurationProvisioningState = "Accepted"
	ClusterMetricsConfigurationProvisioningStateCanceled     ClusterMetricsConfigurationProvisioningState = "Canceled"
	ClusterMetricsConfigurationProvisioningStateFailed       ClusterMetricsConfigurationProvisioningState = "Failed"
	ClusterMetricsConfigurationProvisioningStateProvisioning ClusterMetricsConfigurationProvisioningState = "Provisioning"
	ClusterMetricsConfigurationProvisioningStateSucceeded    ClusterMetricsConfigurationProvisioningState = "Succeeded"
)

func PossibleValuesForClusterMetricsConfigurationProvisioningState() []string {
	return []string{
		string(ClusterMetricsConfigurationProvisioningStateAccepted),
		string(ClusterMetricsConfigurationProvisioningStateCanceled),
		string(ClusterMetricsConfigurationProvisioningStateFailed),
		string(ClusterMetricsConfigurationProvisioningStateProvisioning),
		string(ClusterMetricsConfigurationProvisioningStateSucceeded),
	}
}

func (s *ClusterMetricsConfigurationProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterMetricsConfigurationProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterMetricsConfigurationProvisioningState(input string) (*ClusterMetricsConfigurationProvisioningState, error) {
	vals := map[string]ClusterMetricsConfigurationProvisioningState{
		"accepted":     ClusterMetricsConfigurationProvisioningStateAccepted,
		"canceled":     ClusterMetricsConfigurationProvisioningStateCanceled,
		"failed":       ClusterMetricsConfigurationProvisioningStateFailed,
		"provisioning": ClusterMetricsConfigurationProvisioningStateProvisioning,
		"succeeded":    ClusterMetricsConfigurationProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterMetricsConfigurationProvisioningState(input)
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

type ConsoleDetailedStatus string

const (
	ConsoleDetailedStatusError ConsoleDetailedStatus = "Error"
	ConsoleDetailedStatusReady ConsoleDetailedStatus = "Ready"
)

func PossibleValuesForConsoleDetailedStatus() []string {
	return []string{
		string(ConsoleDetailedStatusError),
		string(ConsoleDetailedStatusReady),
	}
}

func (s *ConsoleDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleDetailedStatus(input string) (*ConsoleDetailedStatus, error) {
	vals := map[string]ConsoleDetailedStatus{
		"error": ConsoleDetailedStatusError,
		"ready": ConsoleDetailedStatusReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleDetailedStatus(input)
	return &out, nil
}

type ConsoleEnabled string

const (
	ConsoleEnabledFalse ConsoleEnabled = "False"
	ConsoleEnabledTrue  ConsoleEnabled = "True"
)

func PossibleValuesForConsoleEnabled() []string {
	return []string{
		string(ConsoleEnabledFalse),
		string(ConsoleEnabledTrue),
	}
}

func (s *ConsoleEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleEnabled(input string) (*ConsoleEnabled, error) {
	vals := map[string]ConsoleEnabled{
		"false": ConsoleEnabledFalse,
		"true":  ConsoleEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleEnabled(input)
	return &out, nil
}

type ConsoleProvisioningState string

const (
	ConsoleProvisioningStateAccepted     ConsoleProvisioningState = "Accepted"
	ConsoleProvisioningStateCanceled     ConsoleProvisioningState = "Canceled"
	ConsoleProvisioningStateFailed       ConsoleProvisioningState = "Failed"
	ConsoleProvisioningStateProvisioning ConsoleProvisioningState = "Provisioning"
	ConsoleProvisioningStateSucceeded    ConsoleProvisioningState = "Succeeded"
)

func PossibleValuesForConsoleProvisioningState() []string {
	return []string{
		string(ConsoleProvisioningStateAccepted),
		string(ConsoleProvisioningStateCanceled),
		string(ConsoleProvisioningStateFailed),
		string(ConsoleProvisioningStateProvisioning),
		string(ConsoleProvisioningStateSucceeded),
	}
}

func (s *ConsoleProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleProvisioningState(input string) (*ConsoleProvisioningState, error) {
	vals := map[string]ConsoleProvisioningState{
		"accepted":     ConsoleProvisioningStateAccepted,
		"canceled":     ConsoleProvisioningStateCanceled,
		"failed":       ConsoleProvisioningStateFailed,
		"provisioning": ConsoleProvisioningStateProvisioning,
		"succeeded":    ConsoleProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleProvisioningState(input)
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

type DefaultGateway string

const (
	DefaultGatewayFalse DefaultGateway = "False"
	DefaultGatewayTrue  DefaultGateway = "True"
)

func PossibleValuesForDefaultGateway() []string {
	return []string{
		string(DefaultGatewayFalse),
		string(DefaultGatewayTrue),
	}
}

func (s *DefaultGateway) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultGateway(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultGateway(input string) (*DefaultGateway, error) {
	vals := map[string]DefaultGateway{
		"false": DefaultGatewayFalse,
		"true":  DefaultGatewayTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultGateway(input)
	return &out, nil
}

type DeviceConnectionType string

const (
	DeviceConnectionTypePCI DeviceConnectionType = "PCI"
)

func PossibleValuesForDeviceConnectionType() []string {
	return []string{
		string(DeviceConnectionTypePCI),
	}
}

func (s *DeviceConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConnectionType(input string) (*DeviceConnectionType, error) {
	vals := map[string]DeviceConnectionType{
		"pci": DeviceConnectionTypePCI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConnectionType(input)
	return &out, nil
}

type DiskType string

const (
	DiskTypeHDD DiskType = "HDD"
	DiskTypeSSD DiskType = "SSD"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypeHDD),
		string(DiskTypeSSD),
	}
}

func (s *DiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiskType(input string) (*DiskType, error) {
	vals := map[string]DiskType{
		"hdd": DiskTypeHDD,
		"ssd": DiskTypeSSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskType(input)
	return &out, nil
}

type FabricPeeringEnabled string

const (
	FabricPeeringEnabledFalse FabricPeeringEnabled = "False"
	FabricPeeringEnabledTrue  FabricPeeringEnabled = "True"
)

func PossibleValuesForFabricPeeringEnabled() []string {
	return []string{
		string(FabricPeeringEnabledFalse),
		string(FabricPeeringEnabledTrue),
	}
}

func (s *FabricPeeringEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFabricPeeringEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFabricPeeringEnabled(input string) (*FabricPeeringEnabled, error) {
	vals := map[string]FabricPeeringEnabled{
		"false": FabricPeeringEnabledFalse,
		"true":  FabricPeeringEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FabricPeeringEnabled(input)
	return &out, nil
}

type FeatureDetailedStatus string

const (
	FeatureDetailedStatusFailed  FeatureDetailedStatus = "Failed"
	FeatureDetailedStatusRunning FeatureDetailedStatus = "Running"
	FeatureDetailedStatusUnknown FeatureDetailedStatus = "Unknown"
)

func PossibleValuesForFeatureDetailedStatus() []string {
	return []string{
		string(FeatureDetailedStatusFailed),
		string(FeatureDetailedStatusRunning),
		string(FeatureDetailedStatusUnknown),
	}
}

func (s *FeatureDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureDetailedStatus(input string) (*FeatureDetailedStatus, error) {
	vals := map[string]FeatureDetailedStatus{
		"failed":  FeatureDetailedStatusFailed,
		"running": FeatureDetailedStatusRunning,
		"unknown": FeatureDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureDetailedStatus(input)
	return &out, nil
}

type HugepagesSize string

const (
	HugepagesSizeOneG HugepagesSize = "1G"
	HugepagesSizeTwoM HugepagesSize = "2M"
)

func PossibleValuesForHugepagesSize() []string {
	return []string{
		string(HugepagesSizeOneG),
		string(HugepagesSizeTwoM),
	}
}

func (s *HugepagesSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHugepagesSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHugepagesSize(input string) (*HugepagesSize, error) {
	vals := map[string]HugepagesSize{
		"1g": HugepagesSizeOneG,
		"2m": HugepagesSizeTwoM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HugepagesSize(input)
	return &out, nil
}

type HybridAksIPamEnabled string

const (
	HybridAksIPamEnabledFalse HybridAksIPamEnabled = "False"
	HybridAksIPamEnabledTrue  HybridAksIPamEnabled = "True"
)

func PossibleValuesForHybridAksIPamEnabled() []string {
	return []string{
		string(HybridAksIPamEnabledFalse),
		string(HybridAksIPamEnabledTrue),
	}
}

func (s *HybridAksIPamEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHybridAksIPamEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHybridAksIPamEnabled(input string) (*HybridAksIPamEnabled, error) {
	vals := map[string]HybridAksIPamEnabled{
		"false": HybridAksIPamEnabledFalse,
		"true":  HybridAksIPamEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HybridAksIPamEnabled(input)
	return &out, nil
}

type HybridAksPluginType string

const (
	HybridAksPluginTypeDPDK     HybridAksPluginType = "DPDK"
	HybridAksPluginTypeOSDevice HybridAksPluginType = "OSDevice"
	HybridAksPluginTypeSRIOV    HybridAksPluginType = "SRIOV"
)

func PossibleValuesForHybridAksPluginType() []string {
	return []string{
		string(HybridAksPluginTypeDPDK),
		string(HybridAksPluginTypeOSDevice),
		string(HybridAksPluginTypeSRIOV),
	}
}

func (s *HybridAksPluginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHybridAksPluginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHybridAksPluginType(input string) (*HybridAksPluginType, error) {
	vals := map[string]HybridAksPluginType{
		"dpdk":     HybridAksPluginTypeDPDK,
		"osdevice": HybridAksPluginTypeOSDevice,
		"sriov":    HybridAksPluginTypeSRIOV,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HybridAksPluginType(input)
	return &out, nil
}

type IPAllocationType string

const (
	IPAllocationTypeDualStack IPAllocationType = "DualStack"
	IPAllocationTypeIPVFour   IPAllocationType = "IPV4"
	IPAllocationTypeIPVSix    IPAllocationType = "IPV6"
)

func PossibleValuesForIPAllocationType() []string {
	return []string{
		string(IPAllocationTypeDualStack),
		string(IPAllocationTypeIPVFour),
		string(IPAllocationTypeIPVSix),
	}
}

func (s *IPAllocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPAllocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPAllocationType(input string) (*IPAllocationType, error) {
	vals := map[string]IPAllocationType{
		"dualstack": IPAllocationTypeDualStack,
		"ipv4":      IPAllocationTypeIPVFour,
		"ipv6":      IPAllocationTypeIPVSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAllocationType(input)
	return &out, nil
}

type KubernetesClusterDetailedStatus string

const (
	KubernetesClusterDetailedStatusAvailable    KubernetesClusterDetailedStatus = "Available"
	KubernetesClusterDetailedStatusError        KubernetesClusterDetailedStatus = "Error"
	KubernetesClusterDetailedStatusProvisioning KubernetesClusterDetailedStatus = "Provisioning"
)

func PossibleValuesForKubernetesClusterDetailedStatus() []string {
	return []string{
		string(KubernetesClusterDetailedStatusAvailable),
		string(KubernetesClusterDetailedStatusError),
		string(KubernetesClusterDetailedStatusProvisioning),
	}
}

func (s *KubernetesClusterDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterDetailedStatus(input string) (*KubernetesClusterDetailedStatus, error) {
	vals := map[string]KubernetesClusterDetailedStatus{
		"available":    KubernetesClusterDetailedStatusAvailable,
		"error":        KubernetesClusterDetailedStatusError,
		"provisioning": KubernetesClusterDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterFeatureAvailabilityLifecycle string

const (
	KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable KubernetesClusterFeatureAvailabilityLifecycle = "GenerallyAvailable"
	KubernetesClusterFeatureAvailabilityLifecyclePreview            KubernetesClusterFeatureAvailabilityLifecycle = "Preview"
)

func PossibleValuesForKubernetesClusterFeatureAvailabilityLifecycle() []string {
	return []string{
		string(KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
		string(KubernetesClusterFeatureAvailabilityLifecyclePreview),
	}
}

func (s *KubernetesClusterFeatureAvailabilityLifecycle) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureAvailabilityLifecycle(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureAvailabilityLifecycle(input string) (*KubernetesClusterFeatureAvailabilityLifecycle, error) {
	vals := map[string]KubernetesClusterFeatureAvailabilityLifecycle{
		"generallyavailable": KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable,
		"preview":            KubernetesClusterFeatureAvailabilityLifecyclePreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureAvailabilityLifecycle(input)
	return &out, nil
}

type KubernetesClusterFeatureDetailedStatus string

const (
	KubernetesClusterFeatureDetailedStatusError        KubernetesClusterFeatureDetailedStatus = "Error"
	KubernetesClusterFeatureDetailedStatusInstalled    KubernetesClusterFeatureDetailedStatus = "Installed"
	KubernetesClusterFeatureDetailedStatusProvisioning KubernetesClusterFeatureDetailedStatus = "Provisioning"
)

func PossibleValuesForKubernetesClusterFeatureDetailedStatus() []string {
	return []string{
		string(KubernetesClusterFeatureDetailedStatusError),
		string(KubernetesClusterFeatureDetailedStatusInstalled),
		string(KubernetesClusterFeatureDetailedStatusProvisioning),
	}
}

func (s *KubernetesClusterFeatureDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureDetailedStatus(input string) (*KubernetesClusterFeatureDetailedStatus, error) {
	vals := map[string]KubernetesClusterFeatureDetailedStatus{
		"error":        KubernetesClusterFeatureDetailedStatusError,
		"installed":    KubernetesClusterFeatureDetailedStatusInstalled,
		"provisioning": KubernetesClusterFeatureDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterFeatureProvisioningState string

const (
	KubernetesClusterFeatureProvisioningStateAccepted  KubernetesClusterFeatureProvisioningState = "Accepted"
	KubernetesClusterFeatureProvisioningStateCanceled  KubernetesClusterFeatureProvisioningState = "Canceled"
	KubernetesClusterFeatureProvisioningStateDeleting  KubernetesClusterFeatureProvisioningState = "Deleting"
	KubernetesClusterFeatureProvisioningStateFailed    KubernetesClusterFeatureProvisioningState = "Failed"
	KubernetesClusterFeatureProvisioningStateSucceeded KubernetesClusterFeatureProvisioningState = "Succeeded"
	KubernetesClusterFeatureProvisioningStateUpdating  KubernetesClusterFeatureProvisioningState = "Updating"
)

func PossibleValuesForKubernetesClusterFeatureProvisioningState() []string {
	return []string{
		string(KubernetesClusterFeatureProvisioningStateAccepted),
		string(KubernetesClusterFeatureProvisioningStateCanceled),
		string(KubernetesClusterFeatureProvisioningStateDeleting),
		string(KubernetesClusterFeatureProvisioningStateFailed),
		string(KubernetesClusterFeatureProvisioningStateSucceeded),
		string(KubernetesClusterFeatureProvisioningStateUpdating),
	}
}

func (s *KubernetesClusterFeatureProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureProvisioningState(input string) (*KubernetesClusterFeatureProvisioningState, error) {
	vals := map[string]KubernetesClusterFeatureProvisioningState{
		"accepted":  KubernetesClusterFeatureProvisioningStateAccepted,
		"canceled":  KubernetesClusterFeatureProvisioningStateCanceled,
		"deleting":  KubernetesClusterFeatureProvisioningStateDeleting,
		"failed":    KubernetesClusterFeatureProvisioningStateFailed,
		"succeeded": KubernetesClusterFeatureProvisioningStateSucceeded,
		"updating":  KubernetesClusterFeatureProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureProvisioningState(input)
	return &out, nil
}

type KubernetesClusterFeatureRequired string

const (
	KubernetesClusterFeatureRequiredFalse KubernetesClusterFeatureRequired = "False"
	KubernetesClusterFeatureRequiredTrue  KubernetesClusterFeatureRequired = "True"
)

func PossibleValuesForKubernetesClusterFeatureRequired() []string {
	return []string{
		string(KubernetesClusterFeatureRequiredFalse),
		string(KubernetesClusterFeatureRequiredTrue),
	}
}

func (s *KubernetesClusterFeatureRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterFeatureRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterFeatureRequired(input string) (*KubernetesClusterFeatureRequired, error) {
	vals := map[string]KubernetesClusterFeatureRequired{
		"false": KubernetesClusterFeatureRequiredFalse,
		"true":  KubernetesClusterFeatureRequiredTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterFeatureRequired(input)
	return &out, nil
}

type KubernetesClusterNodeDetailedStatus string

const (
	KubernetesClusterNodeDetailedStatusAvailable    KubernetesClusterNodeDetailedStatus = "Available"
	KubernetesClusterNodeDetailedStatusError        KubernetesClusterNodeDetailedStatus = "Error"
	KubernetesClusterNodeDetailedStatusProvisioning KubernetesClusterNodeDetailedStatus = "Provisioning"
	KubernetesClusterNodeDetailedStatusRunning      KubernetesClusterNodeDetailedStatus = "Running"
	KubernetesClusterNodeDetailedStatusScheduling   KubernetesClusterNodeDetailedStatus = "Scheduling"
	KubernetesClusterNodeDetailedStatusStopped      KubernetesClusterNodeDetailedStatus = "Stopped"
	KubernetesClusterNodeDetailedStatusTerminating  KubernetesClusterNodeDetailedStatus = "Terminating"
	KubernetesClusterNodeDetailedStatusUnknown      KubernetesClusterNodeDetailedStatus = "Unknown"
)

func PossibleValuesForKubernetesClusterNodeDetailedStatus() []string {
	return []string{
		string(KubernetesClusterNodeDetailedStatusAvailable),
		string(KubernetesClusterNodeDetailedStatusError),
		string(KubernetesClusterNodeDetailedStatusProvisioning),
		string(KubernetesClusterNodeDetailedStatusRunning),
		string(KubernetesClusterNodeDetailedStatusScheduling),
		string(KubernetesClusterNodeDetailedStatusStopped),
		string(KubernetesClusterNodeDetailedStatusTerminating),
		string(KubernetesClusterNodeDetailedStatusUnknown),
	}
}

func (s *KubernetesClusterNodeDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterNodeDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterNodeDetailedStatus(input string) (*KubernetesClusterNodeDetailedStatus, error) {
	vals := map[string]KubernetesClusterNodeDetailedStatus{
		"available":    KubernetesClusterNodeDetailedStatusAvailable,
		"error":        KubernetesClusterNodeDetailedStatusError,
		"provisioning": KubernetesClusterNodeDetailedStatusProvisioning,
		"running":      KubernetesClusterNodeDetailedStatusRunning,
		"scheduling":   KubernetesClusterNodeDetailedStatusScheduling,
		"stopped":      KubernetesClusterNodeDetailedStatusStopped,
		"terminating":  KubernetesClusterNodeDetailedStatusTerminating,
		"unknown":      KubernetesClusterNodeDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterNodeDetailedStatus(input)
	return &out, nil
}

type KubernetesClusterProvisioningState string

const (
	KubernetesClusterProvisioningStateAccepted   KubernetesClusterProvisioningState = "Accepted"
	KubernetesClusterProvisioningStateCanceled   KubernetesClusterProvisioningState = "Canceled"
	KubernetesClusterProvisioningStateCreated    KubernetesClusterProvisioningState = "Created"
	KubernetesClusterProvisioningStateDeleting   KubernetesClusterProvisioningState = "Deleting"
	KubernetesClusterProvisioningStateFailed     KubernetesClusterProvisioningState = "Failed"
	KubernetesClusterProvisioningStateInProgress KubernetesClusterProvisioningState = "InProgress"
	KubernetesClusterProvisioningStateSucceeded  KubernetesClusterProvisioningState = "Succeeded"
	KubernetesClusterProvisioningStateUpdating   KubernetesClusterProvisioningState = "Updating"
)

func PossibleValuesForKubernetesClusterProvisioningState() []string {
	return []string{
		string(KubernetesClusterProvisioningStateAccepted),
		string(KubernetesClusterProvisioningStateCanceled),
		string(KubernetesClusterProvisioningStateCreated),
		string(KubernetesClusterProvisioningStateDeleting),
		string(KubernetesClusterProvisioningStateFailed),
		string(KubernetesClusterProvisioningStateInProgress),
		string(KubernetesClusterProvisioningStateSucceeded),
		string(KubernetesClusterProvisioningStateUpdating),
	}
}

func (s *KubernetesClusterProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesClusterProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesClusterProvisioningState(input string) (*KubernetesClusterProvisioningState, error) {
	vals := map[string]KubernetesClusterProvisioningState{
		"accepted":   KubernetesClusterProvisioningStateAccepted,
		"canceled":   KubernetesClusterProvisioningStateCanceled,
		"created":    KubernetesClusterProvisioningStateCreated,
		"deleting":   KubernetesClusterProvisioningStateDeleting,
		"failed":     KubernetesClusterProvisioningStateFailed,
		"inprogress": KubernetesClusterProvisioningStateInProgress,
		"succeeded":  KubernetesClusterProvisioningStateSucceeded,
		"updating":   KubernetesClusterProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesClusterProvisioningState(input)
	return &out, nil
}

type KubernetesNodePowerState string

const (
	KubernetesNodePowerStateOff     KubernetesNodePowerState = "Off"
	KubernetesNodePowerStateOn      KubernetesNodePowerState = "On"
	KubernetesNodePowerStateUnknown KubernetesNodePowerState = "Unknown"
)

func PossibleValuesForKubernetesNodePowerState() []string {
	return []string{
		string(KubernetesNodePowerStateOff),
		string(KubernetesNodePowerStateOn),
		string(KubernetesNodePowerStateUnknown),
	}
}

func (s *KubernetesNodePowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesNodePowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesNodePowerState(input string) (*KubernetesNodePowerState, error) {
	vals := map[string]KubernetesNodePowerState{
		"off":     KubernetesNodePowerStateOff,
		"on":      KubernetesNodePowerStateOn,
		"unknown": KubernetesNodePowerStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesNodePowerState(input)
	return &out, nil
}

type KubernetesNodeRole string

const (
	KubernetesNodeRoleControlPlane KubernetesNodeRole = "ControlPlane"
	KubernetesNodeRoleWorker       KubernetesNodeRole = "Worker"
)

func PossibleValuesForKubernetesNodeRole() []string {
	return []string{
		string(KubernetesNodeRoleControlPlane),
		string(KubernetesNodeRoleWorker),
	}
}

func (s *KubernetesNodeRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesNodeRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesNodeRole(input string) (*KubernetesNodeRole, error) {
	vals := map[string]KubernetesNodeRole{
		"controlplane": KubernetesNodeRoleControlPlane,
		"worker":       KubernetesNodeRoleWorker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesNodeRole(input)
	return &out, nil
}

type KubernetesPluginType string

const (
	KubernetesPluginTypeDPDK     KubernetesPluginType = "DPDK"
	KubernetesPluginTypeIPVLAN   KubernetesPluginType = "IPVLAN"
	KubernetesPluginTypeMACVLAN  KubernetesPluginType = "MACVLAN"
	KubernetesPluginTypeOSDevice KubernetesPluginType = "OSDevice"
	KubernetesPluginTypeSRIOV    KubernetesPluginType = "SRIOV"
)

func PossibleValuesForKubernetesPluginType() []string {
	return []string{
		string(KubernetesPluginTypeDPDK),
		string(KubernetesPluginTypeIPVLAN),
		string(KubernetesPluginTypeMACVLAN),
		string(KubernetesPluginTypeOSDevice),
		string(KubernetesPluginTypeSRIOV),
	}
}

func (s *KubernetesPluginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKubernetesPluginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKubernetesPluginType(input string) (*KubernetesPluginType, error) {
	vals := map[string]KubernetesPluginType{
		"dpdk":     KubernetesPluginTypeDPDK,
		"ipvlan":   KubernetesPluginTypeIPVLAN,
		"macvlan":  KubernetesPluginTypeMACVLAN,
		"osdevice": KubernetesPluginTypeOSDevice,
		"sriov":    KubernetesPluginTypeSRIOV,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesPluginType(input)
	return &out, nil
}

type L2NetworkDetailedStatus string

const (
	L2NetworkDetailedStatusAvailable    L2NetworkDetailedStatus = "Available"
	L2NetworkDetailedStatusError        L2NetworkDetailedStatus = "Error"
	L2NetworkDetailedStatusProvisioning L2NetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForL2NetworkDetailedStatus() []string {
	return []string{
		string(L2NetworkDetailedStatusAvailable),
		string(L2NetworkDetailedStatusError),
		string(L2NetworkDetailedStatusProvisioning),
	}
}

func (s *L2NetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL2NetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL2NetworkDetailedStatus(input string) (*L2NetworkDetailedStatus, error) {
	vals := map[string]L2NetworkDetailedStatus{
		"available":    L2NetworkDetailedStatusAvailable,
		"error":        L2NetworkDetailedStatusError,
		"provisioning": L2NetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L2NetworkDetailedStatus(input)
	return &out, nil
}

type L2NetworkProvisioningState string

const (
	L2NetworkProvisioningStateAccepted     L2NetworkProvisioningState = "Accepted"
	L2NetworkProvisioningStateCanceled     L2NetworkProvisioningState = "Canceled"
	L2NetworkProvisioningStateFailed       L2NetworkProvisioningState = "Failed"
	L2NetworkProvisioningStateProvisioning L2NetworkProvisioningState = "Provisioning"
	L2NetworkProvisioningStateSucceeded    L2NetworkProvisioningState = "Succeeded"
)

func PossibleValuesForL2NetworkProvisioningState() []string {
	return []string{
		string(L2NetworkProvisioningStateAccepted),
		string(L2NetworkProvisioningStateCanceled),
		string(L2NetworkProvisioningStateFailed),
		string(L2NetworkProvisioningStateProvisioning),
		string(L2NetworkProvisioningStateSucceeded),
	}
}

func (s *L2NetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL2NetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL2NetworkProvisioningState(input string) (*L2NetworkProvisioningState, error) {
	vals := map[string]L2NetworkProvisioningState{
		"accepted":     L2NetworkProvisioningStateAccepted,
		"canceled":     L2NetworkProvisioningStateCanceled,
		"failed":       L2NetworkProvisioningStateFailed,
		"provisioning": L2NetworkProvisioningStateProvisioning,
		"succeeded":    L2NetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L2NetworkProvisioningState(input)
	return &out, nil
}

type L3NetworkConfigurationIPamEnabled string

const (
	L3NetworkConfigurationIPamEnabledFalse L3NetworkConfigurationIPamEnabled = "False"
	L3NetworkConfigurationIPamEnabledTrue  L3NetworkConfigurationIPamEnabled = "True"
)

func PossibleValuesForL3NetworkConfigurationIPamEnabled() []string {
	return []string{
		string(L3NetworkConfigurationIPamEnabledFalse),
		string(L3NetworkConfigurationIPamEnabledTrue),
	}
}

func (s *L3NetworkConfigurationIPamEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkConfigurationIPamEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkConfigurationIPamEnabled(input string) (*L3NetworkConfigurationIPamEnabled, error) {
	vals := map[string]L3NetworkConfigurationIPamEnabled{
		"false": L3NetworkConfigurationIPamEnabledFalse,
		"true":  L3NetworkConfigurationIPamEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkConfigurationIPamEnabled(input)
	return &out, nil
}

type L3NetworkDetailedStatus string

const (
	L3NetworkDetailedStatusAvailable    L3NetworkDetailedStatus = "Available"
	L3NetworkDetailedStatusError        L3NetworkDetailedStatus = "Error"
	L3NetworkDetailedStatusProvisioning L3NetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForL3NetworkDetailedStatus() []string {
	return []string{
		string(L3NetworkDetailedStatusAvailable),
		string(L3NetworkDetailedStatusError),
		string(L3NetworkDetailedStatusProvisioning),
	}
}

func (s *L3NetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkDetailedStatus(input string) (*L3NetworkDetailedStatus, error) {
	vals := map[string]L3NetworkDetailedStatus{
		"available":    L3NetworkDetailedStatusAvailable,
		"error":        L3NetworkDetailedStatusError,
		"provisioning": L3NetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkDetailedStatus(input)
	return &out, nil
}

type L3NetworkProvisioningState string

const (
	L3NetworkProvisioningStateAccepted     L3NetworkProvisioningState = "Accepted"
	L3NetworkProvisioningStateCanceled     L3NetworkProvisioningState = "Canceled"
	L3NetworkProvisioningStateFailed       L3NetworkProvisioningState = "Failed"
	L3NetworkProvisioningStateProvisioning L3NetworkProvisioningState = "Provisioning"
	L3NetworkProvisioningStateSucceeded    L3NetworkProvisioningState = "Succeeded"
)

func PossibleValuesForL3NetworkProvisioningState() []string {
	return []string{
		string(L3NetworkProvisioningStateAccepted),
		string(L3NetworkProvisioningStateCanceled),
		string(L3NetworkProvisioningStateFailed),
		string(L3NetworkProvisioningStateProvisioning),
		string(L3NetworkProvisioningStateSucceeded),
	}
}

func (s *L3NetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkProvisioningState(input string) (*L3NetworkProvisioningState, error) {
	vals := map[string]L3NetworkProvisioningState{
		"accepted":     L3NetworkProvisioningStateAccepted,
		"canceled":     L3NetworkProvisioningStateCanceled,
		"failed":       L3NetworkProvisioningStateFailed,
		"provisioning": L3NetworkProvisioningStateProvisioning,
		"succeeded":    L3NetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkProvisioningState(input)
	return &out, nil
}

type MachineSkuDiskConnectionType string

const (
	MachineSkuDiskConnectionTypePCIE MachineSkuDiskConnectionType = "PCIE"
	MachineSkuDiskConnectionTypeRAID MachineSkuDiskConnectionType = "RAID"
	MachineSkuDiskConnectionTypeSAS  MachineSkuDiskConnectionType = "SAS"
	MachineSkuDiskConnectionTypeSATA MachineSkuDiskConnectionType = "SATA"
)

func PossibleValuesForMachineSkuDiskConnectionType() []string {
	return []string{
		string(MachineSkuDiskConnectionTypePCIE),
		string(MachineSkuDiskConnectionTypeRAID),
		string(MachineSkuDiskConnectionTypeSAS),
		string(MachineSkuDiskConnectionTypeSATA),
	}
}

func (s *MachineSkuDiskConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineSkuDiskConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineSkuDiskConnectionType(input string) (*MachineSkuDiskConnectionType, error) {
	vals := map[string]MachineSkuDiskConnectionType{
		"pcie": MachineSkuDiskConnectionTypePCIE,
		"raid": MachineSkuDiskConnectionTypeRAID,
		"sas":  MachineSkuDiskConnectionTypeSAS,
		"sata": MachineSkuDiskConnectionTypeSATA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineSkuDiskConnectionType(input)
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

type OsDiskCreateOption string

const (
	OsDiskCreateOptionEphemeral  OsDiskCreateOption = "Ephemeral"
	OsDiskCreateOptionPersistent OsDiskCreateOption = "Persistent"
)

func PossibleValuesForOsDiskCreateOption() []string {
	return []string{
		string(OsDiskCreateOptionEphemeral),
		string(OsDiskCreateOptionPersistent),
	}
}

func (s *OsDiskCreateOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOsDiskCreateOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOsDiskCreateOption(input string) (*OsDiskCreateOption, error) {
	vals := map[string]OsDiskCreateOption{
		"ephemeral":  OsDiskCreateOptionEphemeral,
		"persistent": OsDiskCreateOptionPersistent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsDiskCreateOption(input)
	return &out, nil
}

type OsDiskDeleteOption string

const (
	OsDiskDeleteOptionDelete OsDiskDeleteOption = "Delete"
)

func PossibleValuesForOsDiskDeleteOption() []string {
	return []string{
		string(OsDiskDeleteOptionDelete),
	}
}

func (s *OsDiskDeleteOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOsDiskDeleteOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOsDiskDeleteOption(input string) (*OsDiskDeleteOption, error) {
	vals := map[string]OsDiskDeleteOption{
		"delete": OsDiskDeleteOptionDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsDiskDeleteOption(input)
	return &out, nil
}

type RackDetailedStatus string

const (
	RackDetailedStatusAvailable    RackDetailedStatus = "Available"
	RackDetailedStatusError        RackDetailedStatus = "Error"
	RackDetailedStatusProvisioning RackDetailedStatus = "Provisioning"
)

func PossibleValuesForRackDetailedStatus() []string {
	return []string{
		string(RackDetailedStatusAvailable),
		string(RackDetailedStatusError),
		string(RackDetailedStatusProvisioning),
	}
}

func (s *RackDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackDetailedStatus(input string) (*RackDetailedStatus, error) {
	vals := map[string]RackDetailedStatus{
		"available":    RackDetailedStatusAvailable,
		"error":        RackDetailedStatusError,
		"provisioning": RackDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackDetailedStatus(input)
	return &out, nil
}

type RackProvisioningState string

const (
	RackProvisioningStateAccepted     RackProvisioningState = "Accepted"
	RackProvisioningStateCanceled     RackProvisioningState = "Canceled"
	RackProvisioningStateFailed       RackProvisioningState = "Failed"
	RackProvisioningStateProvisioning RackProvisioningState = "Provisioning"
	RackProvisioningStateSucceeded    RackProvisioningState = "Succeeded"
)

func PossibleValuesForRackProvisioningState() []string {
	return []string{
		string(RackProvisioningStateAccepted),
		string(RackProvisioningStateCanceled),
		string(RackProvisioningStateFailed),
		string(RackProvisioningStateProvisioning),
		string(RackProvisioningStateSucceeded),
	}
}

func (s *RackProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackProvisioningState(input string) (*RackProvisioningState, error) {
	vals := map[string]RackProvisioningState{
		"accepted":     RackProvisioningStateAccepted,
		"canceled":     RackProvisioningStateCanceled,
		"failed":       RackProvisioningStateFailed,
		"provisioning": RackProvisioningStateProvisioning,
		"succeeded":    RackProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackProvisioningState(input)
	return &out, nil
}

type RackSkuProvisioningState string

const (
	RackSkuProvisioningStateCanceled  RackSkuProvisioningState = "Canceled"
	RackSkuProvisioningStateFailed    RackSkuProvisioningState = "Failed"
	RackSkuProvisioningStateSucceeded RackSkuProvisioningState = "Succeeded"
)

func PossibleValuesForRackSkuProvisioningState() []string {
	return []string{
		string(RackSkuProvisioningStateCanceled),
		string(RackSkuProvisioningStateFailed),
		string(RackSkuProvisioningStateSucceeded),
	}
}

func (s *RackSkuProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackSkuProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackSkuProvisioningState(input string) (*RackSkuProvisioningState, error) {
	vals := map[string]RackSkuProvisioningState{
		"canceled":  RackSkuProvisioningStateCanceled,
		"failed":    RackSkuProvisioningStateFailed,
		"succeeded": RackSkuProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackSkuProvisioningState(input)
	return &out, nil
}

type RackSkuType string

const (
	RackSkuTypeAggregator RackSkuType = "Aggregator"
	RackSkuTypeCompute    RackSkuType = "Compute"
	RackSkuTypeSingle     RackSkuType = "Single"
)

func PossibleValuesForRackSkuType() []string {
	return []string{
		string(RackSkuTypeAggregator),
		string(RackSkuTypeCompute),
		string(RackSkuTypeSingle),
	}
}

func (s *RackSkuType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackSkuType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackSkuType(input string) (*RackSkuType, error) {
	vals := map[string]RackSkuType{
		"aggregator": RackSkuTypeAggregator,
		"compute":    RackSkuTypeCompute,
		"single":     RackSkuTypeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackSkuType(input)
	return &out, nil
}

type RemoteVendorManagementFeature string

const (
	RemoteVendorManagementFeatureSupported   RemoteVendorManagementFeature = "Supported"
	RemoteVendorManagementFeatureUnsupported RemoteVendorManagementFeature = "Unsupported"
)

func PossibleValuesForRemoteVendorManagementFeature() []string {
	return []string{
		string(RemoteVendorManagementFeatureSupported),
		string(RemoteVendorManagementFeatureUnsupported),
	}
}

func (s *RemoteVendorManagementFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteVendorManagementFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteVendorManagementFeature(input string) (*RemoteVendorManagementFeature, error) {
	vals := map[string]RemoteVendorManagementFeature{
		"supported":   RemoteVendorManagementFeatureSupported,
		"unsupported": RemoteVendorManagementFeatureUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteVendorManagementFeature(input)
	return &out, nil
}

type RemoteVendorManagementStatus string

const (
	RemoteVendorManagementStatusDisabled    RemoteVendorManagementStatus = "Disabled"
	RemoteVendorManagementStatusEnabled     RemoteVendorManagementStatus = "Enabled"
	RemoteVendorManagementStatusUnsupported RemoteVendorManagementStatus = "Unsupported"
)

func PossibleValuesForRemoteVendorManagementStatus() []string {
	return []string{
		string(RemoteVendorManagementStatusDisabled),
		string(RemoteVendorManagementStatusEnabled),
		string(RemoteVendorManagementStatusUnsupported),
	}
}

func (s *RemoteVendorManagementStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteVendorManagementStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteVendorManagementStatus(input string) (*RemoteVendorManagementStatus, error) {
	vals := map[string]RemoteVendorManagementStatus{
		"disabled":    RemoteVendorManagementStatusDisabled,
		"enabled":     RemoteVendorManagementStatusEnabled,
		"unsupported": RemoteVendorManagementStatusUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteVendorManagementStatus(input)
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

type SkipShutdown string

const (
	SkipShutdownFalse SkipShutdown = "False"
	SkipShutdownTrue  SkipShutdown = "True"
)

func PossibleValuesForSkipShutdown() []string {
	return []string{
		string(SkipShutdownFalse),
		string(SkipShutdownTrue),
	}
}

func (s *SkipShutdown) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkipShutdown(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkipShutdown(input string) (*SkipShutdown, error) {
	vals := map[string]SkipShutdown{
		"false": SkipShutdownFalse,
		"true":  SkipShutdownTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkipShutdown(input)
	return &out, nil
}

type StorageApplianceDetailedStatus string

const (
	StorageApplianceDetailedStatusAvailable    StorageApplianceDetailedStatus = "Available"
	StorageApplianceDetailedStatusDegraded     StorageApplianceDetailedStatus = "Degraded"
	StorageApplianceDetailedStatusError        StorageApplianceDetailedStatus = "Error"
	StorageApplianceDetailedStatusProvisioning StorageApplianceDetailedStatus = "Provisioning"
)

func PossibleValuesForStorageApplianceDetailedStatus() []string {
	return []string{
		string(StorageApplianceDetailedStatusAvailable),
		string(StorageApplianceDetailedStatusDegraded),
		string(StorageApplianceDetailedStatusError),
		string(StorageApplianceDetailedStatusProvisioning),
	}
}

func (s *StorageApplianceDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageApplianceDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageApplianceDetailedStatus(input string) (*StorageApplianceDetailedStatus, error) {
	vals := map[string]StorageApplianceDetailedStatus{
		"available":    StorageApplianceDetailedStatusAvailable,
		"degraded":     StorageApplianceDetailedStatusDegraded,
		"error":        StorageApplianceDetailedStatusError,
		"provisioning": StorageApplianceDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageApplianceDetailedStatus(input)
	return &out, nil
}

type StorageApplianceProvisioningState string

const (
	StorageApplianceProvisioningStateAccepted     StorageApplianceProvisioningState = "Accepted"
	StorageApplianceProvisioningStateCanceled     StorageApplianceProvisioningState = "Canceled"
	StorageApplianceProvisioningStateFailed       StorageApplianceProvisioningState = "Failed"
	StorageApplianceProvisioningStateProvisioning StorageApplianceProvisioningState = "Provisioning"
	StorageApplianceProvisioningStateSucceeded    StorageApplianceProvisioningState = "Succeeded"
)

func PossibleValuesForStorageApplianceProvisioningState() []string {
	return []string{
		string(StorageApplianceProvisioningStateAccepted),
		string(StorageApplianceProvisioningStateCanceled),
		string(StorageApplianceProvisioningStateFailed),
		string(StorageApplianceProvisioningStateProvisioning),
		string(StorageApplianceProvisioningStateSucceeded),
	}
}

func (s *StorageApplianceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageApplianceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageApplianceProvisioningState(input string) (*StorageApplianceProvisioningState, error) {
	vals := map[string]StorageApplianceProvisioningState{
		"accepted":     StorageApplianceProvisioningStateAccepted,
		"canceled":     StorageApplianceProvisioningStateCanceled,
		"failed":       StorageApplianceProvisioningStateFailed,
		"provisioning": StorageApplianceProvisioningStateProvisioning,
		"succeeded":    StorageApplianceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageApplianceProvisioningState(input)
	return &out, nil
}

type TrunkedNetworkDetailedStatus string

const (
	TrunkedNetworkDetailedStatusAvailable    TrunkedNetworkDetailedStatus = "Available"
	TrunkedNetworkDetailedStatusError        TrunkedNetworkDetailedStatus = "Error"
	TrunkedNetworkDetailedStatusProvisioning TrunkedNetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForTrunkedNetworkDetailedStatus() []string {
	return []string{
		string(TrunkedNetworkDetailedStatusAvailable),
		string(TrunkedNetworkDetailedStatusError),
		string(TrunkedNetworkDetailedStatusProvisioning),
	}
}

func (s *TrunkedNetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrunkedNetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrunkedNetworkDetailedStatus(input string) (*TrunkedNetworkDetailedStatus, error) {
	vals := map[string]TrunkedNetworkDetailedStatus{
		"available":    TrunkedNetworkDetailedStatusAvailable,
		"error":        TrunkedNetworkDetailedStatusError,
		"provisioning": TrunkedNetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrunkedNetworkDetailedStatus(input)
	return &out, nil
}

type TrunkedNetworkProvisioningState string

const (
	TrunkedNetworkProvisioningStateAccepted     TrunkedNetworkProvisioningState = "Accepted"
	TrunkedNetworkProvisioningStateCanceled     TrunkedNetworkProvisioningState = "Canceled"
	TrunkedNetworkProvisioningStateFailed       TrunkedNetworkProvisioningState = "Failed"
	TrunkedNetworkProvisioningStateProvisioning TrunkedNetworkProvisioningState = "Provisioning"
	TrunkedNetworkProvisioningStateSucceeded    TrunkedNetworkProvisioningState = "Succeeded"
)

func PossibleValuesForTrunkedNetworkProvisioningState() []string {
	return []string{
		string(TrunkedNetworkProvisioningStateAccepted),
		string(TrunkedNetworkProvisioningStateCanceled),
		string(TrunkedNetworkProvisioningStateFailed),
		string(TrunkedNetworkProvisioningStateProvisioning),
		string(TrunkedNetworkProvisioningStateSucceeded),
	}
}

func (s *TrunkedNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrunkedNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrunkedNetworkProvisioningState(input string) (*TrunkedNetworkProvisioningState, error) {
	vals := map[string]TrunkedNetworkProvisioningState{
		"accepted":     TrunkedNetworkProvisioningStateAccepted,
		"canceled":     TrunkedNetworkProvisioningStateCanceled,
		"failed":       TrunkedNetworkProvisioningStateFailed,
		"provisioning": TrunkedNetworkProvisioningStateProvisioning,
		"succeeded":    TrunkedNetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrunkedNetworkProvisioningState(input)
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

type VirtualMachineBootMethod string

const (
	VirtualMachineBootMethodBIOS VirtualMachineBootMethod = "BIOS"
	VirtualMachineBootMethodUEFI VirtualMachineBootMethod = "UEFI"
)

func PossibleValuesForVirtualMachineBootMethod() []string {
	return []string{
		string(VirtualMachineBootMethodBIOS),
		string(VirtualMachineBootMethodUEFI),
	}
}

func (s *VirtualMachineBootMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineBootMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineBootMethod(input string) (*VirtualMachineBootMethod, error) {
	vals := map[string]VirtualMachineBootMethod{
		"bios": VirtualMachineBootMethodBIOS,
		"uefi": VirtualMachineBootMethodUEFI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineBootMethod(input)
	return &out, nil
}

type VirtualMachineDetailedStatus string

const (
	VirtualMachineDetailedStatusAvailable    VirtualMachineDetailedStatus = "Available"
	VirtualMachineDetailedStatusError        VirtualMachineDetailedStatus = "Error"
	VirtualMachineDetailedStatusProvisioning VirtualMachineDetailedStatus = "Provisioning"
	VirtualMachineDetailedStatusRunning      VirtualMachineDetailedStatus = "Running"
	VirtualMachineDetailedStatusScheduling   VirtualMachineDetailedStatus = "Scheduling"
	VirtualMachineDetailedStatusStopped      VirtualMachineDetailedStatus = "Stopped"
	VirtualMachineDetailedStatusTerminating  VirtualMachineDetailedStatus = "Terminating"
	VirtualMachineDetailedStatusUnknown      VirtualMachineDetailedStatus = "Unknown"
)

func PossibleValuesForVirtualMachineDetailedStatus() []string {
	return []string{
		string(VirtualMachineDetailedStatusAvailable),
		string(VirtualMachineDetailedStatusError),
		string(VirtualMachineDetailedStatusProvisioning),
		string(VirtualMachineDetailedStatusRunning),
		string(VirtualMachineDetailedStatusScheduling),
		string(VirtualMachineDetailedStatusStopped),
		string(VirtualMachineDetailedStatusTerminating),
		string(VirtualMachineDetailedStatusUnknown),
	}
}

func (s *VirtualMachineDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineDetailedStatus(input string) (*VirtualMachineDetailedStatus, error) {
	vals := map[string]VirtualMachineDetailedStatus{
		"available":    VirtualMachineDetailedStatusAvailable,
		"error":        VirtualMachineDetailedStatusError,
		"provisioning": VirtualMachineDetailedStatusProvisioning,
		"running":      VirtualMachineDetailedStatusRunning,
		"scheduling":   VirtualMachineDetailedStatusScheduling,
		"stopped":      VirtualMachineDetailedStatusStopped,
		"terminating":  VirtualMachineDetailedStatusTerminating,
		"unknown":      VirtualMachineDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineDetailedStatus(input)
	return &out, nil
}

type VirtualMachineDeviceModelType string

const (
	VirtualMachineDeviceModelTypeTOne   VirtualMachineDeviceModelType = "T1"
	VirtualMachineDeviceModelTypeTThree VirtualMachineDeviceModelType = "T3"
	VirtualMachineDeviceModelTypeTTwo   VirtualMachineDeviceModelType = "T2"
)

func PossibleValuesForVirtualMachineDeviceModelType() []string {
	return []string{
		string(VirtualMachineDeviceModelTypeTOne),
		string(VirtualMachineDeviceModelTypeTThree),
		string(VirtualMachineDeviceModelTypeTTwo),
	}
}

func (s *VirtualMachineDeviceModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineDeviceModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineDeviceModelType(input string) (*VirtualMachineDeviceModelType, error) {
	vals := map[string]VirtualMachineDeviceModelType{
		"t1": VirtualMachineDeviceModelTypeTOne,
		"t3": VirtualMachineDeviceModelTypeTThree,
		"t2": VirtualMachineDeviceModelTypeTTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineDeviceModelType(input)
	return &out, nil
}

type VirtualMachineIPAllocationMethod string

const (
	VirtualMachineIPAllocationMethodDisabled VirtualMachineIPAllocationMethod = "Disabled"
	VirtualMachineIPAllocationMethodDynamic  VirtualMachineIPAllocationMethod = "Dynamic"
	VirtualMachineIPAllocationMethodStatic   VirtualMachineIPAllocationMethod = "Static"
)

func PossibleValuesForVirtualMachineIPAllocationMethod() []string {
	return []string{
		string(VirtualMachineIPAllocationMethodDisabled),
		string(VirtualMachineIPAllocationMethodDynamic),
		string(VirtualMachineIPAllocationMethodStatic),
	}
}

func (s *VirtualMachineIPAllocationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineIPAllocationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineIPAllocationMethod(input string) (*VirtualMachineIPAllocationMethod, error) {
	vals := map[string]VirtualMachineIPAllocationMethod{
		"disabled": VirtualMachineIPAllocationMethodDisabled,
		"dynamic":  VirtualMachineIPAllocationMethodDynamic,
		"static":   VirtualMachineIPAllocationMethodStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineIPAllocationMethod(input)
	return &out, nil
}

type VirtualMachineIsolateEmulatorThread string

const (
	VirtualMachineIsolateEmulatorThreadFalse VirtualMachineIsolateEmulatorThread = "False"
	VirtualMachineIsolateEmulatorThreadTrue  VirtualMachineIsolateEmulatorThread = "True"
)

func PossibleValuesForVirtualMachineIsolateEmulatorThread() []string {
	return []string{
		string(VirtualMachineIsolateEmulatorThreadFalse),
		string(VirtualMachineIsolateEmulatorThreadTrue),
	}
}

func (s *VirtualMachineIsolateEmulatorThread) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineIsolateEmulatorThread(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineIsolateEmulatorThread(input string) (*VirtualMachineIsolateEmulatorThread, error) {
	vals := map[string]VirtualMachineIsolateEmulatorThread{
		"false": VirtualMachineIsolateEmulatorThreadFalse,
		"true":  VirtualMachineIsolateEmulatorThreadTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineIsolateEmulatorThread(input)
	return &out, nil
}

type VirtualMachinePlacementHintPodAffinityScope string

const (
	VirtualMachinePlacementHintPodAffinityScopeMachine VirtualMachinePlacementHintPodAffinityScope = "Machine"
	VirtualMachinePlacementHintPodAffinityScopeRack    VirtualMachinePlacementHintPodAffinityScope = "Rack"
)

func PossibleValuesForVirtualMachinePlacementHintPodAffinityScope() []string {
	return []string{
		string(VirtualMachinePlacementHintPodAffinityScopeMachine),
		string(VirtualMachinePlacementHintPodAffinityScopeRack),
	}
}

func (s *VirtualMachinePlacementHintPodAffinityScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePlacementHintPodAffinityScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePlacementHintPodAffinityScope(input string) (*VirtualMachinePlacementHintPodAffinityScope, error) {
	vals := map[string]VirtualMachinePlacementHintPodAffinityScope{
		"machine": VirtualMachinePlacementHintPodAffinityScopeMachine,
		"rack":    VirtualMachinePlacementHintPodAffinityScopeRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePlacementHintPodAffinityScope(input)
	return &out, nil
}

type VirtualMachinePlacementHintType string

const (
	VirtualMachinePlacementHintTypeAffinity     VirtualMachinePlacementHintType = "Affinity"
	VirtualMachinePlacementHintTypeAntiAffinity VirtualMachinePlacementHintType = "AntiAffinity"
)

func PossibleValuesForVirtualMachinePlacementHintType() []string {
	return []string{
		string(VirtualMachinePlacementHintTypeAffinity),
		string(VirtualMachinePlacementHintTypeAntiAffinity),
	}
}

func (s *VirtualMachinePlacementHintType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePlacementHintType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePlacementHintType(input string) (*VirtualMachinePlacementHintType, error) {
	vals := map[string]VirtualMachinePlacementHintType{
		"affinity":     VirtualMachinePlacementHintTypeAffinity,
		"antiaffinity": VirtualMachinePlacementHintTypeAntiAffinity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePlacementHintType(input)
	return &out, nil
}

type VirtualMachinePowerState string

const (
	VirtualMachinePowerStateOff     VirtualMachinePowerState = "Off"
	VirtualMachinePowerStateOn      VirtualMachinePowerState = "On"
	VirtualMachinePowerStateUnknown VirtualMachinePowerState = "Unknown"
)

func PossibleValuesForVirtualMachinePowerState() []string {
	return []string{
		string(VirtualMachinePowerStateOff),
		string(VirtualMachinePowerStateOn),
		string(VirtualMachinePowerStateUnknown),
	}
}

func (s *VirtualMachinePowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePowerState(input string) (*VirtualMachinePowerState, error) {
	vals := map[string]VirtualMachinePowerState{
		"off":     VirtualMachinePowerStateOff,
		"on":      VirtualMachinePowerStateOn,
		"unknown": VirtualMachinePowerStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePowerState(input)
	return &out, nil
}

type VirtualMachineProvisioningState string

const (
	VirtualMachineProvisioningStateAccepted     VirtualMachineProvisioningState = "Accepted"
	VirtualMachineProvisioningStateCanceled     VirtualMachineProvisioningState = "Canceled"
	VirtualMachineProvisioningStateFailed       VirtualMachineProvisioningState = "Failed"
	VirtualMachineProvisioningStateProvisioning VirtualMachineProvisioningState = "Provisioning"
	VirtualMachineProvisioningStateSucceeded    VirtualMachineProvisioningState = "Succeeded"
)

func PossibleValuesForVirtualMachineProvisioningState() []string {
	return []string{
		string(VirtualMachineProvisioningStateAccepted),
		string(VirtualMachineProvisioningStateCanceled),
		string(VirtualMachineProvisioningStateFailed),
		string(VirtualMachineProvisioningStateProvisioning),
		string(VirtualMachineProvisioningStateSucceeded),
	}
}

func (s *VirtualMachineProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineProvisioningState(input string) (*VirtualMachineProvisioningState, error) {
	vals := map[string]VirtualMachineProvisioningState{
		"accepted":     VirtualMachineProvisioningStateAccepted,
		"canceled":     VirtualMachineProvisioningStateCanceled,
		"failed":       VirtualMachineProvisioningStateFailed,
		"provisioning": VirtualMachineProvisioningStateProvisioning,
		"succeeded":    VirtualMachineProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineProvisioningState(input)
	return &out, nil
}

type VirtualMachineSchedulingExecution string

const (
	VirtualMachineSchedulingExecutionHard VirtualMachineSchedulingExecution = "Hard"
	VirtualMachineSchedulingExecutionSoft VirtualMachineSchedulingExecution = "Soft"
)

func PossibleValuesForVirtualMachineSchedulingExecution() []string {
	return []string{
		string(VirtualMachineSchedulingExecutionHard),
		string(VirtualMachineSchedulingExecutionSoft),
	}
}

func (s *VirtualMachineSchedulingExecution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineSchedulingExecution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineSchedulingExecution(input string) (*VirtualMachineSchedulingExecution, error) {
	vals := map[string]VirtualMachineSchedulingExecution{
		"hard": VirtualMachineSchedulingExecutionHard,
		"soft": VirtualMachineSchedulingExecutionSoft,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineSchedulingExecution(input)
	return &out, nil
}

type VirtualMachineVirtioInterfaceType string

const (
	VirtualMachineVirtioInterfaceTypeModern       VirtualMachineVirtioInterfaceType = "Modern"
	VirtualMachineVirtioInterfaceTypeTransitional VirtualMachineVirtioInterfaceType = "Transitional"
)

func PossibleValuesForVirtualMachineVirtioInterfaceType() []string {
	return []string{
		string(VirtualMachineVirtioInterfaceTypeModern),
		string(VirtualMachineVirtioInterfaceTypeTransitional),
	}
}

func (s *VirtualMachineVirtioInterfaceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineVirtioInterfaceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineVirtioInterfaceType(input string) (*VirtualMachineVirtioInterfaceType, error) {
	vals := map[string]VirtualMachineVirtioInterfaceType{
		"modern":       VirtualMachineVirtioInterfaceTypeModern,
		"transitional": VirtualMachineVirtioInterfaceTypeTransitional,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineVirtioInterfaceType(input)
	return &out, nil
}

type VolumeDetailedStatus string

const (
	VolumeDetailedStatusActive       VolumeDetailedStatus = "Active"
	VolumeDetailedStatusError        VolumeDetailedStatus = "Error"
	VolumeDetailedStatusProvisioning VolumeDetailedStatus = "Provisioning"
)

func PossibleValuesForVolumeDetailedStatus() []string {
	return []string{
		string(VolumeDetailedStatusActive),
		string(VolumeDetailedStatusError),
		string(VolumeDetailedStatusProvisioning),
	}
}

func (s *VolumeDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeDetailedStatus(input string) (*VolumeDetailedStatus, error) {
	vals := map[string]VolumeDetailedStatus{
		"active":       VolumeDetailedStatusActive,
		"error":        VolumeDetailedStatusError,
		"provisioning": VolumeDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeDetailedStatus(input)
	return &out, nil
}

type VolumeProvisioningState string

const (
	VolumeProvisioningStateAccepted     VolumeProvisioningState = "Accepted"
	VolumeProvisioningStateCanceled     VolumeProvisioningState = "Canceled"
	VolumeProvisioningStateFailed       VolumeProvisioningState = "Failed"
	VolumeProvisioningStateProvisioning VolumeProvisioningState = "Provisioning"
	VolumeProvisioningStateSucceeded    VolumeProvisioningState = "Succeeded"
)

func PossibleValuesForVolumeProvisioningState() []string {
	return []string{
		string(VolumeProvisioningStateAccepted),
		string(VolumeProvisioningStateCanceled),
		string(VolumeProvisioningStateFailed),
		string(VolumeProvisioningStateProvisioning),
		string(VolumeProvisioningStateSucceeded),
	}
}

func (s *VolumeProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeProvisioningState(input string) (*VolumeProvisioningState, error) {
	vals := map[string]VolumeProvisioningState{
		"accepted":     VolumeProvisioningStateAccepted,
		"canceled":     VolumeProvisioningStateCanceled,
		"failed":       VolumeProvisioningStateFailed,
		"provisioning": VolumeProvisioningStateProvisioning,
		"succeeded":    VolumeProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeProvisioningState(input)
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
