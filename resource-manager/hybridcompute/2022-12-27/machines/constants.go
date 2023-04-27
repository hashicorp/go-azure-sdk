package machines

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentConfigurationMode string

const (
	AgentConfigurationModeFull    AgentConfigurationMode = "full"
	AgentConfigurationModeMonitor AgentConfigurationMode = "monitor"
)

func PossibleValuesForAgentConfigurationMode() []string {
	return []string{
		string(AgentConfigurationModeFull),
		string(AgentConfigurationModeMonitor),
	}
}

func parseAgentConfigurationMode(input string) (*AgentConfigurationMode, error) {
	vals := map[string]AgentConfigurationMode{
		"full":    AgentConfigurationModeFull,
		"monitor": AgentConfigurationModeMonitor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentConfigurationMode(input)
	return &out, nil
}

type AssessmentModeTypes string

const (
	AssessmentModeTypesAutomaticByPlatform AssessmentModeTypes = "AutomaticByPlatform"
	AssessmentModeTypesImageDefault        AssessmentModeTypes = "ImageDefault"
)

func PossibleValuesForAssessmentModeTypes() []string {
	return []string{
		string(AssessmentModeTypesAutomaticByPlatform),
		string(AssessmentModeTypesImageDefault),
	}
}

func parseAssessmentModeTypes(input string) (*AssessmentModeTypes, error) {
	vals := map[string]AssessmentModeTypes{
		"automaticbyplatform": AssessmentModeTypesAutomaticByPlatform,
		"imagedefault":        AssessmentModeTypesImageDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssessmentModeTypes(input)
	return &out, nil
}

type InstanceViewTypes string

const (
	InstanceViewTypesInstanceView InstanceViewTypes = "instanceView"
)

func PossibleValuesForInstanceViewTypes() []string {
	return []string{
		string(InstanceViewTypesInstanceView),
	}
}

func parseInstanceViewTypes(input string) (*InstanceViewTypes, error) {
	vals := map[string]InstanceViewTypes{
		"instanceview": InstanceViewTypesInstanceView,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstanceViewTypes(input)
	return &out, nil
}

type LastAttemptStatusEnum string

const (
	LastAttemptStatusEnumFailed  LastAttemptStatusEnum = "Failed"
	LastAttemptStatusEnumSuccess LastAttemptStatusEnum = "Success"
)

func PossibleValuesForLastAttemptStatusEnum() []string {
	return []string{
		string(LastAttemptStatusEnumFailed),
		string(LastAttemptStatusEnumSuccess),
	}
}

func parseLastAttemptStatusEnum(input string) (*LastAttemptStatusEnum, error) {
	vals := map[string]LastAttemptStatusEnum{
		"failed":  LastAttemptStatusEnumFailed,
		"success": LastAttemptStatusEnumSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LastAttemptStatusEnum(input)
	return &out, nil
}

type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeLinux),
		string(OsTypeWindows),
	}
}

func parseOsType(input string) (*OsType, error) {
	vals := map[string]OsType{
		"linux":   OsTypeLinux,
		"windows": OsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsType(input)
	return &out, nil
}

type PatchModeTypes string

const (
	PatchModeTypesAutomaticByOS       PatchModeTypes = "AutomaticByOS"
	PatchModeTypesAutomaticByPlatform PatchModeTypes = "AutomaticByPlatform"
	PatchModeTypesImageDefault        PatchModeTypes = "ImageDefault"
	PatchModeTypesManual              PatchModeTypes = "Manual"
)

func PossibleValuesForPatchModeTypes() []string {
	return []string{
		string(PatchModeTypesAutomaticByOS),
		string(PatchModeTypesAutomaticByPlatform),
		string(PatchModeTypesImageDefault),
		string(PatchModeTypesManual),
	}
}

func parsePatchModeTypes(input string) (*PatchModeTypes, error) {
	vals := map[string]PatchModeTypes{
		"automaticbyos":       PatchModeTypesAutomaticByOS,
		"automaticbyplatform": PatchModeTypesAutomaticByPlatform,
		"imagedefault":        PatchModeTypesImageDefault,
		"manual":              PatchModeTypesManual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchModeTypes(input)
	return &out, nil
}

type PatchOperationStartedBy string

const (
	PatchOperationStartedByPlatform PatchOperationStartedBy = "Platform"
	PatchOperationStartedByUser     PatchOperationStartedBy = "User"
)

func PossibleValuesForPatchOperationStartedBy() []string {
	return []string{
		string(PatchOperationStartedByPlatform),
		string(PatchOperationStartedByUser),
	}
}

func parsePatchOperationStartedBy(input string) (*PatchOperationStartedBy, error) {
	vals := map[string]PatchOperationStartedBy{
		"platform": PatchOperationStartedByPlatform,
		"user":     PatchOperationStartedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchOperationStartedBy(input)
	return &out, nil
}

type PatchOperationStatus string

const (
	PatchOperationStatusCompletedWithWarnings PatchOperationStatus = "CompletedWithWarnings"
	PatchOperationStatusFailed                PatchOperationStatus = "Failed"
	PatchOperationStatusInProgress            PatchOperationStatus = "InProgress"
	PatchOperationStatusSucceeded             PatchOperationStatus = "Succeeded"
	PatchOperationStatusUnknown               PatchOperationStatus = "Unknown"
)

func PossibleValuesForPatchOperationStatus() []string {
	return []string{
		string(PatchOperationStatusCompletedWithWarnings),
		string(PatchOperationStatusFailed),
		string(PatchOperationStatusInProgress),
		string(PatchOperationStatusSucceeded),
		string(PatchOperationStatusUnknown),
	}
}

func parsePatchOperationStatus(input string) (*PatchOperationStatus, error) {
	vals := map[string]PatchOperationStatus{
		"completedwithwarnings": PatchOperationStatusCompletedWithWarnings,
		"failed":                PatchOperationStatusFailed,
		"inprogress":            PatchOperationStatusInProgress,
		"succeeded":             PatchOperationStatusSucceeded,
		"unknown":               PatchOperationStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchOperationStatus(input)
	return &out, nil
}

type PatchServiceUsed string

const (
	PatchServiceUsedAPT     PatchServiceUsed = "APT"
	PatchServiceUsedUnknown PatchServiceUsed = "Unknown"
	PatchServiceUsedWU      PatchServiceUsed = "WU"
	PatchServiceUsedWUWSUS  PatchServiceUsed = "WU_WSUS"
	PatchServiceUsedYUM     PatchServiceUsed = "YUM"
	PatchServiceUsedZypper  PatchServiceUsed = "Zypper"
)

func PossibleValuesForPatchServiceUsed() []string {
	return []string{
		string(PatchServiceUsedAPT),
		string(PatchServiceUsedUnknown),
		string(PatchServiceUsedWU),
		string(PatchServiceUsedWUWSUS),
		string(PatchServiceUsedYUM),
		string(PatchServiceUsedZypper),
	}
}

func parsePatchServiceUsed(input string) (*PatchServiceUsed, error) {
	vals := map[string]PatchServiceUsed{
		"apt":     PatchServiceUsedAPT,
		"unknown": PatchServiceUsedUnknown,
		"wu":      PatchServiceUsedWU,
		"wu_wsus": PatchServiceUsedWUWSUS,
		"yum":     PatchServiceUsedYUM,
		"zypper":  PatchServiceUsedZypper,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchServiceUsed(input)
	return &out, nil
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}

func parseStatusLevelTypes(input string) (*StatusLevelTypes, error) {
	vals := map[string]StatusLevelTypes{
		"error":   StatusLevelTypesError,
		"info":    StatusLevelTypesInfo,
		"warning": StatusLevelTypesWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusLevelTypes(input)
	return &out, nil
}

type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

func PossibleValuesForStatusTypes() []string {
	return []string{
		string(StatusTypesConnected),
		string(StatusTypesDisconnected),
		string(StatusTypesError),
	}
}

func parseStatusTypes(input string) (*StatusTypes, error) {
	vals := map[string]StatusTypes{
		"connected":    StatusTypesConnected,
		"disconnected": StatusTypesDisconnected,
		"error":        StatusTypesError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusTypes(input)
	return &out, nil
}

type VMGuestPatchClassificationLinux string

const (
	VMGuestPatchClassificationLinuxCritical VMGuestPatchClassificationLinux = "Critical"
	VMGuestPatchClassificationLinuxOther    VMGuestPatchClassificationLinux = "Other"
	VMGuestPatchClassificationLinuxSecurity VMGuestPatchClassificationLinux = "Security"
)

func PossibleValuesForVMGuestPatchClassificationLinux() []string {
	return []string{
		string(VMGuestPatchClassificationLinuxCritical),
		string(VMGuestPatchClassificationLinuxOther),
		string(VMGuestPatchClassificationLinuxSecurity),
	}
}

func parseVMGuestPatchClassificationLinux(input string) (*VMGuestPatchClassificationLinux, error) {
	vals := map[string]VMGuestPatchClassificationLinux{
		"critical": VMGuestPatchClassificationLinuxCritical,
		"other":    VMGuestPatchClassificationLinuxOther,
		"security": VMGuestPatchClassificationLinuxSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGuestPatchClassificationLinux(input)
	return &out, nil
}

type VMGuestPatchClassificationWindows string

const (
	VMGuestPatchClassificationWindowsCritical     VMGuestPatchClassificationWindows = "Critical"
	VMGuestPatchClassificationWindowsDefinition   VMGuestPatchClassificationWindows = "Definition"
	VMGuestPatchClassificationWindowsFeaturePack  VMGuestPatchClassificationWindows = "FeaturePack"
	VMGuestPatchClassificationWindowsSecurity     VMGuestPatchClassificationWindows = "Security"
	VMGuestPatchClassificationWindowsServicePack  VMGuestPatchClassificationWindows = "ServicePack"
	VMGuestPatchClassificationWindowsTools        VMGuestPatchClassificationWindows = "Tools"
	VMGuestPatchClassificationWindowsUpdateRollUp VMGuestPatchClassificationWindows = "UpdateRollUp"
	VMGuestPatchClassificationWindowsUpdates      VMGuestPatchClassificationWindows = "Updates"
)

func PossibleValuesForVMGuestPatchClassificationWindows() []string {
	return []string{
		string(VMGuestPatchClassificationWindowsCritical),
		string(VMGuestPatchClassificationWindowsDefinition),
		string(VMGuestPatchClassificationWindowsFeaturePack),
		string(VMGuestPatchClassificationWindowsSecurity),
		string(VMGuestPatchClassificationWindowsServicePack),
		string(VMGuestPatchClassificationWindowsTools),
		string(VMGuestPatchClassificationWindowsUpdateRollUp),
		string(VMGuestPatchClassificationWindowsUpdates),
	}
}

func parseVMGuestPatchClassificationWindows(input string) (*VMGuestPatchClassificationWindows, error) {
	vals := map[string]VMGuestPatchClassificationWindows{
		"critical":     VMGuestPatchClassificationWindowsCritical,
		"definition":   VMGuestPatchClassificationWindowsDefinition,
		"featurepack":  VMGuestPatchClassificationWindowsFeaturePack,
		"security":     VMGuestPatchClassificationWindowsSecurity,
		"servicepack":  VMGuestPatchClassificationWindowsServicePack,
		"tools":        VMGuestPatchClassificationWindowsTools,
		"updaterollup": VMGuestPatchClassificationWindowsUpdateRollUp,
		"updates":      VMGuestPatchClassificationWindowsUpdates,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGuestPatchClassificationWindows(input)
	return &out, nil
}

type VMGuestPatchRebootSetting string

const (
	VMGuestPatchRebootSettingAlways     VMGuestPatchRebootSetting = "Always"
	VMGuestPatchRebootSettingIfRequired VMGuestPatchRebootSetting = "IfRequired"
	VMGuestPatchRebootSettingNever      VMGuestPatchRebootSetting = "Never"
)

func PossibleValuesForVMGuestPatchRebootSetting() []string {
	return []string{
		string(VMGuestPatchRebootSettingAlways),
		string(VMGuestPatchRebootSettingIfRequired),
		string(VMGuestPatchRebootSettingNever),
	}
}

func parseVMGuestPatchRebootSetting(input string) (*VMGuestPatchRebootSetting, error) {
	vals := map[string]VMGuestPatchRebootSetting{
		"always":     VMGuestPatchRebootSettingAlways,
		"ifrequired": VMGuestPatchRebootSettingIfRequired,
		"never":      VMGuestPatchRebootSettingNever,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGuestPatchRebootSetting(input)
	return &out, nil
}

type VMGuestPatchRebootStatus string

const (
	VMGuestPatchRebootStatusCompleted VMGuestPatchRebootStatus = "Completed"
	VMGuestPatchRebootStatusFailed    VMGuestPatchRebootStatus = "Failed"
	VMGuestPatchRebootStatusNotNeeded VMGuestPatchRebootStatus = "NotNeeded"
	VMGuestPatchRebootStatusRequired  VMGuestPatchRebootStatus = "Required"
	VMGuestPatchRebootStatusStarted   VMGuestPatchRebootStatus = "Started"
	VMGuestPatchRebootStatusUnknown   VMGuestPatchRebootStatus = "Unknown"
)

func PossibleValuesForVMGuestPatchRebootStatus() []string {
	return []string{
		string(VMGuestPatchRebootStatusCompleted),
		string(VMGuestPatchRebootStatusFailed),
		string(VMGuestPatchRebootStatusNotNeeded),
		string(VMGuestPatchRebootStatusRequired),
		string(VMGuestPatchRebootStatusStarted),
		string(VMGuestPatchRebootStatusUnknown),
	}
}

func parseVMGuestPatchRebootStatus(input string) (*VMGuestPatchRebootStatus, error) {
	vals := map[string]VMGuestPatchRebootStatus{
		"completed": VMGuestPatchRebootStatusCompleted,
		"failed":    VMGuestPatchRebootStatusFailed,
		"notneeded": VMGuestPatchRebootStatusNotNeeded,
		"required":  VMGuestPatchRebootStatusRequired,
		"started":   VMGuestPatchRebootStatusStarted,
		"unknown":   VMGuestPatchRebootStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGuestPatchRebootStatus(input)
	return &out, nil
}
