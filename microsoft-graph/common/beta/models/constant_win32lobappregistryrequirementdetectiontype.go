package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRequirementDetectionType string

const (
	Win32LobAppRegistryRequirementDetectionTypedoesNotExist  Win32LobAppRegistryRequirementDetectionType = "DoesNotExist"
	Win32LobAppRegistryRequirementDetectionTypeexists        Win32LobAppRegistryRequirementDetectionType = "Exists"
	Win32LobAppRegistryRequirementDetectionTypeinteger       Win32LobAppRegistryRequirementDetectionType = "Integer"
	Win32LobAppRegistryRequirementDetectionTypenotConfigured Win32LobAppRegistryRequirementDetectionType = "NotConfigured"
	Win32LobAppRegistryRequirementDetectionTypestring        Win32LobAppRegistryRequirementDetectionType = "String"
	Win32LobAppRegistryRequirementDetectionTypeversion       Win32LobAppRegistryRequirementDetectionType = "Version"
)

func PossibleValuesForWin32LobAppRegistryRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppRegistryRequirementDetectionTypedoesNotExist),
		string(Win32LobAppRegistryRequirementDetectionTypeexists),
		string(Win32LobAppRegistryRequirementDetectionTypeinteger),
		string(Win32LobAppRegistryRequirementDetectionTypenotConfigured),
		string(Win32LobAppRegistryRequirementDetectionTypestring),
		string(Win32LobAppRegistryRequirementDetectionTypeversion),
	}
}

func parseWin32LobAppRegistryRequirementDetectionType(input string) (*Win32LobAppRegistryRequirementDetectionType, error) {
	vals := map[string]Win32LobAppRegistryRequirementDetectionType{
		"doesnotexist":  Win32LobAppRegistryRequirementDetectionTypedoesNotExist,
		"exists":        Win32LobAppRegistryRequirementDetectionTypeexists,
		"integer":       Win32LobAppRegistryRequirementDetectionTypeinteger,
		"notconfigured": Win32LobAppRegistryRequirementDetectionTypenotConfigured,
		"string":        Win32LobAppRegistryRequirementDetectionTypestring,
		"version":       Win32LobAppRegistryRequirementDetectionTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRequirementDetectionType(input)
	return &out, nil
}
