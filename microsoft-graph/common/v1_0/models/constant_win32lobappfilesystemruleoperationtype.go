package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRuleOperationType string

const (
	Win32LobAppFileSystemRuleOperationTypecreatedDate   Win32LobAppFileSystemRuleOperationType = "CreatedDate"
	Win32LobAppFileSystemRuleOperationTypeexists        Win32LobAppFileSystemRuleOperationType = "Exists"
	Win32LobAppFileSystemRuleOperationTypemodifiedDate  Win32LobAppFileSystemRuleOperationType = "ModifiedDate"
	Win32LobAppFileSystemRuleOperationTypenotConfigured Win32LobAppFileSystemRuleOperationType = "NotConfigured"
	Win32LobAppFileSystemRuleOperationTypesizeInMB      Win32LobAppFileSystemRuleOperationType = "SizeInMB"
	Win32LobAppFileSystemRuleOperationTypeversion       Win32LobAppFileSystemRuleOperationType = "Version"
)

func PossibleValuesForWin32LobAppFileSystemRuleOperationType() []string {
	return []string{
		string(Win32LobAppFileSystemRuleOperationTypecreatedDate),
		string(Win32LobAppFileSystemRuleOperationTypeexists),
		string(Win32LobAppFileSystemRuleOperationTypemodifiedDate),
		string(Win32LobAppFileSystemRuleOperationTypenotConfigured),
		string(Win32LobAppFileSystemRuleOperationTypesizeInMB),
		string(Win32LobAppFileSystemRuleOperationTypeversion),
	}
}

func parseWin32LobAppFileSystemRuleOperationType(input string) (*Win32LobAppFileSystemRuleOperationType, error) {
	vals := map[string]Win32LobAppFileSystemRuleOperationType{
		"createddate":   Win32LobAppFileSystemRuleOperationTypecreatedDate,
		"exists":        Win32LobAppFileSystemRuleOperationTypeexists,
		"modifieddate":  Win32LobAppFileSystemRuleOperationTypemodifiedDate,
		"notconfigured": Win32LobAppFileSystemRuleOperationTypenotConfigured,
		"sizeinmb":      Win32LobAppFileSystemRuleOperationTypesizeInMB,
		"version":       Win32LobAppFileSystemRuleOperationTypeversion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRuleOperationType(input)
	return &out, nil
}
