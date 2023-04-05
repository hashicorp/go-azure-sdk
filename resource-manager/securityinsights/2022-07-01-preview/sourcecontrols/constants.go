package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentType string

const (
	ContentTypeAnalyticRule ContentType = "AnalyticRule"
	ContentTypeWorkbook     ContentType = "Workbook"
)

func PossibleValuesForContentType() []string {
	return []string{
		string(ContentTypeAnalyticRule),
		string(ContentTypeWorkbook),
	}
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

type RepoType string

const (
	RepoTypeDevOps RepoType = "DevOps"
	RepoTypeGithub RepoType = "Github"
)

func PossibleValuesForRepoType() []string {
	return []string{
		string(RepoTypeDevOps),
		string(RepoTypeGithub),
	}
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
