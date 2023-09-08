package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeAuditResult string

const (
	CustomSecurityAttributeAuditResultfailure CustomSecurityAttributeAuditResult = "Failure"
	CustomSecurityAttributeAuditResultsuccess CustomSecurityAttributeAuditResult = "Success"
	CustomSecurityAttributeAuditResulttimeout CustomSecurityAttributeAuditResult = "Timeout"
)

func PossibleValuesForCustomSecurityAttributeAuditResult() []string {
	return []string{
		string(CustomSecurityAttributeAuditResultfailure),
		string(CustomSecurityAttributeAuditResultsuccess),
		string(CustomSecurityAttributeAuditResulttimeout),
	}
}

func parseCustomSecurityAttributeAuditResult(input string) (*CustomSecurityAttributeAuditResult, error) {
	vals := map[string]CustomSecurityAttributeAuditResult{
		"failure": CustomSecurityAttributeAuditResultfailure,
		"success": CustomSecurityAttributeAuditResultsuccess,
		"timeout": CustomSecurityAttributeAuditResulttimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomSecurityAttributeAuditResult(input)
	return &out, nil
}
