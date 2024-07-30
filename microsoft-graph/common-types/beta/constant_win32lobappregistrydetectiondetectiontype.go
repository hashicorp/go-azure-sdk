package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryDetectionDetectionType string

const (
	Win32LobAppRegistryDetectionDetectionType_DoesNotExist  Win32LobAppRegistryDetectionDetectionType = "doesNotExist"
	Win32LobAppRegistryDetectionDetectionType_Exists        Win32LobAppRegistryDetectionDetectionType = "exists"
	Win32LobAppRegistryDetectionDetectionType_Integer       Win32LobAppRegistryDetectionDetectionType = "integer"
	Win32LobAppRegistryDetectionDetectionType_NotConfigured Win32LobAppRegistryDetectionDetectionType = "notConfigured"
	Win32LobAppRegistryDetectionDetectionType_String        Win32LobAppRegistryDetectionDetectionType = "string"
	Win32LobAppRegistryDetectionDetectionType_Version       Win32LobAppRegistryDetectionDetectionType = "version"
)

func PossibleValuesForWin32LobAppRegistryDetectionDetectionType() []string {
	return []string{
		string(Win32LobAppRegistryDetectionDetectionType_DoesNotExist),
		string(Win32LobAppRegistryDetectionDetectionType_Exists),
		string(Win32LobAppRegistryDetectionDetectionType_Integer),
		string(Win32LobAppRegistryDetectionDetectionType_NotConfigured),
		string(Win32LobAppRegistryDetectionDetectionType_String),
		string(Win32LobAppRegistryDetectionDetectionType_Version),
	}
}

func (s *Win32LobAppRegistryDetectionDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryDetectionDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryDetectionDetectionType(input string) (*Win32LobAppRegistryDetectionDetectionType, error) {
	vals := map[string]Win32LobAppRegistryDetectionDetectionType{
		"doesnotexist":  Win32LobAppRegistryDetectionDetectionType_DoesNotExist,
		"exists":        Win32LobAppRegistryDetectionDetectionType_Exists,
		"integer":       Win32LobAppRegistryDetectionDetectionType_Integer,
		"notconfigured": Win32LobAppRegistryDetectionDetectionType_NotConfigured,
		"string":        Win32LobAppRegistryDetectionDetectionType_String,
		"version":       Win32LobAppRegistryDetectionDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryDetectionDetectionType(input)
	return &out, nil
}
