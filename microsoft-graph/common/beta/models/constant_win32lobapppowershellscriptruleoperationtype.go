package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleOperationType string

const (
	Win32LobAppPowerShellScriptRuleOperationTypeboolean       Win32LobAppPowerShellScriptRuleOperationType = "Boolean"
	Win32LobAppPowerShellScriptRuleOperationTypedateTime      Win32LobAppPowerShellScriptRuleOperationType = "DateTime"
	Win32LobAppPowerShellScriptRuleOperationTypefloat         Win32LobAppPowerShellScriptRuleOperationType = "Float"
	Win32LobAppPowerShellScriptRuleOperationTypeinteger       Win32LobAppPowerShellScriptRuleOperationType = "Integer"
	Win32LobAppPowerShellScriptRuleOperationTypenotConfigured Win32LobAppPowerShellScriptRuleOperationType = "NotConfigured"
	Win32LobAppPowerShellScriptRuleOperationTypestring        Win32LobAppPowerShellScriptRuleOperationType = "String"
	Win32LobAppPowerShellScriptRuleOperationTypeversion       Win32LobAppPowerShellScriptRuleOperationType = "Version"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleOperationType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleOperationTypeboolean),
		string(Win32LobAppPowerShellScriptRuleOperationTypedateTime),
		string(Win32LobAppPowerShellScriptRuleOperationTypefloat),
		string(Win32LobAppPowerShellScriptRuleOperationTypeinteger),
		string(Win32LobAppPowerShellScriptRuleOperationTypenotConfigured),
		string(Win32LobAppPowerShellScriptRuleOperationTypestring),
		string(Win32LobAppPowerShellScriptRuleOperationTypeversion),
	}
}

func parseWin32LobAppPowerShellScriptRuleOperationType(input string) (*Win32LobAppPowerShellScriptRuleOperationType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleOperationType{
		"boolean":       Win32LobAppPowerShellScriptRuleOperationTypeboolean,
		"datetime":      Win32LobAppPowerShellScriptRuleOperationTypedateTime,
		"float":         Win32LobAppPowerShellScriptRuleOperationTypefloat,
		"integer":       Win32LobAppPowerShellScriptRuleOperationTypeinteger,
		"notconfigured": Win32LobAppPowerShellScriptRuleOperationTypenotConfigured,
		"string":        Win32LobAppPowerShellScriptRuleOperationTypestring,
		"version":       Win32LobAppPowerShellScriptRuleOperationTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleOperationType(input)
	return &out, nil
}
