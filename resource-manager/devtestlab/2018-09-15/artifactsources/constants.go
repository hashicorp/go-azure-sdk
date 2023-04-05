package artifactsources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableStatus string

const (
	EnableStatusDisabled EnableStatus = "Disabled"
	EnableStatusEnabled  EnableStatus = "Enabled"
)

func PossibleValuesForEnableStatus() []string {
	return []string{
		string(EnableStatusDisabled),
		string(EnableStatusEnabled),
	}
}

type SourceControlType string

const (
	SourceControlTypeGitHub         SourceControlType = "GitHub"
	SourceControlTypeStorageAccount SourceControlType = "StorageAccount"
	SourceControlTypeVsoGit         SourceControlType = "VsoGit"
)

func PossibleValuesForSourceControlType() []string {
	return []string{
		string(SourceControlTypeGitHub),
		string(SourceControlTypeStorageAccount),
		string(SourceControlTypeVsoGit),
	}
}
