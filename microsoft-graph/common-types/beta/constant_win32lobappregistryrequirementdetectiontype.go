package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRequirementDetectionType string

const (
	Win32LobAppRegistryRequirementDetectionType_DoesNotExist  Win32LobAppRegistryRequirementDetectionType = "doesNotExist"
	Win32LobAppRegistryRequirementDetectionType_Exists        Win32LobAppRegistryRequirementDetectionType = "exists"
	Win32LobAppRegistryRequirementDetectionType_Integer       Win32LobAppRegistryRequirementDetectionType = "integer"
	Win32LobAppRegistryRequirementDetectionType_NotConfigured Win32LobAppRegistryRequirementDetectionType = "notConfigured"
	Win32LobAppRegistryRequirementDetectionType_String        Win32LobAppRegistryRequirementDetectionType = "string"
	Win32LobAppRegistryRequirementDetectionType_Version       Win32LobAppRegistryRequirementDetectionType = "version"
)

func PossibleValuesForWin32LobAppRegistryRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppRegistryRequirementDetectionType_DoesNotExist),
		string(Win32LobAppRegistryRequirementDetectionType_Exists),
		string(Win32LobAppRegistryRequirementDetectionType_Integer),
		string(Win32LobAppRegistryRequirementDetectionType_NotConfigured),
		string(Win32LobAppRegistryRequirementDetectionType_String),
		string(Win32LobAppRegistryRequirementDetectionType_Version),
	}
}

func (s *Win32LobAppRegistryRequirementDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryRequirementDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryRequirementDetectionType(input string) (*Win32LobAppRegistryRequirementDetectionType, error) {
	vals := map[string]Win32LobAppRegistryRequirementDetectionType{
		"doesnotexist":  Win32LobAppRegistryRequirementDetectionType_DoesNotExist,
		"exists":        Win32LobAppRegistryRequirementDetectionType_Exists,
		"integer":       Win32LobAppRegistryRequirementDetectionType_Integer,
		"notconfigured": Win32LobAppRegistryRequirementDetectionType_NotConfigured,
		"string":        Win32LobAppRegistryRequirementDetectionType_String,
		"version":       Win32LobAppRegistryRequirementDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRequirementDetectionType(input)
	return &out, nil
}
