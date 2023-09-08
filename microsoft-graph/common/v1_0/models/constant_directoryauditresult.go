package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAuditResult string

const (
	DirectoryAuditResultfailure DirectoryAuditResult = "Failure"
	DirectoryAuditResultsuccess DirectoryAuditResult = "Success"
	DirectoryAuditResulttimeout DirectoryAuditResult = "Timeout"
)

func PossibleValuesForDirectoryAuditResult() []string {
	return []string{
		string(DirectoryAuditResultfailure),
		string(DirectoryAuditResultsuccess),
		string(DirectoryAuditResulttimeout),
	}
}

func parseDirectoryAuditResult(input string) (*DirectoryAuditResult, error) {
	vals := map[string]DirectoryAuditResult{
		"failure": DirectoryAuditResultfailure,
		"success": DirectoryAuditResultsuccess,
		"timeout": DirectoryAuditResulttimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryAuditResult(input)
	return &out, nil
}
