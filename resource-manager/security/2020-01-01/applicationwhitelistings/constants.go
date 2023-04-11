package applicationwhitelistings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Action string

const (
	ActionAdd         Action = "Add"
	ActionRecommended Action = "Recommended"
	ActionRemove      Action = "Remove"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAdd),
		string(ActionRecommended),
		string(ActionRemove),
	}
}

func parseAction(input string) (*Action, error) {
	vals := map[string]Action{
		"add":         ActionAdd,
		"recommended": ActionRecommended,
		"remove":      ActionRemove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Action(input)
	return &out, nil
}

type ConfigurationStatus string

const (
	ConfigurationStatusConfigured    ConfigurationStatus = "Configured"
	ConfigurationStatusFailed        ConfigurationStatus = "Failed"
	ConfigurationStatusInProgress    ConfigurationStatus = "InProgress"
	ConfigurationStatusNoStatus      ConfigurationStatus = "NoStatus"
	ConfigurationStatusNotConfigured ConfigurationStatus = "NotConfigured"
)

func PossibleValuesForConfigurationStatus() []string {
	return []string{
		string(ConfigurationStatusConfigured),
		string(ConfigurationStatusFailed),
		string(ConfigurationStatusInProgress),
		string(ConfigurationStatusNoStatus),
		string(ConfigurationStatusNotConfigured),
	}
}

func parseConfigurationStatus(input string) (*ConfigurationStatus, error) {
	vals := map[string]ConfigurationStatus{
		"configured":    ConfigurationStatusConfigured,
		"failed":        ConfigurationStatusFailed,
		"inprogress":    ConfigurationStatusInProgress,
		"nostatus":      ConfigurationStatusNoStatus,
		"notconfigured": ConfigurationStatusNotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationStatus(input)
	return &out, nil
}

type EnforcementMode string

const (
	EnforcementModeAudit   EnforcementMode = "Audit"
	EnforcementModeEnforce EnforcementMode = "Enforce"
	EnforcementModeNone    EnforcementMode = "None"
)

func PossibleValuesForEnforcementMode() []string {
	return []string{
		string(EnforcementModeAudit),
		string(EnforcementModeEnforce),
		string(EnforcementModeNone),
	}
}

func parseEnforcementMode(input string) (*EnforcementMode, error) {
	vals := map[string]EnforcementMode{
		"audit":   EnforcementModeAudit,
		"enforce": EnforcementModeEnforce,
		"none":    EnforcementModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnforcementMode(input)
	return &out, nil
}

type EnforcementSupport string

const (
	EnforcementSupportNotSupported EnforcementSupport = "NotSupported"
	EnforcementSupportSupported    EnforcementSupport = "Supported"
	EnforcementSupportUnknown      EnforcementSupport = "Unknown"
)

func PossibleValuesForEnforcementSupport() []string {
	return []string{
		string(EnforcementSupportNotSupported),
		string(EnforcementSupportSupported),
		string(EnforcementSupportUnknown),
	}
}

func parseEnforcementSupport(input string) (*EnforcementSupport, error) {
	vals := map[string]EnforcementSupport{
		"notsupported": EnforcementSupportNotSupported,
		"supported":    EnforcementSupportSupported,
		"unknown":      EnforcementSupportUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnforcementSupport(input)
	return &out, nil
}

type Exe string

const (
	ExeAudit   Exe = "Audit"
	ExeEnforce Exe = "Enforce"
	ExeNone    Exe = "None"
)

func PossibleValuesForExe() []string {
	return []string{
		string(ExeAudit),
		string(ExeEnforce),
		string(ExeNone),
	}
}

func parseExe(input string) (*Exe, error) {
	vals := map[string]Exe{
		"audit":   ExeAudit,
		"enforce": ExeEnforce,
		"none":    ExeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Exe(input)
	return &out, nil
}

type Executable string

const (
	ExecutableAudit   Executable = "Audit"
	ExecutableEnforce Executable = "Enforce"
	ExecutableNone    Executable = "None"
)

func PossibleValuesForExecutable() []string {
	return []string{
		string(ExecutableAudit),
		string(ExecutableEnforce),
		string(ExecutableNone),
	}
}

func parseExecutable(input string) (*Executable, error) {
	vals := map[string]Executable{
		"audit":   ExecutableAudit,
		"enforce": ExecutableEnforce,
		"none":    ExecutableNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Executable(input)
	return &out, nil
}

type FileType string

const (
	FileTypeDll        FileType = "Dll"
	FileTypeExe        FileType = "Exe"
	FileTypeExecutable FileType = "Executable"
	FileTypeMsi        FileType = "Msi"
	FileTypeScript     FileType = "Script"
	FileTypeUnknown    FileType = "Unknown"
)

func PossibleValuesForFileType() []string {
	return []string{
		string(FileTypeDll),
		string(FileTypeExe),
		string(FileTypeExecutable),
		string(FileTypeMsi),
		string(FileTypeScript),
		string(FileTypeUnknown),
	}
}

func parseFileType(input string) (*FileType, error) {
	vals := map[string]FileType{
		"dll":        FileTypeDll,
		"exe":        FileTypeExe,
		"executable": FileTypeExecutable,
		"msi":        FileTypeMsi,
		"script":     FileTypeScript,
		"unknown":    FileTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileType(input)
	return &out, nil
}

type Issue string

const (
	IssueExecutableViolationsAudited   Issue = "ExecutableViolationsAudited"
	IssueMsiAndScriptViolationsAudited Issue = "MsiAndScriptViolationsAudited"
	IssueMsiAndScriptViolationsBlocked Issue = "MsiAndScriptViolationsBlocked"
	IssueRulesViolatedManually         Issue = "RulesViolatedManually"
	IssueViolationsAudited             Issue = "ViolationsAudited"
	IssueViolationsBlocked             Issue = "ViolationsBlocked"
)

func PossibleValuesForIssue() []string {
	return []string{
		string(IssueExecutableViolationsAudited),
		string(IssueMsiAndScriptViolationsAudited),
		string(IssueMsiAndScriptViolationsBlocked),
		string(IssueRulesViolatedManually),
		string(IssueViolationsAudited),
		string(IssueViolationsBlocked),
	}
}

func parseIssue(input string) (*Issue, error) {
	vals := map[string]Issue{
		"executableviolationsaudited":   IssueExecutableViolationsAudited,
		"msiandscriptviolationsaudited": IssueMsiAndScriptViolationsAudited,
		"msiandscriptviolationsblocked": IssueMsiAndScriptViolationsBlocked,
		"rulesviolatedmanually":         IssueRulesViolatedManually,
		"violationsaudited":             IssueViolationsAudited,
		"violationsblocked":             IssueViolationsBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Issue(input)
	return &out, nil
}

type Msi string

const (
	MsiAudit   Msi = "Audit"
	MsiEnforce Msi = "Enforce"
	MsiNone    Msi = "None"
)

func PossibleValuesForMsi() []string {
	return []string{
		string(MsiAudit),
		string(MsiEnforce),
		string(MsiNone),
	}
}

func parseMsi(input string) (*Msi, error) {
	vals := map[string]Msi{
		"audit":   MsiAudit,
		"enforce": MsiEnforce,
		"none":    MsiNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Msi(input)
	return &out, nil
}

type RecommendationAction string

const (
	RecommendationActionAdd         RecommendationAction = "Add"
	RecommendationActionRecommended RecommendationAction = "Recommended"
	RecommendationActionRemove      RecommendationAction = "Remove"
)

func PossibleValuesForRecommendationAction() []string {
	return []string{
		string(RecommendationActionAdd),
		string(RecommendationActionRecommended),
		string(RecommendationActionRemove),
	}
}

func parseRecommendationAction(input string) (*RecommendationAction, error) {
	vals := map[string]RecommendationAction{
		"add":         RecommendationActionAdd,
		"recommended": RecommendationActionRecommended,
		"remove":      RecommendationActionRemove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationAction(input)
	return &out, nil
}

type RecommendationStatus string

const (
	RecommendationStatusNoStatus       RecommendationStatus = "NoStatus"
	RecommendationStatusNotAvailable   RecommendationStatus = "NotAvailable"
	RecommendationStatusNotRecommended RecommendationStatus = "NotRecommended"
	RecommendationStatusRecommended    RecommendationStatus = "Recommended"
)

func PossibleValuesForRecommendationStatus() []string {
	return []string{
		string(RecommendationStatusNoStatus),
		string(RecommendationStatusNotAvailable),
		string(RecommendationStatusNotRecommended),
		string(RecommendationStatusRecommended),
	}
}

func parseRecommendationStatus(input string) (*RecommendationStatus, error) {
	vals := map[string]RecommendationStatus{
		"nostatus":       RecommendationStatusNoStatus,
		"notavailable":   RecommendationStatusNotAvailable,
		"notrecommended": RecommendationStatusNotRecommended,
		"recommended":    RecommendationStatusRecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationStatus(input)
	return &out, nil
}

type Script string

const (
	ScriptAudit   Script = "Audit"
	ScriptEnforce Script = "Enforce"
	ScriptNone    Script = "None"
)

func PossibleValuesForScript() []string {
	return []string{
		string(ScriptAudit),
		string(ScriptEnforce),
		string(ScriptNone),
	}
}

func parseScript(input string) (*Script, error) {
	vals := map[string]Script{
		"audit":   ScriptAudit,
		"enforce": ScriptEnforce,
		"none":    ScriptNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Script(input)
	return &out, nil
}

type SourceSystem string

const (
	SourceSystemAzureAppLocker    SourceSystem = "Azure_AppLocker"
	SourceSystemAzureAuditD       SourceSystem = "Azure_AuditD"
	SourceSystemNonAzureAppLocker SourceSystem = "NonAzure_AppLocker"
	SourceSystemNonAzureAuditD    SourceSystem = "NonAzure_AuditD"
	SourceSystemNone              SourceSystem = "None"
)

func PossibleValuesForSourceSystem() []string {
	return []string{
		string(SourceSystemAzureAppLocker),
		string(SourceSystemAzureAuditD),
		string(SourceSystemNonAzureAppLocker),
		string(SourceSystemNonAzureAuditD),
		string(SourceSystemNone),
	}
}

func parseSourceSystem(input string) (*SourceSystem, error) {
	vals := map[string]SourceSystem{
		"azure_applocker":    SourceSystemAzureAppLocker,
		"azure_auditd":       SourceSystemAzureAuditD,
		"nonazure_applocker": SourceSystemNonAzureAppLocker,
		"nonazure_auditd":    SourceSystemNonAzureAuditD,
		"none":               SourceSystemNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceSystem(input)
	return &out, nil
}

type Type string

const (
	TypeBinarySignature          Type = "BinarySignature"
	TypeFile                     Type = "File"
	TypeFileHash                 Type = "FileHash"
	TypeProductSignature         Type = "ProductSignature"
	TypePublisherSignature       Type = "PublisherSignature"
	TypeVersionAndAboveSignature Type = "VersionAndAboveSignature"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeBinarySignature),
		string(TypeFile),
		string(TypeFileHash),
		string(TypeProductSignature),
		string(TypePublisherSignature),
		string(TypeVersionAndAboveSignature),
	}
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"binarysignature":          TypeBinarySignature,
		"file":                     TypeFile,
		"filehash":                 TypeFileHash,
		"productsignature":         TypeProductSignature,
		"publishersignature":       TypePublisherSignature,
		"versionandabovesignature": TypeVersionAndAboveSignature,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
