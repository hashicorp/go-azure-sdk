package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRequirementDetectionType string

const (
	Win32LobAppFileSystemRequirementDetectionTypecreatedDate   Win32LobAppFileSystemRequirementDetectionType = "CreatedDate"
	Win32LobAppFileSystemRequirementDetectionTypedoesNotExist  Win32LobAppFileSystemRequirementDetectionType = "DoesNotExist"
	Win32LobAppFileSystemRequirementDetectionTypeexists        Win32LobAppFileSystemRequirementDetectionType = "Exists"
	Win32LobAppFileSystemRequirementDetectionTypemodifiedDate  Win32LobAppFileSystemRequirementDetectionType = "ModifiedDate"
	Win32LobAppFileSystemRequirementDetectionTypenotConfigured Win32LobAppFileSystemRequirementDetectionType = "NotConfigured"
	Win32LobAppFileSystemRequirementDetectionTypesizeInMB      Win32LobAppFileSystemRequirementDetectionType = "SizeInMB"
	Win32LobAppFileSystemRequirementDetectionTypeversion       Win32LobAppFileSystemRequirementDetectionType = "Version"
)

func PossibleValuesForWin32LobAppFileSystemRequirementDetectionType() []string {
	return []string{
		string(Win32LobAppFileSystemRequirementDetectionTypecreatedDate),
		string(Win32LobAppFileSystemRequirementDetectionTypedoesNotExist),
		string(Win32LobAppFileSystemRequirementDetectionTypeexists),
		string(Win32LobAppFileSystemRequirementDetectionTypemodifiedDate),
		string(Win32LobAppFileSystemRequirementDetectionTypenotConfigured),
		string(Win32LobAppFileSystemRequirementDetectionTypesizeInMB),
		string(Win32LobAppFileSystemRequirementDetectionTypeversion),
	}
}

func parseWin32LobAppFileSystemRequirementDetectionType(input string) (*Win32LobAppFileSystemRequirementDetectionType, error) {
	vals := map[string]Win32LobAppFileSystemRequirementDetectionType{
		"createddate":   Win32LobAppFileSystemRequirementDetectionTypecreatedDate,
		"doesnotexist":  Win32LobAppFileSystemRequirementDetectionTypedoesNotExist,
		"exists":        Win32LobAppFileSystemRequirementDetectionTypeexists,
		"modifieddate":  Win32LobAppFileSystemRequirementDetectionTypemodifiedDate,
		"notconfigured": Win32LobAppFileSystemRequirementDetectionTypenotConfigured,
		"sizeinmb":      Win32LobAppFileSystemRequirementDetectionTypesizeInMB,
		"version":       Win32LobAppFileSystemRequirementDetectionTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRequirementDetectionType(input)
	return &out, nil
}
