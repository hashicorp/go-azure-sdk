package runs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Architecture string

const (
	ArchitectureAmdSixFour    Architecture = "amd64"
	ArchitectureArm           Architecture = "arm"
	ArchitectureArmSixFour    Architecture = "arm64"
	ArchitectureThreeEightSix Architecture = "386"
	ArchitectureXEightSix     Architecture = "x86"
)

func PossibleValuesForArchitecture() []string {
	return []string{
		string(ArchitectureAmdSixFour),
		string(ArchitectureArm),
		string(ArchitectureArmSixFour),
		string(ArchitectureThreeEightSix),
		string(ArchitectureXEightSix),
	}
}

func (s *Architecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForArchitecture() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Architecture(decoded)
	return nil
}

type OS string

const (
	OSLinux   OS = "Linux"
	OSWindows OS = "Windows"
)

func PossibleValuesForOS() []string {
	return []string{
		string(OSLinux),
		string(OSWindows),
	}
}

func (s *OS) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForOS() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = OS(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type RunStatus string

const (
	RunStatusCanceled  RunStatus = "Canceled"
	RunStatusError     RunStatus = "Error"
	RunStatusFailed    RunStatus = "Failed"
	RunStatusQueued    RunStatus = "Queued"
	RunStatusRunning   RunStatus = "Running"
	RunStatusStarted   RunStatus = "Started"
	RunStatusSucceeded RunStatus = "Succeeded"
	RunStatusTimeout   RunStatus = "Timeout"
)

func PossibleValuesForRunStatus() []string {
	return []string{
		string(RunStatusCanceled),
		string(RunStatusError),
		string(RunStatusFailed),
		string(RunStatusQueued),
		string(RunStatusRunning),
		string(RunStatusStarted),
		string(RunStatusSucceeded),
		string(RunStatusTimeout),
	}
}

func (s *RunStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRunStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RunStatus(decoded)
	return nil
}

type RunType string

const (
	RunTypeAutoBuild  RunType = "AutoBuild"
	RunTypeAutoRun    RunType = "AutoRun"
	RunTypeQuickBuild RunType = "QuickBuild"
	RunTypeQuickRun   RunType = "QuickRun"
)

func PossibleValuesForRunType() []string {
	return []string{
		string(RunTypeAutoBuild),
		string(RunTypeAutoRun),
		string(RunTypeQuickBuild),
		string(RunTypeQuickRun),
	}
}

func (s *RunType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRunType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RunType(decoded)
	return nil
}

type Variant string

const (
	VariantVEight Variant = "v8"
	VariantVSeven Variant = "v7"
	VariantVSix   Variant = "v6"
)

func PossibleValuesForVariant() []string {
	return []string{
		string(VariantVEight),
		string(VariantVSeven),
		string(VariantVSix),
	}
}

func (s *Variant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVariant() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Variant(decoded)
	return nil
}
