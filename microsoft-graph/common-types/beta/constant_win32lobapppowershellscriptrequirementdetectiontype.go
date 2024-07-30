package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementDetectionType string

const (
	Win32LobAppPowerShellScriptRequirementDetectionType_Boolean       Win32LobAppPowerShellScriptRequirementDetectionType = "boolean"
	Win32LobAppPowerShellScriptRequirementDetectionType_DateTime      Win32LobAppPowerShellScriptRequirementDetectionType = "dateTime"
	Win32LobAppPowerShellScriptRequirementDetectionType_Float         Win32LobAppPowerShellScriptRequirementDetectionType = "float"
	Win32LobAppPowerShellScriptRequirementDetectionType_Integer       Win32LobAppPowerShellScriptRequirementDetectionType = "integer"
	Win32LobAppPowerShellScriptRequirementDetectionType_NotConfigured Win32LobAppPowerShellScriptRequirementDetectionType = "notConfigured"
	Win32LobAppPowerShellScriptRequirementDetectionType_String        Win32LobAppPowerShellScriptRequirementDetectionType = "string"
	Win32LobAppPowerShellScriptRequirementDetectionType_Version       Win32LobAppPowerShellScriptRequirementDetectionType = "version"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementDetectionType_Boolean),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_DateTime),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_Float),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_Integer),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_NotConfigured),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_String),
		string(Win32LobAppPowerShellScriptRequirementDetectionType_Version),
	}
}

func (s *Win32LobAppPowerShellScriptRequirementDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRequirementDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRequirementDetectionType(input string) (*Win32LobAppPowerShellScriptRequirementDetectionType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementDetectionType{
		"boolean":       Win32LobAppPowerShellScriptRequirementDetectionType_Boolean,
		"datetime":      Win32LobAppPowerShellScriptRequirementDetectionType_DateTime,
		"float":         Win32LobAppPowerShellScriptRequirementDetectionType_Float,
		"integer":       Win32LobAppPowerShellScriptRequirementDetectionType_Integer,
		"notconfigured": Win32LobAppPowerShellScriptRequirementDetectionType_NotConfigured,
		"string":        Win32LobAppPowerShellScriptRequirementDetectionType_String,
		"version":       Win32LobAppPowerShellScriptRequirementDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementDetectionType(input)
	return &out, nil
}
