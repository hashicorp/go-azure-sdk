package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryDetectionDetectionType string

const (
	Win32LobAppRegistryDetectionDetectionTypedoesNotExist  Win32LobAppRegistryDetectionDetectionType = "DoesNotExist"
	Win32LobAppRegistryDetectionDetectionTypeexists        Win32LobAppRegistryDetectionDetectionType = "Exists"
	Win32LobAppRegistryDetectionDetectionTypeinteger       Win32LobAppRegistryDetectionDetectionType = "Integer"
	Win32LobAppRegistryDetectionDetectionTypenotConfigured Win32LobAppRegistryDetectionDetectionType = "NotConfigured"
	Win32LobAppRegistryDetectionDetectionTypestring        Win32LobAppRegistryDetectionDetectionType = "String"
	Win32LobAppRegistryDetectionDetectionTypeversion       Win32LobAppRegistryDetectionDetectionType = "Version"
)

func PossibleValuesForWin32LobAppRegistryDetectionDetectionType() []string {
	return []string{
		string(Win32LobAppRegistryDetectionDetectionTypedoesNotExist),
		string(Win32LobAppRegistryDetectionDetectionTypeexists),
		string(Win32LobAppRegistryDetectionDetectionTypeinteger),
		string(Win32LobAppRegistryDetectionDetectionTypenotConfigured),
		string(Win32LobAppRegistryDetectionDetectionTypestring),
		string(Win32LobAppRegistryDetectionDetectionTypeversion),
	}
}

func parseWin32LobAppRegistryDetectionDetectionType(input string) (*Win32LobAppRegistryDetectionDetectionType, error) {
	vals := map[string]Win32LobAppRegistryDetectionDetectionType{
		"doesnotexist":  Win32LobAppRegistryDetectionDetectionTypedoesNotExist,
		"exists":        Win32LobAppRegistryDetectionDetectionTypeexists,
		"integer":       Win32LobAppRegistryDetectionDetectionTypeinteger,
		"notconfigured": Win32LobAppRegistryDetectionDetectionTypenotConfigured,
		"string":        Win32LobAppRegistryDetectionDetectionTypestring,
		"version":       Win32LobAppRegistryDetectionDetectionTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryDetectionDetectionType(input)
	return &out, nil
}
