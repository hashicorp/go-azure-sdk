package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementDetectionType string

const (
	Win32LobAppPowerShellScriptRequirementDetectionTypeboolean       Win32LobAppPowerShellScriptRequirementDetectionType = "Boolean"
	Win32LobAppPowerShellScriptRequirementDetectionTypedateTime      Win32LobAppPowerShellScriptRequirementDetectionType = "DateTime"
	Win32LobAppPowerShellScriptRequirementDetectionTypefloat         Win32LobAppPowerShellScriptRequirementDetectionType = "Float"
	Win32LobAppPowerShellScriptRequirementDetectionTypeinteger       Win32LobAppPowerShellScriptRequirementDetectionType = "Integer"
	Win32LobAppPowerShellScriptRequirementDetectionTypenotConfigured Win32LobAppPowerShellScriptRequirementDetectionType = "NotConfigured"
	Win32LobAppPowerShellScriptRequirementDetectionTypestring        Win32LobAppPowerShellScriptRequirementDetectionType = "String"
	Win32LobAppPowerShellScriptRequirementDetectionTypeversion       Win32LobAppPowerShellScriptRequirementDetectionType = "Version"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementDetectionTypeboolean),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypedateTime),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypefloat),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypeinteger),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypenotConfigured),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypestring),
		string(Win32LobAppPowerShellScriptRequirementDetectionTypeversion),
	}
}

func parseWin32LobAppPowerShellScriptRequirementDetectionType(input string) (*Win32LobAppPowerShellScriptRequirementDetectionType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementDetectionType{
		"boolean":       Win32LobAppPowerShellScriptRequirementDetectionTypeboolean,
		"datetime":      Win32LobAppPowerShellScriptRequirementDetectionTypedateTime,
		"float":         Win32LobAppPowerShellScriptRequirementDetectionTypefloat,
		"integer":       Win32LobAppPowerShellScriptRequirementDetectionTypeinteger,
		"notconfigured": Win32LobAppPowerShellScriptRequirementDetectionTypenotConfigured,
		"string":        Win32LobAppPowerShellScriptRequirementDetectionTypestring,
		"version":       Win32LobAppPowerShellScriptRequirementDetectionTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementDetectionType(input)
	return &out, nil
}
