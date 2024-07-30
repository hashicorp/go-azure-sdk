package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetectionDetectionType string

const (
	Win32LobAppFileSystemDetectionDetectionType_CreatedDate   Win32LobAppFileSystemDetectionDetectionType = "createdDate"
	Win32LobAppFileSystemDetectionDetectionType_DoesNotExist  Win32LobAppFileSystemDetectionDetectionType = "doesNotExist"
	Win32LobAppFileSystemDetectionDetectionType_Exists        Win32LobAppFileSystemDetectionDetectionType = "exists"
	Win32LobAppFileSystemDetectionDetectionType_ModifiedDate  Win32LobAppFileSystemDetectionDetectionType = "modifiedDate"
	Win32LobAppFileSystemDetectionDetectionType_NotConfigured Win32LobAppFileSystemDetectionDetectionType = "notConfigured"
	Win32LobAppFileSystemDetectionDetectionType_SizeInMB      Win32LobAppFileSystemDetectionDetectionType = "sizeInMB"
	Win32LobAppFileSystemDetectionDetectionType_Version       Win32LobAppFileSystemDetectionDetectionType = "version"
)

func PossibleValuesForWin32LobAppFileSystemDetectionDetectionType() []string {
	return []string{
		string(Win32LobAppFileSystemDetectionDetectionType_CreatedDate),
		string(Win32LobAppFileSystemDetectionDetectionType_DoesNotExist),
		string(Win32LobAppFileSystemDetectionDetectionType_Exists),
		string(Win32LobAppFileSystemDetectionDetectionType_ModifiedDate),
		string(Win32LobAppFileSystemDetectionDetectionType_NotConfigured),
		string(Win32LobAppFileSystemDetectionDetectionType_SizeInMB),
		string(Win32LobAppFileSystemDetectionDetectionType_Version),
	}
}

func (s *Win32LobAppFileSystemDetectionDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemDetectionDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemDetectionDetectionType(input string) (*Win32LobAppFileSystemDetectionDetectionType, error) {
	vals := map[string]Win32LobAppFileSystemDetectionDetectionType{
		"createddate":   Win32LobAppFileSystemDetectionDetectionType_CreatedDate,
		"doesnotexist":  Win32LobAppFileSystemDetectionDetectionType_DoesNotExist,
		"exists":        Win32LobAppFileSystemDetectionDetectionType_Exists,
		"modifieddate":  Win32LobAppFileSystemDetectionDetectionType_ModifiedDate,
		"notconfigured": Win32LobAppFileSystemDetectionDetectionType_NotConfigured,
		"sizeinmb":      Win32LobAppFileSystemDetectionDetectionType_SizeInMB,
		"version":       Win32LobAppFileSystemDetectionDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemDetectionDetectionType(input)
	return &out, nil
}
