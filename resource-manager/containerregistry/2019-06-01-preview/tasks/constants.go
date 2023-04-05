package tasks

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

type BaseImageDependencyType string

const (
	BaseImageDependencyTypeBuildTime BaseImageDependencyType = "BuildTime"
	BaseImageDependencyTypeRunTime   BaseImageDependencyType = "RunTime"
)

func PossibleValuesForBaseImageDependencyType() []string {
	return []string{
		string(BaseImageDependencyTypeBuildTime),
		string(BaseImageDependencyTypeRunTime),
	}
}

func (s *BaseImageDependencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForBaseImageDependencyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = BaseImageDependencyType(decoded)
	return nil
}

type BaseImageTriggerType string

const (
	BaseImageTriggerTypeAll     BaseImageTriggerType = "All"
	BaseImageTriggerTypeRuntime BaseImageTriggerType = "Runtime"
)

func PossibleValuesForBaseImageTriggerType() []string {
	return []string{
		string(BaseImageTriggerTypeAll),
		string(BaseImageTriggerTypeRuntime),
	}
}

func (s *BaseImageTriggerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForBaseImageTriggerType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = BaseImageTriggerType(decoded)
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

type SecretObjectType string

const (
	SecretObjectTypeOpaque      SecretObjectType = "Opaque"
	SecretObjectTypeVaultsecret SecretObjectType = "Vaultsecret"
)

func PossibleValuesForSecretObjectType() []string {
	return []string{
		string(SecretObjectTypeOpaque),
		string(SecretObjectTypeVaultsecret),
	}
}

func (s *SecretObjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSecretObjectType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SecretObjectType(decoded)
	return nil
}

type SourceControlType string

const (
	SourceControlTypeGithub                  SourceControlType = "Github"
	SourceControlTypeVisualStudioTeamService SourceControlType = "VisualStudioTeamService"
)

func PossibleValuesForSourceControlType() []string {
	return []string{
		string(SourceControlTypeGithub),
		string(SourceControlTypeVisualStudioTeamService),
	}
}

func (s *SourceControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSourceControlType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SourceControlType(decoded)
	return nil
}

type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = "Default"
	SourceRegistryLoginModeNone    SourceRegistryLoginMode = "None"
)

func PossibleValuesForSourceRegistryLoginMode() []string {
	return []string{
		string(SourceRegistryLoginModeDefault),
		string(SourceRegistryLoginModeNone),
	}
}

func (s *SourceRegistryLoginMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSourceRegistryLoginMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SourceRegistryLoginMode(decoded)
	return nil
}

type SourceTriggerEvent string

const (
	SourceTriggerEventCommit      SourceTriggerEvent = "commit"
	SourceTriggerEventPullrequest SourceTriggerEvent = "pullrequest"
)

func PossibleValuesForSourceTriggerEvent() []string {
	return []string{
		string(SourceTriggerEventCommit),
		string(SourceTriggerEventPullrequest),
	}
}

func (s *SourceTriggerEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSourceTriggerEvent() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SourceTriggerEvent(decoded)
	return nil
}

type StepType string

const (
	StepTypeDocker      StepType = "Docker"
	StepTypeEncodedTask StepType = "EncodedTask"
	StepTypeFileTask    StepType = "FileTask"
)

func PossibleValuesForStepType() []string {
	return []string{
		string(StepTypeDocker),
		string(StepTypeEncodedTask),
		string(StepTypeFileTask),
	}
}

func (s *StepType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStepType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StepType(decoded)
	return nil
}

type TaskStatus string

const (
	TaskStatusDisabled TaskStatus = "Disabled"
	TaskStatusEnabled  TaskStatus = "Enabled"
)

func PossibleValuesForTaskStatus() []string {
	return []string{
		string(TaskStatusDisabled),
		string(TaskStatusEnabled),
	}
}

func (s *TaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTaskStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TaskStatus(decoded)
	return nil
}

type TokenType string

const (
	TokenTypeOAuth TokenType = "OAuth"
	TokenTypePAT   TokenType = "PAT"
)

func PossibleValuesForTokenType() []string {
	return []string{
		string(TokenTypeOAuth),
		string(TokenTypePAT),
	}
}

func (s *TokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenType(decoded)
	return nil
}

type TriggerStatus string

const (
	TriggerStatusDisabled TriggerStatus = "Disabled"
	TriggerStatusEnabled  TriggerStatus = "Enabled"
)

func PossibleValuesForTriggerStatus() []string {
	return []string{
		string(TriggerStatusDisabled),
		string(TriggerStatusEnabled),
	}
}

func (s *TriggerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTriggerStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TriggerStatus(decoded)
	return nil
}

type UpdateTriggerPayloadType string

const (
	UpdateTriggerPayloadTypeDefault UpdateTriggerPayloadType = "Default"
	UpdateTriggerPayloadTypeToken   UpdateTriggerPayloadType = "Token"
)

func PossibleValuesForUpdateTriggerPayloadType() []string {
	return []string{
		string(UpdateTriggerPayloadTypeDefault),
		string(UpdateTriggerPayloadTypeToken),
	}
}

func (s *UpdateTriggerPayloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForUpdateTriggerPayloadType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = UpdateTriggerPayloadType(decoded)
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
