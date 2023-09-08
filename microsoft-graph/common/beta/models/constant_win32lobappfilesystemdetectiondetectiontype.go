package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetectionDetectionType string

const (
	Win32LobAppFileSystemDetectionDetectionTypecreatedDate   Win32LobAppFileSystemDetectionDetectionType = "CreatedDate"
	Win32LobAppFileSystemDetectionDetectionTypedoesNotExist  Win32LobAppFileSystemDetectionDetectionType = "DoesNotExist"
	Win32LobAppFileSystemDetectionDetectionTypeexists        Win32LobAppFileSystemDetectionDetectionType = "Exists"
	Win32LobAppFileSystemDetectionDetectionTypemodifiedDate  Win32LobAppFileSystemDetectionDetectionType = "ModifiedDate"
	Win32LobAppFileSystemDetectionDetectionTypenotConfigured Win32LobAppFileSystemDetectionDetectionType = "NotConfigured"
	Win32LobAppFileSystemDetectionDetectionTypesizeInMB      Win32LobAppFileSystemDetectionDetectionType = "SizeInMB"
	Win32LobAppFileSystemDetectionDetectionTypeversion       Win32LobAppFileSystemDetectionDetectionType = "Version"
)

func PossibleValuesForWin32LobAppFileSystemDetectionDetectionType() []string {
	return []string{
		string(Win32LobAppFileSystemDetectionDetectionTypecreatedDate),
		string(Win32LobAppFileSystemDetectionDetectionTypedoesNotExist),
		string(Win32LobAppFileSystemDetectionDetectionTypeexists),
		string(Win32LobAppFileSystemDetectionDetectionTypemodifiedDate),
		string(Win32LobAppFileSystemDetectionDetectionTypenotConfigured),
		string(Win32LobAppFileSystemDetectionDetectionTypesizeInMB),
		string(Win32LobAppFileSystemDetectionDetectionTypeversion),
	}
}

func parseWin32LobAppFileSystemDetectionDetectionType(input string) (*Win32LobAppFileSystemDetectionDetectionType, error) {
	vals := map[string]Win32LobAppFileSystemDetectionDetectionType{
		"createddate":   Win32LobAppFileSystemDetectionDetectionTypecreatedDate,
		"doesnotexist":  Win32LobAppFileSystemDetectionDetectionTypedoesNotExist,
		"exists":        Win32LobAppFileSystemDetectionDetectionTypeexists,
		"modifieddate":  Win32LobAppFileSystemDetectionDetectionTypemodifiedDate,
		"notconfigured": Win32LobAppFileSystemDetectionDetectionTypenotConfigured,
		"sizeinmb":      Win32LobAppFileSystemDetectionDetectionTypesizeInMB,
		"version":       Win32LobAppFileSystemDetectionDetectionTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemDetectionDetectionType(input)
	return &out, nil
}
