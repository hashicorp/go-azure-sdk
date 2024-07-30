package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRequirementDetectionType string

const (
	Win32LobAppFileSystemRequirementDetectionType_CreatedDate   Win32LobAppFileSystemRequirementDetectionType = "createdDate"
	Win32LobAppFileSystemRequirementDetectionType_DoesNotExist  Win32LobAppFileSystemRequirementDetectionType = "doesNotExist"
	Win32LobAppFileSystemRequirementDetectionType_Exists        Win32LobAppFileSystemRequirementDetectionType = "exists"
	Win32LobAppFileSystemRequirementDetectionType_ModifiedDate  Win32LobAppFileSystemRequirementDetectionType = "modifiedDate"
	Win32LobAppFileSystemRequirementDetectionType_NotConfigured Win32LobAppFileSystemRequirementDetectionType = "notConfigured"
	Win32LobAppFileSystemRequirementDetectionType_SizeInMB      Win32LobAppFileSystemRequirementDetectionType = "sizeInMB"
	Win32LobAppFileSystemRequirementDetectionType_Version       Win32LobAppFileSystemRequirementDetectionType = "version"
)

func PossibleValuesForWin32LobAppFileSystemRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppFileSystemRequirementDetectionType_CreatedDate),
		string(Win32LobAppFileSystemRequirementDetectionType_DoesNotExist),
		string(Win32LobAppFileSystemRequirementDetectionType_Exists),
		string(Win32LobAppFileSystemRequirementDetectionType_ModifiedDate),
		string(Win32LobAppFileSystemRequirementDetectionType_NotConfigured),
		string(Win32LobAppFileSystemRequirementDetectionType_SizeInMB),
		string(Win32LobAppFileSystemRequirementDetectionType_Version),
	}
}

func (s *Win32LobAppFileSystemRequirementDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemRequirementDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemRequirementDetectionType(input string) (*Win32LobAppFileSystemRequirementDetectionType, error) {
	vals := map[string]Win32LobAppFileSystemRequirementDetectionType{
		"createddate":   Win32LobAppFileSystemRequirementDetectionType_CreatedDate,
		"doesnotexist":  Win32LobAppFileSystemRequirementDetectionType_DoesNotExist,
		"exists":        Win32LobAppFileSystemRequirementDetectionType_Exists,
		"modifieddate":  Win32LobAppFileSystemRequirementDetectionType_ModifiedDate,
		"notconfigured": Win32LobAppFileSystemRequirementDetectionType_NotConfigured,
		"sizeinmb":      Win32LobAppFileSystemRequirementDetectionType_SizeInMB,
		"version":       Win32LobAppFileSystemRequirementDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRequirementDetectionType(input)
	return &out, nil
}
