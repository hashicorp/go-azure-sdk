package sourcecontrols

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentType string

const (
	ContentTypeAnalyticRule   ContentType = "AnalyticRule"
	ContentTypeAutomationRule ContentType = "AutomationRule"
	ContentTypeHuntingQuery   ContentType = "HuntingQuery"
	ContentTypeParser         ContentType = "Parser"
	ContentTypePlaybook       ContentType = "Playbook"
	ContentTypeWorkbook       ContentType = "Workbook"
)

func PossibleValuesForContentType() []string {
	return []string{
		string(ContentTypeAnalyticRule),
		string(ContentTypeAutomationRule),
		string(ContentTypeHuntingQuery),
		string(ContentTypeParser),
		string(ContentTypePlaybook),
		string(ContentTypeWorkbook),
	}
}

func parseContentType(input string) (*ContentType, error) {
	vals := map[string]ContentType{
		"analyticrule":   ContentTypeAnalyticRule,
		"automationrule": ContentTypeAutomationRule,
		"huntingquery":   ContentTypeHuntingQuery,
		"parser":         ContentTypeParser,
		"playbook":       ContentTypePlaybook,
		"workbook":       ContentTypeWorkbook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentType(input)
	return &out, nil
}

type DeploymentFetchStatus string

const (
	DeploymentFetchStatusNotFound     DeploymentFetchStatus = "NotFound"
	DeploymentFetchStatusSuccess      DeploymentFetchStatus = "Success"
	DeploymentFetchStatusUnauthorized DeploymentFetchStatus = "Unauthorized"
)

func PossibleValuesForDeploymentFetchStatus() []string {
	return []string{
		string(DeploymentFetchStatusNotFound),
		string(DeploymentFetchStatusSuccess),
		string(DeploymentFetchStatusUnauthorized),
	}
}

func parseDeploymentFetchStatus(input string) (*DeploymentFetchStatus, error) {
	vals := map[string]DeploymentFetchStatus{
		"notfound":     DeploymentFetchStatusNotFound,
		"success":      DeploymentFetchStatusSuccess,
		"unauthorized": DeploymentFetchStatusUnauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentFetchStatus(input)
	return &out, nil
}

type DeploymentResult string

const (
	DeploymentResultCanceled DeploymentResult = "Canceled"
	DeploymentResultFailed   DeploymentResult = "Failed"
	DeploymentResultSuccess  DeploymentResult = "Success"
)

func PossibleValuesForDeploymentResult() []string {
	return []string{
		string(DeploymentResultCanceled),
		string(DeploymentResultFailed),
		string(DeploymentResultSuccess),
	}
}

func parseDeploymentResult(input string) (*DeploymentResult, error) {
	vals := map[string]DeploymentResult{
		"canceled": DeploymentResultCanceled,
		"failed":   DeploymentResultFailed,
		"success":  DeploymentResultSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentResult(input)
	return &out, nil
}

type DeploymentState string

const (
	DeploymentStateCanceling  DeploymentState = "Canceling"
	DeploymentStateCompleted  DeploymentState = "Completed"
	DeploymentStateInProgress DeploymentState = "In_Progress"
	DeploymentStateQueued     DeploymentState = "Queued"
)

func PossibleValuesForDeploymentState() []string {
	return []string{
		string(DeploymentStateCanceling),
		string(DeploymentStateCompleted),
		string(DeploymentStateInProgress),
		string(DeploymentStateQueued),
	}
}

func parseDeploymentState(input string) (*DeploymentState, error) {
	vals := map[string]DeploymentState{
		"canceling":   DeploymentStateCanceling,
		"completed":   DeploymentStateCompleted,
		"in_progress": DeploymentStateInProgress,
		"queued":      DeploymentStateQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentState(input)
	return &out, nil
}

type RepoType string

const (
	RepoTypeAzureDevOps RepoType = "AzureDevOps"
	RepoTypeGithub      RepoType = "Github"
)

func PossibleValuesForRepoType() []string {
	return []string{
		string(RepoTypeAzureDevOps),
		string(RepoTypeGithub),
	}
}

func parseRepoType(input string) (*RepoType, error) {
	vals := map[string]RepoType{
		"azuredevops": RepoTypeAzureDevOps,
		"github":      RepoTypeGithub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepoType(input)
	return &out, nil
}

type RepositoryAccessKind string

const (
	RepositoryAccessKindApp   RepositoryAccessKind = "App"
	RepositoryAccessKindOAuth RepositoryAccessKind = "OAuth"
	RepositoryAccessKindPAT   RepositoryAccessKind = "PAT"
)

func PossibleValuesForRepositoryAccessKind() []string {
	return []string{
		string(RepositoryAccessKindApp),
		string(RepositoryAccessKindOAuth),
		string(RepositoryAccessKindPAT),
	}
}

func parseRepositoryAccessKind(input string) (*RepositoryAccessKind, error) {
	vals := map[string]RepositoryAccessKind{
		"app":   RepositoryAccessKindApp,
		"oauth": RepositoryAccessKindOAuth,
		"pat":   RepositoryAccessKindPAT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepositoryAccessKind(input)
	return &out, nil
}

type State string

const (
	StateClosed State = "Closed"
	StateOpen   State = "Open"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateClosed),
		string(StateOpen),
	}
}

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"closed": StateClosed,
		"open":   StateOpen,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}

type Version string

const (
	VersionVOne Version = "V1"
	VersionVTwo Version = "V2"
)

func PossibleValuesForVersion() []string {
	return []string{
		string(VersionVOne),
		string(VersionVTwo),
	}
}

func parseVersion(input string) (*Version, error) {
	vals := map[string]Version{
		"v1": VersionVOne,
		"v2": VersionVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Version(input)
	return &out, nil
}

type WarningCode string

const (
	WarningCodeSourceControlDeletedWithWarnings                      WarningCode = "SourceControl_DeletedWithWarnings"
	WarningCodeSourceControlWarningDeletePipelineFromAzureDevOps     WarningCode = "SourceControlWarning_DeletePipelineFromAzureDevOps"
	WarningCodeSourceControlWarningDeleteRoleAssignment              WarningCode = "SourceControlWarning_DeleteRoleAssignment"
	WarningCodeSourceControlWarningDeleteServicePrincipal            WarningCode = "SourceControlWarning_DeleteServicePrincipal"
	WarningCodeSourceControlWarningDeleteWorkflowAndSecretFromGitHub WarningCode = "SourceControlWarning_DeleteWorkflowAndSecretFromGitHub"
)

func PossibleValuesForWarningCode() []string {
	return []string{
		string(WarningCodeSourceControlDeletedWithWarnings),
		string(WarningCodeSourceControlWarningDeletePipelineFromAzureDevOps),
		string(WarningCodeSourceControlWarningDeleteRoleAssignment),
		string(WarningCodeSourceControlWarningDeleteServicePrincipal),
		string(WarningCodeSourceControlWarningDeleteWorkflowAndSecretFromGitHub),
	}
}

func parseWarningCode(input string) (*WarningCode, error) {
	vals := map[string]WarningCode{
		"sourcecontrol_deletedwithwarnings":                      WarningCodeSourceControlDeletedWithWarnings,
		"sourcecontrolwarning_deletepipelinefromazuredevops":     WarningCodeSourceControlWarningDeletePipelineFromAzureDevOps,
		"sourcecontrolwarning_deleteroleassignment":              WarningCodeSourceControlWarningDeleteRoleAssignment,
		"sourcecontrolwarning_deleteserviceprincipal":            WarningCodeSourceControlWarningDeleteServicePrincipal,
		"sourcecontrolwarning_deleteworkflowandsecretfromgithub": WarningCodeSourceControlWarningDeleteWorkflowAndSecretFromGitHub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WarningCode(input)
	return &out, nil
}
