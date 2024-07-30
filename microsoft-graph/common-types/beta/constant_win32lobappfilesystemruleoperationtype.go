package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRuleOperationType string

const (
	Win32LobAppFileSystemRuleOperationType_CreatedDate   Win32LobAppFileSystemRuleOperationType = "createdDate"
	Win32LobAppFileSystemRuleOperationType_DoesNotExist  Win32LobAppFileSystemRuleOperationType = "doesNotExist"
	Win32LobAppFileSystemRuleOperationType_Exists        Win32LobAppFileSystemRuleOperationType = "exists"
	Win32LobAppFileSystemRuleOperationType_ModifiedDate  Win32LobAppFileSystemRuleOperationType = "modifiedDate"
	Win32LobAppFileSystemRuleOperationType_NotConfigured Win32LobAppFileSystemRuleOperationType = "notConfigured"
	Win32LobAppFileSystemRuleOperationType_SizeInMB      Win32LobAppFileSystemRuleOperationType = "sizeInMB"
	Win32LobAppFileSystemRuleOperationType_Version       Win32LobAppFileSystemRuleOperationType = "version"
)

func PossibleValuesForWin32LobAppFileSystemRuleOperationType() []string {
	return []string{
		string(Win32LobAppFileSystemRuleOperationType_CreatedDate),
		string(Win32LobAppFileSystemRuleOperationType_DoesNotExist),
		string(Win32LobAppFileSystemRuleOperationType_Exists),
		string(Win32LobAppFileSystemRuleOperationType_ModifiedDate),
		string(Win32LobAppFileSystemRuleOperationType_NotConfigured),
		string(Win32LobAppFileSystemRuleOperationType_SizeInMB),
		string(Win32LobAppFileSystemRuleOperationType_Version),
	}
}

func (s *Win32LobAppFileSystemRuleOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemRuleOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemRuleOperationType(input string) (*Win32LobAppFileSystemRuleOperationType, error) {
	vals := map[string]Win32LobAppFileSystemRuleOperationType{
		"createddate":   Win32LobAppFileSystemRuleOperationType_CreatedDate,
		"doesnotexist":  Win32LobAppFileSystemRuleOperationType_DoesNotExist,
		"exists":        Win32LobAppFileSystemRuleOperationType_Exists,
		"modifieddate":  Win32LobAppFileSystemRuleOperationType_ModifiedDate,
		"notconfigured": Win32LobAppFileSystemRuleOperationType_NotConfigured,
		"sizeinmb":      Win32LobAppFileSystemRuleOperationType_SizeInMB,
		"version":       Win32LobAppFileSystemRuleOperationType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRuleOperationType(input)
	return &out, nil
}
