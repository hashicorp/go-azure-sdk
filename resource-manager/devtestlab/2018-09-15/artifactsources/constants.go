package artifactsources

import "strings"

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

func parseEnableStatus(input string) (*EnableStatus, error) {
	vals := map[string]EnableStatus{
		"disabled": EnableStatusDisabled,
		"enabled":  EnableStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnableStatus(input)
	return &out, nil
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

func parseSourceControlType(input string) (*SourceControlType, error) {
	vals := map[string]SourceControlType{
		"github":         SourceControlTypeGitHub,
		"storageaccount": SourceControlTypeStorageAccount,
		"vsogit":         SourceControlTypeVsoGit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceControlType(input)
	return &out, nil
}
